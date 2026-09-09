package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"resume/generate/data"
)

type Resume = data.Resume

func ReadResumeJSON(path string) (*Resume, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var resume Resume
	err = json.Unmarshal(bytes, &resume)
	if err != nil {
		return nil, err
	}

	return &resume, nil
}

func FindAndReadResume(arg string) (*Resume, error) {
	// First priority is the argument
	resume, err := ReadResumeJSON(arg)
	if resume != nil {
		return resume, nil
	}

	// Check in a data folder
	resume, err = ReadResumeJSON(fmt.Sprintf("data/%s", arg))
	return resume, err
}
