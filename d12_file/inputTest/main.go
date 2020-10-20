package main

import (
	"bufio"
	"os"
)

func main() {
	pWriter := bufio.NewWriter(os.Stdout)
	_, _ = pWriter.WriteString("Benjamin\n")
	_, _ = pWriter.WriteString("Stella\n")
	_ = pWriter.Flush()

}
