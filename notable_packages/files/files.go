package files

import (
	"fmt"
	"os"
)

// Create and write to a new file
func CreateFile() {
	// Create new file if it does no exist
	file, err := os.Create("notable_packages/files/myfile.txt")

	if err != nil {
		fmt.Println("failed to create file!")
		return
	}

	fmt.Println("File created successfully!")

	// Write message to file as a string
	sizeString, err := file.WriteString("Hello, World!\n")

	if err != nil {
		fmt.Println("failed to write to file using WriteString!")
		return
	}

	fmt.Printf("Successfully wrote message to file! Size: %d bytes\n", sizeString)

	// Write message to file as a byte array
	sizeBytes, err := file.Write([]byte("Writing data to file!\n"))

	if err != nil {
		fmt.Println("failed to write to file using Write!")
		return
	}

	fmt.Printf("Successfully wrote message to file! Size: %d bytes\n", sizeBytes)

	// Remember to close file at the end
	file.Close()
}

// Read the entire file content all at once
func ReadFile() {
	// Read file all at once
	data, err := os.ReadFile("notable_packages/files/myfile.txt")

	if err != nil {
		fmt.Println("failed to read file!")
		return
	}

	fmt.Printf("Reading from file: %s\n", string(data))
}

// How to read the file piece by piece
func ReadFileInParts() {
	_, err := os.Open("notable_packages/files/myfile.txt")

	if err != nil {
		fmt.Println("failed to open file!")
		return
	}
}
