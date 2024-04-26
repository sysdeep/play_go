package main

import (
	"errors"
	"flag"
)

type Config struct {
	FilePath string
}

func ConfigFromArgs() (error, *Config) {
	var file_path string
	flag.StringVar(&file_path, "file-path", "", "path to file")

	flag.Parse()

	if len(file_path) == 0 {
		return errors.New("file not specifed"), nil
	}

	return nil, &Config{file_path}

}
