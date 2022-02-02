package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
)

func WriteToFile(filename string, fileMode os.FileMode, data string) error {
	file, err := os.OpenFile(filename, int(fileMode), 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write([]byte((data))); err != nil {
		return err
	}
	return nil
}
func main() {
	file, _ := os.Open("file1.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() { // internally, it advances token based on sperator
		fmt.Println((scanner.Text()))

	}
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter data to write-")
	scanner1.Scan()
	text := scanner1.Text()
	fmt.Println("ur entered data is ", text)
	err := WriteToFile("file1.txt", fs.FileMode(os.O_APPEND), text)
	if err != nil {
		fmt.Println("error in writing to file-", err)
	}
}
