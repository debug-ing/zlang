package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/debug-ing/zlang/interpreter"
	"github.com/debug-ing/zlang/parser"
	"github.com/debug-ing/zlang/tokenizer"
)

// Show the source line and position of a parser or interpreter error
func showErrorSource(source []byte, pos tokenizer.Position, dividerLen int) {
	divider := strings.Repeat("-", dividerLen)
	if divider != "" {
		fmt.Println(divider)
	}
	lines := bytes.Split(source, []byte{'\n'})
	errorLine := string(lines[pos.Line-1])
	numTabs := strings.Count(errorLine[:pos.Column-1], "\t")
	fmt.Println(strings.Replace(errorLine, "\t", "    ", -1))
	fmt.Println(strings.Repeat(" ", pos.Column-1) + strings.Repeat("   ", numTabs) + "^")
	if divider != "" {
		fmt.Println(divider)
	}
}

func main() {
	if len(os.Args) < 2 || (os.Args[1] == "-stats" && len(os.Args) < 3) {
		fmt.Printf("usage: zlang [-stats] source_filename.z\n")
		os.Exit(1)
	}
	showStats := false
	filename := os.Args[1]
	execArgs := os.Args[2:]
	if os.Args[1] == "-stats" {
		showStats = true
		filename = os.Args[2]
		execArgs = os.Args[3:]
	}
	// check format file
	name := strings.Split(filename, "/")[len(strings.Split(filename, "/"))-1]
	if strings.Split(name, ".")[len(strings.Split(name, "."))-1] != "z" {
		fmt.Printf("error: %q is not a zlang file\n", os.Args[1])
		os.Exit(1)
	}

	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error reading %q\n", os.Args[1])
		os.Exit(1)
	}

	prog, err := parser.ParseProgram(input)
	if err != nil {
		errorMessage := fmt.Sprintf("%s", err)
		if e, ok := err.(parser.Error); ok {
			showErrorSource(input, e.Position, len(errorMessage))
		}
		fmt.Println(errorMessage)
		os.Exit(1)
	}

	startTime := time.Now()
	stats, err := interpreter.Execute(prog, &interpreter.Config{Args: execArgs})
	if err != nil {
		errorMessage := fmt.Sprintf("%s", err)
		if e, ok := err.(interpreter.Error); ok {
			showErrorSource(input, e.Position(), len(errorMessage))
		}
		fmt.Println(errorMessage)
		os.Exit(1)
	}
	if showStats {
		elapsed := time.Since(startTime)
		fmt.Printf("%s elapsed: %d ops (%.0f/s), %d builtin calls (%.0f/s), %d user calls (%.0f/s)\n",
			elapsed,
			stats.Ops, float64(stats.Ops)/elapsed.Seconds(),
			stats.BuiltinCalls, float64(stats.BuiltinCalls)/elapsed.Seconds(),
			stats.UserCalls, float64(stats.UserCalls)/elapsed.Seconds(),
		)
	}
}
