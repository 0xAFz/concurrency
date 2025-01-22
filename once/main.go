package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type FileLogger struct {
	file *os.File
	once *sync.Once
}

func NewFileLogger(path string) (FileLogger, error) {
	dir, err := os.MkdirTemp(path, "log")
	if err != nil {
		return FileLogger{}, err
	}

	file, err := os.CreateTemp(dir, "log")
	if err != nil {
		return FileLogger{}, err
	}

	return FileLogger{
		file: file,
		once: &sync.Once{},
	}, nil
}

func (f FileLogger) Log(s string) error {
	if _, err := fmt.Fprintf(f.file, "%s - %s\n", time.Now(), s); err != nil {
		return err
	}
	return nil
}

func (f FileLogger) Close() (err error) {
	f.once.Do(func() {
		err = f.file.Close()
	})

	return
}

func main() {
	logger, err := NewFileLogger(".")
	if err != nil {
		panic(err)
	}
	logger.Log("Hello, World!")

	if err := logger.Close(); err != nil {
		panic(err)
	}
	if err := logger.Close(); err != nil {
		panic(err)
	}
}
