package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	ROOT_PATH = "C:"
)

/*
func ReadDirRecur(startPath string, target string) bool {
	dirEntries, err := os.ReadDir(startPath + "\\")
	if err != nil {
		return false
	}

	var found bool = false

	for _, entry := range dirEntries {
		if strings.Contains(entry.Name(), target) && entry.IsDir() == false {
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
*/

// TODO
func ReadDirectories(startPath string, inCh chan<- string) {
	entries, err := os.ReadDir(startPath + "\\")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("Sending ", entry.Name(), " into channel")
			inCh <- startPath + "\\" + entry.Name()
			ReadDirectories(startPath+"\\"+entry.Name(), inCh)
		}
	}
}

func SearchDir(dir, target string) bool {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	fmt.Println("\tOpened ", dir)
	for _, entry := range entries {
		fmt.Println("\t\tLooking at: ", entry.Name())
		if strings.Compare(entry.Name(), target) == 0 {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("usage <number of threads> <file to search>\n")
	}

	num_workers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	directory := make(chan string)

	go ReadDirectories(ROOT_PATH, directory)

	var wg sync.WaitGroup
	for i := 0; i < num_workers; i++ {
		wg.Add(1)
		go func() {
			for entries := range directory {
				if SearchDir(entries, os.Args[2]) {
					fmt.Println("Successfully found ", os.Args[2])
					os.Exit(0)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
