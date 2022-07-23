package main

import (
	"file_search/workers"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	ROOT_PATH = "C:"
)

func ReadDirRecur(startPath string, target string) bool {
	fmt.Println("Opening dir: ", startPath)
	dirEntries, err := os.ReadDir(startPath + "\\")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	var found bool = false

	for _, entry := range dirEntries {
		fmt.Println("\tlooking at: ", entry.Name())
		if strings.Contains(entry.Name(), target) && entry.IsDir() == false {
			println("found: ", entry.Name())
			return true
		}
		if entry.IsDir() {
			if found = ReadDirRecur(startPath+"\\"+entry.Name(), target); found {
				return true
			}
		}
	}
	return found
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage <file to search>\n")
	}

	worker_pool := workers.NewWorkerPool(workers.MAX_WORKERS)

	if found := ReadDirRecur(ROOT_PATH, os.Args[1]); found {
		fmt.Println("Successfully found: ", os.Args[1])
	} else {
		fmt.Println("Failed to find: ", os.Args[1])
	}
}
