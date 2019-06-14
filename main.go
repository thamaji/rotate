package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var help, version bool
	var dir string
	var pattern string

	flag.BoolVar(&help, "h", false, "show help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.StringVar(&dir, "o", ".", "set output directory")
	flag.StringVar(&pattern, "p", "2006-01-02-15-04-05.log", "set output filename pattern")

	flag.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: " + os.Args[0] + " [OPTIONS]")
		fmt.Println()
		fmt.Println("Rotate stdin")
		fmt.Println()
		fmt.Println("Options:")
		flag.CommandLine.PrintDefaults()
		fmt.Println()
	}

	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	if version {
		fmt.Println("v1.0.0")
		return
	}

	if err := run(dir, pattern); err != nil {
		os.Exit(1)
	}
}

func run(dir string, pattern string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	path := filepath.Join(dir, time.Now().Format(pattern))

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	r := bufio.NewReader(os.Stdin)
	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			file.Close()
			if err != io.EOF {
				return err
			}
			return nil
		}

		if _, err := file.Write(line); err != nil {
			file.Close()
			return err
		}

		if isPrefix {
			continue
		}

		if _, err := file.Write([]byte{'\n'}); err != nil {
			file.Close()
			return err
		}

		next := filepath.Join(dir, time.Now().Format(pattern))
		if path != next {
			file.Close()
			path = next
			file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return err
			}
		}
	}
}
