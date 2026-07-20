package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/Endless33/vrp-engineer-evaluation-kit/internal/output"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("VRP Output Demo")
	fmt.Println("========================================")

	outDir := "demo-output"

	if err := output.EnsureDirectory(outDir); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}

	files := map[string]string{
		"evaluation.txt": "VRP Engineering Evaluation\nSTATUS=PASS\n",
		"runtime.log":    "Runtime initialized successfully.\n",
		"summary.txt":    "Public engineering evaluation completed.\n",
	}

	fmt.Println()
	fmt.Println("Writing Files")
	fmt.Println("----------------------------------------")

	for name, content := range files {
		path, err := output.WriteFile(outDir, name, []byte(content))
		if err != nil {
			log.Fatalf("failed to write %s: %v", name, err)
		}

		fmt.Printf("Created: %s\n", path)

		if !output.Exists(path) {
			log.Fatalf("verification failed for %s", path)
		}
	}

	fmt.Println()
	fmt.Println("Output Directory")
	fmt.Println("----------------------------------------")
	fmt.Println(filepath.Clean(outDir))

	fmt.Println()
	fmt.Println("All output files were written successfully.")
}