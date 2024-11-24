package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Princesso/gopherquill/analyze"
	"github.com/Princesso/gopherquill/utils"
	"github.com/joho/godotenv"
)

func main() {
	var dir string

	// Command-line flag for directory
	flag.StringVar(&dir, "dir", ".", "Git directory to analyze")
	flag.Parse()

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve OpenAI API key
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY not found in environment variables")
	}

	// Initialize GPT client
	analyze.InitializeGPT(apiKey)

	// Resolve the Git directory
	gitDir, err := utils.GetGitRoot(dir)
	if err != nil {
		log.Fatalf("Failed to find Git root for directory %s: %v", dir, err)
	}

	// Get staged files
	stagedFiles, err := utils.GetStagedFiles(gitDir)
	if err != nil {
		log.Fatalf("Failed to get staged files: %v", err)
	}

	// Analyze each staged file
	for _, file := range stagedFiles {
		fullPath := filepath.Join(gitDir, file)
		fmt.Printf("Analyzing file: %s\n", fullPath)

		feedback, err := analyze.ReviewFile(fullPath)
		if err != nil {
			fmt.Printf("Error reviewing file %s: %v\n", fullPath, err)
		} else {
			fmt.Printf("Feedback for %s:\n%s\n\n", file, feedback)
		}
	}
}
