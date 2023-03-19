package main

import (
	"bufio"
	"distributedComputing-CA1/grammar"
	"fmt"
	"log"
	"os"
	"sync"
)

type Line struct {
	lineNumber int
	content    string
}

type WorkerNumber struct {
	mu     sync.Mutex
	number int
}

func (c *WorkerNumber) Dec() {
	c.mu.Lock()
	c.number--
	c.mu.Unlock()
}

func (c *WorkerNumber) GetNumber() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.number
}

func worker(chIn chan Line, chOut chan Line, quit chan int, workerNumber *WorkerNumber) {
	for {
		select {
		case x := <-chIn:
			sentences := grammar.SentenceTokenizer(x.content)
			result := ""
			for _, sentence := range sentences {
				sentence = grammar.OrdinalizeNumbers(sentence)
				result += grammar.CapitalizeFirstLetter(sentence)
			}
			x.content = result
			chOut <- x
		case <-quit:
			workerNumber.Dec()
			if workerNumber.GetNumber() == 0 {
				close(chOut)
			}
			return
		}
	}
}

func ReadFileLineByLine(filepath string, ch chan Line, quit chan int) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lineNumber := 1
	for scanner.Scan() {
		ch <- Line{lineNumber, scanner.Text()}
		lineNumber++
	}
	close(ch)
	close(quit)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func WriteFileLineByLine(filepath string, ch chan Line, done chan bool) {
	f, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()
	dw := bufio.NewWriter(f)
	lineToContent := make(map[int]string)
	for l := range ch {
		lineToContent[l.lineNumber] = l.content
	}
	for i := 1; i <= len(lineToContent); i++ {
		dw.WriteString(lineToContent[i])
		if (i != len(lineToContent)) {
			dw.WriteString("\n")
		}
	}

	dw.Flush()
	done <- true
}

func main() {
	var n int
	fmt.Scan(&n)
	chIn := make(chan Line, n)
	chOut := make(chan Line, n)
	quit := make(chan int)
	done := make(chan bool)
	workerNumber := WorkerNumber{number: n}
	go ReadFileLineByLine("input.txt", chIn, quit)
	for i := 0; i < n; i++ {
		go worker(chIn, chOut, quit, &workerNumber)
	}
	go WriteFileLineByLine("output.txt", chOut, done)

	<-done
}