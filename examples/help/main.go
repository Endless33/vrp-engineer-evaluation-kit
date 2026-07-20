package main

import (
	"fmt"
	"log"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/cli"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Help Example")
	fmt.Println("========================================")

	application := cli.New()

	if err := cli.RegisterDefaultCommands(application); err != nil {
		log.Fatalf("failed to register commands: %v", err)
	}

	fmt.Println()
	fmt.Println("Built-in Help")
	fmt.Println("----------------------------------------")

	if err := application.Execute("help", nil); err != nil {
		log.Fatalf("help command failed: %v", err)
	}

	fmt.Println()
	fmt.Println("Help demonstration completed successfully.")
}
