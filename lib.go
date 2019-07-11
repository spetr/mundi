package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func errorChecker(err error) {
	if err != nil {
		fmt.Println(errors.Wrap(err, "ERROR:"))
		panic("")
	}
}

func appChDir() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	os.Chdir(dir)
}
