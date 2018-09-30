package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// GetEnv ... get value from env variables or return the fallback value
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// LsFiles ... list file in directory
func LsFiles(pattern string) {
	err := filepath.Walk(pattern,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
