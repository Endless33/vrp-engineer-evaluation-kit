package main

import (
	"fmt"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/version"
)

func main() {
	info := version.Current()

	fmt.Println("========================================")
	fmt.Println("VRP Engineer Evaluation Kit")
	fmt.Println("Version Information")
	fmt.Println("========================================")

	fmt.Printf("Version      : %s\n", info.Version)
	fmt.Printf("Release Name : %s\n", info.ReleaseName)
	fmt.Printf("Go Version   : %s\n", info.GoVersion)
	fmt.Printf("Platform     : %s\n", info.Platform)

	fmt.Println()
	fmt.Println("Formatted Version")
	fmt.Println("----------------------------------------")
	fmt.Println(version.String())

	fmt.Println()
	fmt.Println("Version example completed successfully.")
}