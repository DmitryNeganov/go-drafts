package main

import (
	"bufio"
	"fmt"
	"go-drafts/pkg/train"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	// Чтение количества тестов
	var t int
	fmt.Fscanln(reader, &t)

	startTests(t)
}

func startTests(t int) {
	for i := 0; i < t; i++ {
		train.DoFirstTest()
	}
}
