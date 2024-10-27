package reader

import (
	"encoding/json"
	"os"
)


type StdinReader struct{}

func NewStdinReader() *StdinReader {
	return &StdinReader{}
}

func (r *StdinReader) ReadData() ([]int, error) {
	var data []int
	if err := json.NewDecoder(os.Stdin).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
