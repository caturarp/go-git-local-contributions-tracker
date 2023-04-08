package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

// Search for Git repos by accepting `path`
func scan(folder string) {
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
func recursiveScan(folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}

// returns the dot file for the `repos` list.
func getDotFilePath() string {
	// Creates it and the enclosing folder if it does not exist.
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	// our defined gogitlocalstats
	dotFile := usr.HomeDir + "/.gogitlocalstats"

	return dotFile
}

// // User represents a user account.
// type User struct {
//     Uid string
//     Gid string     // primary group ID.
//     Username string
//     Name string
//     HomeDir string
// }

// addNewSliceElementsToFile given a slice of strings representing paths, stores them
// to the filesystem
func addSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(filePath)
	repos := joinSlices(newRepos, existingRepos)
	dumpStringsSliceToFile(repos, filePath)
}

// Overwrites content to file in `filePath`
func dumpStringsSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	ioutil.WriteFile(filePath, []byte(content), 0755)
}

// Adds new element if not existing inside slice
func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}
	return existing
}

// Check if `slice` contains value
func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

// Gets the content of each line and parses it to a slice of strings.
func parseFileLinesToSlice(filePath string) []string {
	f := openFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			panic(err)
		}
	}

	return lines
}

// Opens file as per `filePath` and creates it if not existing
func openFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		if os.IsNotExist(err) {
			// if file not exist, then create file
			_, err = os.Create(filePath)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	return f
}

// Generate stats from contributions with email
func stats(email string) {
	print("stats")
}
