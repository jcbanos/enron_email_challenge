package directory_thread

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Email struct {
	Subject string
	From    string
	To      string
	Date    string
	Content string
}

func NewEmail() Email {
	email := Email{}
	email.Subject = "Not available"
	email.From = "Not available"
	email.To = "Not available"
	email.Date = "Not available"
	return email
}

func DirectoryWalkThread(dirName string, zinchost string) {
	zincapi := "http://" + zinchost + ":5080/api/default/emails/_json"
	// Optimization #2, instead of growing an empty slice, find the number of files beforehand and allocate memory to a slice.
	numberOfFiles, err := extractNumberOfFiles(dirName)
	if err != nil {
		fmt.Println(err)
	}
	emailFiles := make([]Email, numberOfFiles)
	err = extractInformationFromEmails(dirName, emailFiles)
	if err != nil {
		fmt.Println(err)
	}
	sendEmails(emailFiles, zincapi)
}

func extractNumberOfFiles(dirName string) (int, error) {
	numberOfFiles := 0
	err := filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			numberOfFiles++
		}
		return nil
	})
	return numberOfFiles, err
}

func extractInformationFromEmails(dirName string, files []Email) error {
	count := 0
	err := filepath.WalkDir(dirName, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			files[count] = parseEmailFileToJson(path)
			count++
		}
		return nil
	})
	return err
}

// Optimization #3 instead of using scanner which requires more string conversions, read file and parse it as bytes
func parseEmailFileToJson(emailPath string) Email {
	email := NewEmail()
	file_bytes, err := os.ReadFile(emailPath)
	if err != nil {
		fmt.Println(err)
	}
	index := bytes.Index(file_bytes, []byte("\r\n\r\n"))
	header := file_bytes[:index]
	content := file_bytes[index:]
	email.Content = strings.TrimSpace(string(content))
	headerParts := bytes.Split(header, []byte("\r\n"))
	subject_bytes := []byte("Subject:")
	fromBytes := []byte("From:")
	toBytes := []byte("To:")
	dateBytes := []byte("Date:")
	for _, header := range headerParts {
		if bytes.HasPrefix(header, subject_bytes) {
			subject := strings.TrimSpace(strings.TrimPrefix(string(header), "Subject:"))
			email.Subject = subject
		} else if bytes.HasPrefix(header, fromBytes) {
			from := strings.TrimSpace(strings.TrimPrefix(string(header), "From:"))
			email.From = from
		} else if bytes.HasPrefix(header, toBytes) {
			to := strings.TrimSpace(strings.TrimPrefix(string(header), "To:"))
			email.To = to
		} else if bytes.HasPrefix(header, dateBytes) {
			date := strings.TrimSpace(strings.TrimPrefix(string(header), "Date:"))
			email.Date = date
		}
	}
	return email
}

func sendEmails(emails []Email, zincapi string) {

	// Put emails in required format
	jsonEmails, err := json.Marshal(emails)
	if err != nil {
		log.Fatal(err)
	}

	// Create request and send it
	req, err := http.NewRequest("POST", zincapi, strings.NewReader(string(jsonEmails)))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("root@example.com", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
