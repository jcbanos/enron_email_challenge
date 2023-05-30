package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"strings"

	b "enron_email_challenge/indexer_f/directory_thread"
)

// Struct to represent an email
type Email struct {
	Subject string
	From    string
	To      string
	Date    string
	Content string
}

func main() {

	cpuprofile := "cpuprofile.prof"
	memprofile := "memprofile.prof"

	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// Get command line arguments then parse everything
	directory, zinchost := parseCommandLineArguments()
	parseAllDocuments(directory, zinchost)

	if memprofile != "" {
		f, err := os.Create(memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func parseCommandLineArguments() (string, string) {
	// Check number of arguments.
	if len(os.Args) != 3 {
		log.Fatal("Wrong number of arguments provided")
	}

	// Make sure the filepath contains no ..
	directory := os.Args[1]
	zinchost := os.Args[2]
	if strings.Contains(directory, "..") {
		log.Fatal("Poorly formatted path, please do not use ..")
	}
	return directory, zinchost
}

func parseAllDocuments(directory string, zinchost string) {
	emailsRootDirectory := filepath.Join(directory, "maildir")
	dirs, err := ioutil.ReadDir(emailsRootDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// Optimization #1, instead of processing directories sequentially, launch one go routine per directory.
	for _, dir := range dirs {
		b.DirectoryWalkThread(filepath.Join(emailsRootDirectory, dir.Name()), zinchost)
		// go b.DirectoryWalkThread(filepath.Join(emailsRootDirectory, dir.Name())) ##### Overwhelms the server
		// time.Sleep(150 * time.Millisecond) // To prevent overwhelming local sever
	}
}
