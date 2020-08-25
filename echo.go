package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func echo(out *os.File) {
	args := os.Args[1:]

	//flag.Parse()
	//args := flag.Args()

	var result []string

	for _, arg := range args {
		result = append(result, arg)
	}
	result = append(result, "\n")
	if _, err := out.WriteString(strings.Join(result, " ")); err != nil {
		log.Fatal(err)
	}
}

func echo1_1(out *os.File) {
	var result []string

	args := os.Args
	args[0] = filepath.Base(args[0])

	for _, arg := range args {
		result = append(result, arg)
	}
	result = append(result, "\n")
	if out != nil {
		if _, err := out.WriteString(strings.Join(result, " ")); err != nil {
			log.Fatal(err)
		}
	}
}

func echo1_2(out *os.File) {
	var result []string
	args := os.Args[1:]

	for index, arg := range args {
		result = append(result, fmt.Sprintf("%d %v\n", index+1, arg))
	}

	result = append(result, "\n")
	if out != nil {
		if _, err := out.WriteString(strings.Join(result, "")); err != nil {
			log.Fatal(err)
		}
	}
}

func echo1_3(f1, f2 func(*os.File)) int64 {
	now := time.Now()
	var f1Time int64
	var f2Time int64

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func(ch chan interface{}) {
		f1(nil)
		close(ch1)
	}(ch1)

	go func(ch chan interface{}) {
		f2(nil)
		close(ch2)
	}(ch2)

	var ok1 = false
	var ok2 = false
	for {
		select {
		case <- ch1:
			ok1 = true
			f1Time = time.Since(now).Nanoseconds()
		case <- ch2:
			ok2 = true
			f2Time = time.Since(now).Nanoseconds()
		}

		if ok1 && ok2 {
			break
		}
	}

	return f1Time - f2Time
}
