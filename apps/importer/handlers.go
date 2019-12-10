package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const baseFilePath = "importer/files/"
const baseTempFilePath = "importer/temp_file/"

func processFile(file string) {
	var files []string

	splitFile(file)

	files, _ = getFiles(baseTempFilePath)
	for i, chunk := range files {
		fmt.Println("Processing new chunk ... " + chunk)
		lines := readFile(chunk)
		if i == 0 {
			//remove headers
			lines = append(lines[:0], lines[1:]...)
		}
		fmt.Println("Inserting data into database")
		for _, line := range lines {
			columns := formatLine(line)
			err := insert("report", columns)
			if err != nil {
				logFailedLine(line)
			}
		}
		fmt.Println("Chunk done!")
	}
	if err := clearTempFiles(); err != nil {
		fmt.Println("Could not clear the temporary files!")
	}

	moveProcessedFile(baseFilePath + file)
	fmt.Println("Processing Complete!")
}

func splitFile(path string) {
	fmt.Println("Splitting file ...")
	var lines []string

	file, err := os.Open(baseFilePath + path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 0
	var slices uint64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if i == 5000 {
			fileName := baseTempFilePath + "file_to_process_" + strconv.FormatUint(slices, 10) + ".txt"
			f, err := os.Create(fileName)
			if err != nil {
				panic(err)
			}

			for _, line := range lines {
				if _, err = f.WriteString(line + "\n"); err != nil {
					panic(err)
				}
			}
			slices++
			i = 0
			lines = lines[:0]
			f.Close()
		}
		lines = append(lines, scanner.Text())
		i++
	}
	fmt.Println("Split into " + strconv.FormatUint(slices, 10) + " files.")
}

func clearTempFiles() error {
	files, err := filepath.Glob(filepath.Join(baseTempFilePath, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func moveProcessedFile(currentLocation string) {
	filename := strings.Split(currentLocation, "/")
	newLocation := baseFilePath + "processed/" + filename[len(filename)-1]
	err := os.Rename(currentLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func getFiles(directory string) ([]string, error) {
	var files []string
	f, err := os.Open(directory)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	defer f.Close()
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

func readFile(filename string) (lines []string) {
	file, err := os.Open(baseTempFilePath + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func formatLine(text string) []string {
	var result []string
	scanner := bufio.NewScanner(strings.NewReader(text))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		if word != "NULL" {
			word = "'" + strings.Replace(word, ",", ".", -1) + "'"
		}
		result = append(result, word)
	}
	return result
}

func logFailedLine(line string) {
	now := time.Now()
	fileName := baseFilePath + "failed/log_" + now.Format("2006_02_01") + ".log"
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(line + "\n"); err != nil {
		log.Println(err)
	}
}
