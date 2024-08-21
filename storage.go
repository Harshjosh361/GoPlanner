package main

import (
	"encoding/json"
	"os"
)

func LoadTask(filename string) ([]Task, error) {
	var tasks []Task

	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}
	defer file.Close()

	// Decoding the JSON data into Go structure
	err = json.NewDecoder(file).Decode(&tasks)
	return tasks, err
}

func SaveTask(filename string, tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(tasks)
}
