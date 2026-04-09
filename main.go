package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/adomascx/Blockchain_v2/lib"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println(lib.PHA256([]byte(scanner.Text())))
	}
}
