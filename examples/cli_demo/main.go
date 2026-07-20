package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/cli"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP CLI Demo")
	fmt.Println("========================================")

	application := cli.New()

	if err := cli.RegisterDefaultCommands(application); err != nil {
		log.Fatalf("failed to register default commands: %v", err)
	}

	fmt.Println()
	application.PrintHelp()

	fmt.Println()
	fmt.Println("Executing version command")
	fmt.Println("----------------------------------------")

	if err := application.Execute("version", nil); err != nil {
		log.Fatalf("command failed: %v", err)
	}

	fmt.Println()
	fmt.Println("CLI demonstration completed successfully.")
}
