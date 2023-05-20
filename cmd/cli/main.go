package main

import (
	"fmt"
	"os"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/assembler"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd"
)

func main() {
	rootCmd, opt := cmd.NewCmdRoot(os.Stdout, os.Stderr)
	assembler.AddSubCommands(rootCmd, opt)

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("command execution error: %v\n", err)
		os.Exit(1)
	}
}
