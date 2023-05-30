package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type Response struct {
	Took       int          `json:"took"`
	Total      int          `json:"total"`
	From       int          `json:"from"`
	Size       int          `json:"size"`
	ScanSize   int          `json:"scan_size"`
	Hits       []Email      `json:"hits"`
	TookDetail []TookDetail `json:"took_detail"`
}

type Email struct {
	Timestamp string `json:"_timestamp"`
	Content   string
	From      string
	Subject   string
	To        string
	Date      string
}

type TookDetail struct {
	Total            int
	WaitQueue        int `json:"wait_queue"`
	ClusterTotal     int `json:"cluster_total"`
	ClusterWaitQueue int `json:"cluster_total_wait_queue"`
}

var zincsearchUrl string

func main() {

	if len(os.Args) < 2 {
		log.Fatal("The number of args should be 2.")
	}

	zincsearchUrl = "http://" + os.Args[1] + ":5080/api/default/_search"

	r := chi.NewRouter()

	r.Get("/*", handleHomePage)

	r.Get("/search", handleSearchRequest)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", r),
	)
}

func handleHomePage(w http.ResponseWriter, r *http.Request) {
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "frontend", "dist")
	if _, err := os.Stat(filesDir + r.URL.Path); errors.Is(err, os.ErrNotExist) {
		http.ServeFile(w, r, filepath.Join(filesDir, "index.html"))
	}
	http.ServeFile(w, r, filesDir+r.URL.Path)
}

func handleSearchRequest(w http.ResponseWriter, r *http.Request) {

	//create query
	searchTerm := r.URL.Query().Get("searchTerm")
	jsonString := fmt.Sprintf(`{
		"query": {
			"sql": "SELECT * FROM emails WHERE match_all('%s')",
			"from": 0,
			"size": 10
		}
	}
	`, searchTerm) // possible sql injection?
	hits := getHits(jsonString)
	jsonHits, err := json.Marshal(hits)
	checkError(err)

	w.Write(jsonHits)
}

func getHits(jsonString string) []Email {
	jsonBody := []byte(jsonString)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, zincsearchUrl, bodyReader)
	checkError(err)
	req.SetBasicAuth("root@example.com", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	res, err := http.DefaultClient.Do(req)
	checkError(err)
	defer res.Body.Close()
	checkError(err)
	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	json.Unmarshal(b, &response)
	hits := response.Hits
	return hits
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
