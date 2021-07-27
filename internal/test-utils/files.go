package test_utils

import (
	"io/ioutil"
	"log"
	"os"
)

func CreateTmpFile(fileName string, content []byte) (string, func()) {
	tmpFile, err := ioutil.TempFile("", fileName)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := tmpFile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}

	return tmpFile.Name(), func() {
		err := os.Remove(tmpFile.Name())
		if err != nil {
			log.Fatal("Failed to remove tmp file")
		}
	}
}
