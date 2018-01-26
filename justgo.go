package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/satori/go.uuid"
	"github.com/urfave/cli"
)

var nonProblematicFiles = map[string]bool{}

func init() {
	// Initialize map of the non-problematic files to ignore.
	// Also, specify whether they could conflict with any files in the zip.
	// and therefore may need to be actively preserved (value == "true")
	nonProblematicFiles = map[string]bool{
		".git":       false,
		".gitignore": false,
		"justgo":     false,
		"LICENSE":    true,
		"README.md":  true,
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "justgo"
	app.Version = "1.2.1"
	app.Usage = "create a new skeleton project for a Go-based API [micro]service"
	app.UsageText = app.Name + " <path>"
	app.ArgsUsage = "path"

	app.Action = func(c *cli.Context) error {
		installPath := "."
		if c.NArg() > 0 {
			installPath = c.Args().Get(0)
		}
		buildProject(installPath)
		return nil
	}

	app.Run(os.Args)
}

func buildProject(path string) {
	fmt.Printf("Building a skeleton microservice at destination path: `%v` \n", path)
	exists, err := pathExists(path)
	abortIfErr(err)

	if exists == true { // need to make sure it's empty!
		isPathEmpty := folderIsEmpty(path)
		if !isPathEmpty {
			err := errors.New("There are files in `" + path + "`. Destination path must be empty. Aborting.")
			abortIfErr(err)
		}
	} else {
		err := os.MkdirAll(path, os.ModePerm)
		abortIfErr(err)
	}

	fileUrl := "https://github.com/inadarei/justgo-microservice/archive/master.zip"
	tmpFilePath := os.TempDir() + string(os.PathSeparator) + "justgo.zip"
	defer os.Remove(tmpFilePath)

	// Move all conflicting files to tmp dir and move them back post-build
	filesToMove := getConflictingFiles(path)
	uniqueToken := uuid.NewV4()
	uniqueTempFolder := filepath.Join(os.TempDir(), fmt.Sprintf("%s", uniqueToken))
	os.MkdirAll(uniqueTempFolder, os.ModePerm)
	defer os.Remove(uniqueTempFolder)

	if filesToMove != nil {
		for _, file := range filesToMove {
			srcPath := filepath.Join(path, file)
			tmpMovedFilePath := filepath.Join(uniqueTempFolder, file)
			err := os.Rename(srcPath, tmpMovedFilePath)
			abortIfErr(err)
			defer os.Remove(tmpMovedFilePath)
			defer os.Rename(tmpMovedFilePath, srcPath)
		}
	}

	err = downloadFile(tmpFilePath, fileUrl)
	abortIfErr(err)

	err = unzip(tmpFilePath, path, true)
	abortIfErr(err)

	cleanup(path)
	fmt.Println("Success. Happy coding!")
}

// Not everything in the download is useful. Remove garbage
func cleanup(path string) {
	var err error

	filesToRemove := []string{"CODE_OF_CONDUCT.md", "CONTRIBUTING.md", "LICENSE"}

	for _, file := range filesToRemove {
		err = os.Remove(filepath.Join(path, file))
		abortIfErr(err)
	}

	err = os.RemoveAll(filepath.Join(path, "cmd"))
	abortIfErr(err)
}

// Simple exit if error, to avoid putting same 4 lines of code in too many places
func abortIfErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

// Check whether path already exists
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// check whether a folder is empty
func folderIsEmpty(path string) bool {
	// What if we were given path to a file or smth?
	fi, err := os.Stat(path)
	abortIfErr(err)
	if fi.IsDir() == false {
		err := errors.New("Destination path `" + path + "` is not a folder!")
		abortIfErr(err)
	}

	f, err := os.Open(path)
	if err != nil {
		f.Close()
		abortIfErr(err)
	}
	defer f.Close()

	filenames, err := f.Readdirnames(0)
	abortIfErr(err)

	if !containsProblematicFiles(filenames) {
		return true
	}

	// If not already exited, scanning children must have errored-out
	abortIfErr(err)
	return false
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Unzip an archive and place contents at 'target' folder.
// Also allows skipping topLevel
func unzip(archive, target string, skipTop bool) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	firstItem := ""

	for _, file := range reader.File {
		var (
			modPath string
			err     error
		)

		if skipTop {
			if firstItem == "" {
				firstItem = file.Name
				continue
			} else {
				modPath, err = filepath.Rel(firstItem, file.Name)
				abortIfErr(err)
			}
		} else {
			modPath = file.Name
		}

		path := filepath.Join(target, modPath)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

// Check whether folder contains any files other than those specified as non-problematic.
func containsProblematicFiles(filesInDir []string) bool {
	if len(filesInDir) > len(nonProblematicFiles) {
		return true
	}

	// check if any files in the folder are considered to be problematic
	for _, filename := range filesInDir {

		// Is the file non-problematic?
		_, exists := nonProblematicFiles[filename]

		if !exists {
			return true
		}
	}
	return false
}

// Get Non-Problematic files in the 'target' folder that conflict with others in the zip.
func getConflictingFiles(path string) []string {
	var filesWithConflicts []string

	for filename, hasConflict := range nonProblematicFiles {

		exists, err := pathExists(filepath.Join(path, filename))
		abortIfErr(err)

		if exists && hasConflict == true {
			filesWithConflicts = append(filesWithConflicts, filename)
		}
	}
	return filesWithConflicts
}
