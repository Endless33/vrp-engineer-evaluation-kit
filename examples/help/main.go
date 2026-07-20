package main

import (
	"fmt"
	"os"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/cli"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Help Example")
	fmt.Println("========================================")

	application := cli.New()

	fmt.Printf("Executable : %s\n", os.Args[0])

	fmt.Println()
	fmt.Println("Available Commands")
	fmt.Println("----------------------------------------")

	for _, command := range application.Commands() {
		fmt.Printf("  %-16s %s\n", command.Name, command.Description)
	}

	fmt.Println()
	fmt.Println("Example Usage")
	fmt.Println("----------------------------------------")
	fmt.Println("vrp-evaluator run")
	fmt.Println("vrp-evaluator version")
	fmt.Println("vrp-evaluator help")

	fmt.Println()
	fmt.Println("Displaying built-in help")
	fmt.Println("----------------------------------------")

	if err := application.Execute([]string{"help"}); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println()
	fmt.Println("Help demonstration completed successfully.")
}
