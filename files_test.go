package files

import (
	"encoding/hex"
	"testing"
)

func TestCreateNewFile(t *testing.T) {
	filename := "temp.txt"
	filesize := int64(1024)
	_, err := CreateNewFile("", filename, filesize)
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadFile(t *testing.T) {
	var err error
	filename := "temp.txt"
	filesize := int64(1024)
	_, err = CreateNewFile("", filename, filesize)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ReadFile("", filename, filesize)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateRandomFile(t *testing.T) {
	filename := "temp.txt"
	filesize := int64(2048)
	_, err := CreateRandomFile("", filename, filesize)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateRandomHashedFile(t *testing.T) {
	filename := "hashed.txt"
	filesize := int64(2048)
	size, hashSum, err := CreateRandomHashedFile("", filename, filesize)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("File=%v, Size=%v, Hash=%v", filename, size, hex.EncodeToString(hashSum))
}

func TestReadRandomHashedFile(t *testing.T) {
	filename := "hashed.txt"
	filesize := int64(2048)
	writesize, writeHashSum, err := CreateRandomHashedFile("", filename, filesize)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Write: File=%v, Size=%v, Hash=%v", filename, writesize, hex.EncodeToString(writeHashSum))
	readsize, readHashSum, err := ReadHashedFile("", filename, filesize)
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
	filesize := int64(2048)
	hashsize := int64(1024)
	writesize, writeHashSum, err := CreateRandomHashedFileWithHashSize("", filename, filesize, hashsize)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Write: File=%v, Size=%v, Hash=%v", filename, writesize, hex.EncodeToString(writeHashSum))
	readsize, readHashSum, err := ReadHashedFile("", filename, hashsize)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Read: File=%v, Size=%v, Hash=%v", filename, readsize, hex.EncodeToString(readHashSum))
	if hex.EncodeToString(writeHashSum) != hex.EncodeToString(readHashSum) {
		t.Fatal("Write and Read Hash doesn't match")
	}
}

func TestReadRandomHashedFileWithHashSize2(t *testing.T) {
	filename := "hashed.txt"
	filesize := int64(2048)
	hashsize := int64(1024)
	writesize, writeHashSum, err := CreateRandomHashedFileWithHashSize("", filename, filesize, hashsize)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Write: File=%v, Size=%v, Hash=%v", filename, writesize, hex.EncodeToString(writeHashSum))

	data, err := ReadFile("", filename, hashsize)
	if err != nil {
		t.Fatal(err)
	}
	hashed, err := HashData(data, hashsize)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Read: File=%v, Size=%v, Hash=%v", filename, hashsize, hex.EncodeToString(hashed))
	if hex.EncodeToString(writeHashSum) != hex.EncodeToString(hashed) {
		t.Fatal("Write and Read Hash doesn't match")
	}
}
