package files

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var buffersize = int64(1024)

// CreateNewFile - creates new file in given size
func CreateNewFile(testDir string, filename string, filesize int64) ([]byte, error) {
	data := make([]byte, filesize)
	_, err := rand.Read(data)
	if err != nil {
		fmt.Printf("Error while generating random bytes, error=%v\n", err)
		return nil, err
	}
	fmt.Printf(fmt.Sprintf("data first 10 bytes=[%v]\n", data[:9]))

	fmt.Printf(fmt.Sprintf("Creating file %v in size %v bytes\n", filename, filesize))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error creating file %v, error=%v\n", filename, err))
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	// io.CopyN()
	n, err := io.Copy(file, buf)
	if err != nil {
		fmt.Printf("Error copying bytes to file, error=%v\n", err)
		return nil, err
	}
	fmt.Printf("Finished writing, length=%v\n", n)

	return data, nil
}

func ReadFile(testDir string, filename string, filesize int64) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error opening file %v, error=%v\n", filename, err))
		return nil, err
	}
	data := make([]byte, filesize)
	var n int
	if n, err = file.Read(data); err != nil {
		fmt.Printf(fmt.Sprintf("Error reading file %v, error=%v\n", filename, err))
		return nil, err
	}
	if int64(n) != filesize {
		fmt.Printf(fmt.Sprintf("Error expected file %v, got=%v\n", filesize, n))
		return nil, err
	}
	return data, nil
}

func CreateRandomFile(testDir string, filename string, filesize int64) (int64, error) {
	var err error
	var i, n int64

	fmt.Printf(fmt.Sprintf("Creating file %v in size %v bytes\n", filename, filesize))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error creating file %v, error=%v\n", filename, err))
		return 0, err
	}

	for i < filesize && err == nil {
		fmt.Printf("working: i=%v, error=%v\n", i, err)
		if filesize-i < buffersize {
			buffersize = int64(filesize - i)
		}
		if n, err = io.CopyN(file, rand.Reader, buffersize); err != nil {
			fmt.Printf(fmt.Sprintf("Error writing to file %v, error=%v\n", filename, err))
			log.Fatal(err)
		}
		i += n
	}
	fmt.Printf("finished: i=%v, error=%v\n", i, err)
	if i != filesize {
		errMsg := fmt.Sprintf("Error - writing to file %v, size mismatch expectation, expected=%v, got=%v\n", filename, filesize, i)
		fmt.Printf(errMsg)
		return i, errors.New(errMsg)
	}

	return i, nil
}

func ReadHashedFile(testDir string, filename string, filesize int64) (int64, []byte, error) {
	var err error
	var i, n int64
	hash := sha512.New()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error opening file %v, error=%v\n", filename, err))
		return 0, nil, err
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf(fmt.Sprintf("Error closing file %v, error=%v\n", filename, err))
			panic(err)
		}
	}()

	for i < filesize && err == nil {
		if filesize-i < buffersize {
			buffersize = int64(filesize - i)
		}
		if n, err = io.CopyN(hash, file, buffersize); err != nil {
			fmt.Printf(fmt.Sprintf("Error writing to file %v, error=%v\n", filename, err))
			log.Fatal(err)
		}
		i += n
	}

	return int64(i), hash.Sum(nil), nil
}

// CreateRandomHashedFile - CreateRandomHashedFileWithHashSize wrapper: hashsize = filesize
func CreateRandomHashedFile(testDir string, filename string, filesize int64) (int64, []byte, error) {
	return CreateRandomHashedFileWithHashSize(testDir, filename, filesize, filesize)
}

// CreateRandomHashedFileWithHashSize - creates a file with random data, returns size written, hash512 according to hashsize value and error
func CreateRandomHashedFileWithHashSize(testDir string, filename string, filesize int64, hashsize int64) (int64, []byte, error) {
	var err error
	var i, n int64
	hash := sha512.New()

	fmt.Printf(fmt.Sprintf("Creating file %v in size %v bytes\n", filename, filesize))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf(fmt.Sprintf("Error creating file %v, error=%v\n", filename, err))
		return 0, nil, err
	}

	writer := io.MultiWriter(file, hash)

	for i < filesize && err == nil {
		if filesize-i < buffersize {
			buffersize = int64(filesize - i)
		}
		if hashsize <= i {
			writer = file
		}
		if n, err = io.CopyN(writer, rand.Reader, buffersize); err != nil {
			fmt.Printf(fmt.Sprintf("Error writing to file %v, error=%v\n", filename, err))
			log.Fatal(err)
		}
		i += n
	}
	if i != filesize {
		errMsg := fmt.Sprintf("Error - writing to file %v, size mismatch expectation, expected=%v, got=%v\n", filename, filesize, i)
		fmt.Printf(errMsg)
		return i, nil, errors.New(errMsg)
	}

	return i, hash.Sum(nil), nil
}
