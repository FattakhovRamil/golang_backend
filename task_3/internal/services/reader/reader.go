package reader

import (
	"encoding/json"
	"os"
)

type Reader interface {
	ReadData() ([]int, error)
}

type FileReader struct {
	filename string
}

func NewFileReader(filename string) *FileReader {
	return &FileReader{filename: filename}
}

func (r *FileReader) ReadData() ([]int, error) {
	file, err := os.Open(r.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
