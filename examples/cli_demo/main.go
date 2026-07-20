package main

import (
	"fmt"
	"os"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/cli"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP CLI Demo")
	fmt.Println("========================================")

	application := cli.New()

	fmt.Printf("Executable : %s\n", os.Args[0])
	fmt.Printf("Arguments  : %v\n\n", os.Args[1:])

	fmt.Println("Registered Commands")
	fmt.Println("----------------------------------------")

	for _, command := range application.Commands() {
		fmt.Printf("Name        : %s\n", command.Name)
		fmt.Printf("Description : %s\n", command.Description)
		fmt.Println()
	}

	fmt.Println("Executing built-in help command...")
	fmt.Println("----------------------------------------")

	if err := application.Execute([]string{"help"}); err != nil {
		fmt.Printf("Command failed: %v\n", err)
		return
	}

	fmt.Println()
	fmt.Println("CLI demonstration completed successfully.")
}
