package data

import (
	"bee/bbee/models"
	"bee/bbee/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getHomeUrl() string {
	dirname, err := os.UserHomeDir()
	check(err)
	return dirname
}

func getDataUrl() string {
	var home = getHomeUrl()
	var path = ".local/bin/scripthub/data.json"
	return fmt.Sprintf("%s/%s", home, path)
}

func getSourcesUrl() string {
	var home = getHomeUrl()
	var path = ".local/bin/scripthub/sources.json"
	return fmt.Sprintf("%s/%s", home, path)
}

func getLastCommandUrl() string {
	var home = getHomeUrl()
	var path = ".local/bin/scripthub/lastcommand"
	return fmt.Sprintf("%s/%s", home, path)
}

func ReadItems() []models.IndexItem {
	path := getDataUrl()
	dat, err := ReadResource(path)

	if err != nil {
		return []models.IndexItem{}
	}

	var items []models.IndexItem
	json.Unmarshal([]byte(dat), &items)
	return items
}

// Function to read a local or remote source File
// and returns a list of typed models.SourceFile
func ReadSourceFile(path string) []models.SourceFile {
	dat, err := ReadResource(path)

	if err != nil {
		return []models.SourceFile{}
	}

	var sources []models.SourceFile
	json.Unmarshal([]byte(dat), &sources)
	return sources
}

// Function that reads the user's default source File
// and returns a list of typed models.SourceFile
func ReadUserSources() []models.SourceFile {
	path := getSourcesUrl()
	return ReadSourceFile(path)
}

// Reads the user's local Sources file and returns the contents
func ReadSourcesRaw() string {
	path := getSourcesUrl()
	dat, err := readFile(path)

	if err != nil {
		return ""
	}

	return string(dat)
}

// Function that takes either a valid local file path or  a valid remote URL
// and returns either an error or the contents of the resource, as a string
func ReadResource(path string) (string, error) {
	if utils.IsHttpUrl(path) {
		return readUrl(path)
	} else {
		return readFile(path)
	}
}

// Reads data from a remote URL that is accessible somehow
func readUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//Convert the body to type string
	return string(body), nil
}

// Reads data from a local file, given a path
func readFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func WriteLastCommand(command string) {
	path := getLastCommandUrl()
	d1 := []byte(command)
	err := os.WriteFile(path, d1, 0644)
	check(err)
}

func WriteItems(items []models.IndexItem) {
	path := getDataUrl()
	json, err := json.MarshalIndent(items, "", "  ")
	check(err)
	ferr := os.WriteFile(path, json, 0644)
	check(ferr)
}

func WriteSources(sources []models.SourceFile) {
	path := getSourcesUrl()
	json, err := json.MarshalIndent(sources, "", "  ")
	check(err)
	ferr := os.WriteFile(path, json, 0644)
	check(ferr)
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return false
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
