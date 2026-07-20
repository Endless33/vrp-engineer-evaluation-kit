package main

import (
	"fmt"
	"os"
)

const (
	Version = "v0.1.0"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Version :", Version)
	fmt.Println("Repository: vrp-engineer-evaluation-kit")
	fmt.Println()
	fmt.Println("This repository provides a public engineering")
	fmt.Println("evaluation environment for VRP.")
	fmt.Println()
	fmt.Println("Available commands:")
	fmt.Println("  evaluate     Run evaluation")
	fmt.Println("  report       Generate report")
	fmt.Println("  version      Show version")
	fmt.Println()

	if len(os.Args) < 2 {
		fmt.Println("No command specified.")
		fmt.Println("Use one of the available commands.")
		return
	}

	switch os.Args[1] {

	case "evaluate":
		fmt.Println("Evaluation command is not yet implemented.")

	case "report":
		fmt.Println("Report generation is not yet implemented.")

	case "version":
		fmt.Println(Version)

	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
