package main

import (
	"bytes"
	"fmt"
	"os"
)

// RLECompress compresses data using Run-Length Encoding
func RLECompress(data []byte) []byte {
	var compressed bytes.Buffer

	i := 0
	for i < len(data) {
		count := 1
		// Count the number of consecutive bytes that are the same
		for i+1 < len(data) && data[i] == data[i+1] && count < 255 {
			count++
			i++
		}
		// Write the byte and its count to the compressed buffer
		compressed.WriteByte(data[i])
		compressed.WriteByte(byte(count))

		i++
	}

	return compressed.Bytes()
}

// RLEDecompress decompresses data using Run-Length Encoding
func RLEDecompress(data []byte) ([]byte, error) {
	var decompressed bytes.Buffer

	// The length of the compressed data should be even (pairs of value and count)
	if len(data)%2 != 0 {
		return nil, fmt.Errorf("invalid RLE data")
	}

	for i := 0; i < len(data); i += 2 {
		value := data[i]   // The byte value
		count := data[i+1] // The run length

		// Write the byte value 'count' times to the decompressed buffer
		for j := 0; j < int(count); j++ {
			decompressed.WriteByte(value)
		}
	}

	return decompressed.Bytes(), nil
}

func main() {
	// Read input from a text file
	inputFile := "input.txt"
	input, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	// Compress the data
	fmt.Println("Compressing data...")
	compressed := RLECompress(input)

	// Calculate and display the compression ratio
	originalSize := len(input)
	compressedSize := len(compressed)
	compressionRatio := float64(compressedSize) / float64(originalSize)
	savings := 1.0 - compressionRatio

	fmt.Printf("\nOriginal Size: %d bytes\n", originalSize)
	fmt.Printf("Compressed Size: %d bytes\n", compressedSize)
	fmt.Printf("Compression Ratio: %.2f\n", compressionRatio)
	fmt.Printf("Space Saved: %.2f%%\n", savings*100)

	// Write compressed data to a text file
	compressedFile := "compressed.txt"
	err = os.WriteFile(compressedFile, compressed, 0644)
	if err != nil {
		fmt.Println("Error writing compressed file:", err)
		return
	}

	// Decompress the data
	fmt.Println("\nDecompressing data...")
	decompressed, err := RLEDecompress(compressed)
	if err != nil {
		fmt.Println("Error decompressing data:", err)
		return
	}

	// Write decompressed data to a text file
	decompressedFile := "decompressed.txt"
	err = os.WriteFile(decompressedFile, decompressed, 0644)
	fmt.Println("Decompressed data written to", decompressedFile)
	if err != nil {
		fmt.Println("Error writing decompressed file:", err)
		return
	}
}
