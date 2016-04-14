package files

import (
	"encoding/hex"
	"testing"
)

func TestCreateNewFile(t *testing.T) {
	_, err := CreateNewFile("", "temp.txt", 1024)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	var err error
	_, err = CreateNewFile("", "temp.txt", 1024)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ReadFile("", "temp.txt", 1024)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateRandomFile(t *testing.T) {
	_, err := CreateRandomFile("", "temp.txt", 2048)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateRandomHashedFile(t *testing.T) {
	filename := "hashed.txt"
	size, hashSum, err := CreateRandomHashedFile("", "hashed.txt", 2048)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("File=%v, Size=%v, Hash=%v", filename, size, hex.EncodeToString(hashSum))
}

func TestReadRandomHashedFile(t *testing.T) {
	filename := "hashed.txt"
	writesize, writeHashSum, err := CreateRandomHashedFile("", "temp.txt", 2048)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Write: File=%v, Size=%v, Hash=%v", filename, writesize, hex.EncodeToString(writeHashSum))
	readsize, readHashSum, err := ReadHashedFile("", "temp.txt", 2048)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Read: File=%v, Size=%v, Hash=%v", filename, readsize, hex.EncodeToString(readHashSum))
	if hex.EncodeToString(writeHashSum) != hex.EncodeToString(readHashSum) {
		t.Fatal("Write and Read Hash doesn't match")
	}
}

func TestReadRandomHashedFileWithHashSize(t *testing.T) {
	filename := "hashed.txt"
	writesize, writeHashSum, err := CreateRandomHashedFileWithHashSize("", "temp.txt", 2048, 1024)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Write: File=%v, Size=%v, Hash=%v", filename, writesize, hex.EncodeToString(writeHashSum))
	readsize, readHashSum, err := ReadHashedFile("", "temp.txt", 1024)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Read: File=%v, Size=%v, Hash=%v", filename, readsize, hex.EncodeToString(readHashSum))
	if hex.EncodeToString(writeHashSum) != hex.EncodeToString(readHashSum) {
		t.Fatal("Write and Read Hash doesn't match")
	}
}
