package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func echo() {
	args := os.Args[1:]

	//flag.Parse()
	//args := flag.Args()

	var result []string

	for _, arg := range args {
		result = append(result, arg)
	}
	result = append(result, "\n")
	fmt.Println(strings.Join(result, " "))
}

/*
	練習問題1-1
		echoプログラムを修正して、そのプログラムを起動したコマンド名である`os.Args[0[`も表示されるようにしなさい
 */
func echo1_1() {
	var result []string

	args := os.Args
	args[0] = filepath.Base(args[0])

	for _, arg := range args {
		result = append(result, arg)
	}
	result = append(result, "\n")
	fmt.Println(strings.Join(result, " "))
}

/*
	練習問題1-2
		echoプログラムを修正して、個々の引数のインデック于sと値の組みを１行ごとに表示しなさい
 */
func echo1_2() {
	var result []string
	args := os.Args[1:]

	for index, arg := range args {
		result = append(result, fmt.Sprintf("%d %v\n", index+1, arg))
	}

	result = append(result, "\n")
	fmt.Println(strings.Join(result, ""))
}

/*
	練習問題1-3
		非効率な可能性のあるバージョン（echo2）と`strings.Join`を使ったバージョン（echo3）とで、実行時の性能差を計測しなさい
 */
func echo1_3() {
	fmt.Print(procCompleteFuncTime(echo2, echo3), "[ns]")
}

func procCompleteFuncTime(f1, f2 func()) int64 {
	now := time.Now()
	var f1Time int64
	var f2Time int64

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	go func(ch chan interface{}) {
		f1()
		close(ch1)
	}(ch1)

	go func(ch chan interface{}) {
		f2()
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

func echo2() {
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}