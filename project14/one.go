package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error Ocured While Creating File :", err)
		return
	}

	content := "Hii My Name Is Sparsh"
	byteWritten, error := io.WriteString(file, content)
	fmt.Println("Bytes written Are :", byteWritten)
	if error != nil {
		fmt.Println("File Written Succesfully")
	}

	file1, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error While Opening File ", err)
	}

	// Create a buffer
	buffer := make([]byte, 1024)
	for {
		n, err := file1.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("An error occurred while reading the file:", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}

	// We can read a file without creating a buffer also

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("error while reading file ", err)
		return
	}
	fmt.Println(string(data))

	defer file.Close()
	defer file1.Close()
}
