package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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
