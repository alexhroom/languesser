package src

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// Learn reads a directory of files for a language and produces a language file.
func Learn(directory string, language string) {

	var outputMap map[string]float64

	// read directory
	files, readDirErr := os.ReadDir(directory)
	if readDirErr != nil {
		log.Fatal(readDirErr)
	}

	// calculate distribution for each file
	for i, file := range files {
		path := directory + file.Name()
		fileContent, readFileErr := os.ReadFile(path)
		if readFileErr != nil {
			log.Fatal(readFileErr)
		}
		fileText := string(fileContent)
		fileDistribution := Distribution(fileText)
		if i == 0 {
			outputMap = fileDistribution
		} else {
			outputMap = AverageMaps(outputMap, fileDistribution)
		}
	}

	// output new map to file
	outputLanguageFile(outputMap, language)
}

// outputLanguageFile writes a language file if one does not exist,
// and edits the language file otherwise
func outputLanguageFile(languageMap map[string]float64, language string) {

	path := "./langfiles/" + language + ".json"
	_, err := os.Stat(path)

	if errors.Is(err, os.ErrNotExist) { // if language not already learned, create file
		writeLanguageFile(languageMap, language)
	} else if err == nil { // if language already learned, combine results
		oldLangMap := readLanguageFile(language)
		newLangMap := AverageMaps(oldLangMap, languageMap)
		writeLanguageFile(newLangMap, language)
	} else { // else it is a different miscellaneous error
		log.Fatal(err)
	}
}

// readLanguageFile reads in a language file with name `language` and turns it into a map
// language files should be in JSON format
func readLanguageFile(language string) map[string]float64 {

	var languageMap map[string]float64

	// read in file from langfiles subfolder
	path := "./langfiles/" + language + ".json"
	languageFile, langReadErr := os.ReadFile(path)
	if langReadErr != nil {
		log.Fatal(langReadErr)
	}

	jsonErr := json.Unmarshal(languageFile, &languageMap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return languageMap
}

// writeLanguageFile writes a language map to a file
func writeLanguageFile(languageMap map[string]float64, language string) {

	languageJSON, jsonErr := json.Marshal(languageMap)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// set path to write file to
	path := "./langfiles/" + language + ".json"

	// write file with read + write access for all users
	os.WriteFile(path, languageJSON, 0666)
}

// AverageMaps takes two language maps and averages each value
func AverageMaps(map1 map[string]float64, map2 map[string]float64) map[string]float64 {

	outputMap := make(map[string]float64)

	for k, v1 := range map1 {
		v2, exists := map2[k]
		if !exists { // add entry directly to output map if in map1 and not map2
			outputMap[k] = v1
		} else { // average values and assign to output map
			outputMap[k] = (v1 + v2) / 2
		}
	}

	// add entry to output map if in map2 and not map1
	for k, v := range map2 {
		_, exists := map1[k]
		if !exists {
			outputMap[k] = v
		}
	}
	return outputMap
}
