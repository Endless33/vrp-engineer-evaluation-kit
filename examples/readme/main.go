package main

import (
	"fmt"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/version"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("README Example")
	fmt.Println("========================================")

	fmt.Println()
	fmt.Println("This example demonstrates the smallest")
	fmt.Println("possible program that imports the public")
	fmt.Println("VRP Engineer Evaluation Kit.")

	fmt.Println()
	fmt.Println("Project Information")
	fmt.Println("----------------------------------------")
	fmt.Printf("Version : %s\n", version.String())

	fmt.Println()
	fmt.Println("Purpose")
	fmt.Println("----------------------------------------")
	fmt.Println("- Public engineering evaluation")
	fmt.Println("- Public evidence generation")
	fmt.Println("- Public reports")
	fmt.Println("- Public integration examples")
	fmt.Println("- No protected runtime logic")

	fmt.Println()
	fmt.Println("For complete examples, see the other")
	fmt.Println("directories inside the examples folder.")

	fmt.Println()
	fmt.Println("README example completed successfully.")
}
