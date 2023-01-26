package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/shirou/gopsutil/disk"
)

const EmptySearchErrorText = "Text to search is empty"
const NoPartitionErrorText = "No partitions in system"
const WriteTextToSearchText = "Write text to search"
const TextToSearchLength = 0
const CurrentPath = "/."
const Slash = "/"

func main() {
	textToSearch := getUserTextToSearch()
	if len(textToSearch) == TextToSearchLength {
		fmt.Println(EmptySearchErrorText)
	}
	partitions, err := getComputerPartitionList()
	if err != nil {
		fmt.Println(NoPartitionErrorText)
	}
	scanPartitions(textToSearch, partitions)
}

func scanPartitions(textToSearch string, partitionsNames []disk.PartitionStat) {
	var filesPath []string

	for _, partition := range partitionsNames {
		getFiles(textToSearch, &filesPath, partition.Mountpoint+CurrentPath)
	}

	for _, path := range filesPath {
		fmt.Println(path)
	}
}

func getComputerPartitionList() ([]disk.PartitionStat, error) {
	return disk.Partitions(false)
}

func getFiles(textToSearch string, filesPath *[]string, discName string) {
	files, err := os.ReadDir(discName)
	if err != nil {
	}

	for _, file := range files {
		fileName := discName + Slash + file.Name()
		if file.IsDir() {
			getFiles(textToSearch, filesPath, fileName)
		}
		if strings.Contains(fileName, textToSearch) {
			*filesPath = append(*filesPath, fileName)
		}
	}
}

func getUserTextToSearch() string {
	fmt.Println(WriteTextToSearchText)
	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	return input.Text()
}
