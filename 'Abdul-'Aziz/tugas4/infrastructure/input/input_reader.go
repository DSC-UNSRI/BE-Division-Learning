package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputReader struct{}

func NewInputReader() *InputReader {
    return &InputReader{}
}

func (r *InputReader) ReadLine() string {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return strings.TrimSpace(scanner.Text())
}

func (r *InputReader) ReadInt(min, max int) int {
    for {
        input := r.ReadLine()
        num, err := strconv.Atoi(input)
        if err == nil && num >= min && num <= max {
            return num
        }
        fmt.Printf("Input harus angka %d-%d: ", min, max)
    }
}