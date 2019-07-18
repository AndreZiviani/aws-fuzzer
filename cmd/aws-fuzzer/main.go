package main

import (
	"fmt"
	"os"

	"github.com/andreziviani/aws-fuzzer"
)

func main() {
	fzf, err := ec2fzf.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fzf.Run()
}
