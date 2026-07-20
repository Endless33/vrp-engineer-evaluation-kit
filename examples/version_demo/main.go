package main

import (
	"fmt"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/version"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Version Information Demo")
	fmt.Println("========================================")

	info := version.Current()

	fmt.Println()
	fmt.Println("Version Information")
	fmt.Println("----------------------------------------")
	fmt.Printf("Version      : %s\n", info.Version)
	fmt.Printf("Release Name : %s\n", info.ReleaseName)
	fmt.Printf("Go Version   : %s\n", info.GoVersion)
	fmt.Printf("Platform     : %s\n", info.Platform)

	fmt.Println()
	fmt.Println("Compact Version String")
	fmt.Println("----------------------------------------")
	fmt.Println(version.String())

	fmt.Println()
	fmt.Println("Version demonstration completed successfully.")
}