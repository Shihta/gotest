package main

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

func exe_cmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).CombinedOutput()
	if err != nil {
		fmt.Printf("<< %s >>\n", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func exe_cmd2(cmd string) {
	fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).CombinedOutput()
	if err != nil {
		fmt.Printf("<< %s >>\n", err)
	}
	fmt.Printf("%s", out)
	// wg.Done() // Need to signal to waitgroup that this goroutine is done
}

func main() {
	fmt.Println("Hello, World111!")
	// wg := new(sync.WaitGroup)
	// wg.Add(1)
	exe_cmd2("go")
	// go exe_cmd("ls", wg)
	// go exe_cmd("env", wg)
	// go exe_cmd("ls ~", wg)
	// wg.Wait()
	// fmt.Println()
}
