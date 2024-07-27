# RLE Compression

This project implements a simple Run-Length Encoding (RLE) compression and decompression algorithm in Go. The program reads input data from a text file, compresses it using RLE, writes the compressed data to another text file, and then decompresses the data to verify the compression.

## Files

- `main.go`: The main Go program that performs RLE compression and decompression.
- `input.txt`: The input text file containing the data to be compressed.
- `compressed.txt`: The output text file containing the compressed data.
- `decompressed.txt`: The output text file containing the decompressed data.

## How to Use

1. **Prepare the Input File**: Create a text file named `input.txt` in the project directory and add the data you want to compress. For example:

    ```text
    aaabbbbccdddddd
    ```

2. **Run the Program**: Execute the Go program to compress and decompress the data.

    ```sh
    go run main.go
    ```

3. **Check the Output**: After running the program, you will find two new files in the project directory:
    - `compressed.txt`: Contains the compressed data.
    - `decompressed.txt`: Contains the decompressed data, which should match the original input.

