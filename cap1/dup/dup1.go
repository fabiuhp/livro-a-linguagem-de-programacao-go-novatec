package main

import (
	"bufio"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	//NOTA: ignorando erros em potencial de input.Err()
	for line, n := range counts {
		if n > 1 {
			println("%d\t%s\n", n, line)
		}
	}
}
