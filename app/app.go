package app

import (
	"errors"
	"fmt"
	"os"
	"time"
)

func GoTouch(files []string) error {
	for _, f := range files {
		if _, err := os.Stat(f); errors.Is(err, os.ErrNotExist) {
			err := createFile(f)
			if err != nil {
				return errors.New(fmt.Sprintf("Couldn't create file: %s", err.Error()))
			}
		} else {
			err := updateLastAccess(f)
			if err != nil {
				return errors.New(fmt.Sprintf("Couldn't update last access time: %s", err.Error()))
			}
		}
	}
	return nil
}

func createFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func updateLastAccess(filePath string) error {
	err := os.Chtimes(filePath, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}
