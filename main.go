package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Search for Git repos by accepting `path`
func scan(path string) {
	fmt.Printf("Found folders:\n\n")
	// get slice of repos
	repositories := recursiveScan(folder)
	filePath := getDotFilePath()
	addSliceElementsToFile(filePath, repositories)
	fmt.Printf("\n\nSuccessfully added\n\n")
}

// scanGitFolders returns a list of subfolders of `folder` ending with `.git`.
// Returns the base folder of the repo, the .git folder parent.
// Recursively searches in the subfolders by passing an existing `folders` slice.
func scanGitFolders(folders []string, folder string) []string {
	// trim unwanted `/`
	folder = strings.TrimSuffix(folder, "/")

	f, err := os.Open(folder)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")
				fmt.Println(path)
				folders = append(folders, path)
				continue
			}
			// skip if node_modules or vendor since it took times
			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}
			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}

// Recursive search of git repos inside the `folder` subtree
func recursiveScanFolder(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

// Generate stats from contributions with email
func stats(email string) {
	print("stats")
}

func main() {
	var folder, email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "your@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
