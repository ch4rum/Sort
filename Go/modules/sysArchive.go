package modules

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Archives struct {
	path string
}

func (self Archives) OpenArchive() (data []int, err error) {
	archive, err := os.Open(self.path)
	if err != nil {
		return nil, fmt.Errorf("Error open file %v", err)
	}
	defer archive.Close()
	scanner := bufio.NewScanner(archive)
	for scanner.Scan() {
		line := scanner.Text()
		num, err2 := strconv.Atoi(line)
		if err2 != nil {
			fmt.Printf("Cannot be converted '%s' to integer: %v\n", line, err2)
			continue
		}
		data = append(data, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}
	return data, nil
}

func (self Archives) ExportArchive(data []int) error {
	archive, err := os.Create(self.path)
	if err != nil {
		return fmt.Errorf("Error created file: %v", err)
	}
	defer archive.Close()
	writer := bufio.NewWriter(archive)
	for _, num := range data {
		_, err := writer.WriteString(fmt.Sprintf("%d\n", num))
		if err != nil {
			return fmt.Errorf("Error writing data: %v", err)
		}
	}
	return writer.Flush()
}
