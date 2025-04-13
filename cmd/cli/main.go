package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"to-do-planner/internal/config"
	"to-do-planner/internal/db"
	service "to-do-planner/internal/service/provider"
)

func main() {
	db.InitDB()

	cfg := config.Load(db.DB)

	if len(os.Args) < 2 {
		// No command provided
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "list":
		// We expect a second argument: "task" or "provider"
		if len(os.Args) < 3 {
			fmt.Println(`Usage: cli list <task|provider>`)
			os.Exit(1)
		}
		switch os.Args[2] {
		case "task":
			listTasks(cfg)
		case "provider":
			listProviders(cfg)
		default:
			fmt.Printf("Unknown list command: %s\n", os.Args[2])
			os.Exit(1)
		}

	case "add":
		// e.g. "add provider"
		if len(os.Args) < 3 {
			fmt.Println(`Usage: cli add <provider> [--json ...]`)
			os.Exit(1)
		}
		switch os.Args[2] {
		case "provider":
			addProviderCmd := flag.NewFlagSet("add provider", flag.ExitOnError)
			// For example, a JSON file or inline JSON
			jsonInput := addProviderCmd.String("json", "", "JSON data for provider")
			addProviderCmd.Parse(os.Args[3:])
			addProvider(cfg, *jsonInput)
		default:
			fmt.Printf("Unknown add command: %s\n", os.Args[2])
			os.Exit(1)
		}

	case "load":
		// e.g. "load task"
		if len(os.Args) < 3 {
			fmt.Println(`Usage: cli load <task>`)
			os.Exit(1)
		}
		switch os.Args[2] {
		case "task":
			loadTasks(cfg)
		default:
			fmt.Printf("Unknown load command: %s\n", os.Args[2])
			os.Exit(1)
		}

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func listTasks(cfg *config.Config) {
	task, err := cfg.Services.Task.GetTasks(context.Background())
	if err != nil {
		fmt.Println("Error loading task:", err)
		return
	}

	fmt.Println("[LIST TASKS]")
	for _, task := range task {
		fmt.Printf("Name: %s, Difficulty: %d, Duration: %d, ProviderName: %s\n", task.Name, task.Difficulty, task.Duration, task.ProviderName)
	}
}

func listProviders(cfg *config.Config) {
	provider, err := cfg.Services.Provider.GetProviders(context.Background())
	if err != nil {
		fmt.Println("Error loading provider:", err)
		return
	}

	fmt.Println("[LIST PROVIDERS]")
	for _, provider := range provider {
		fmt.Printf("Name: %s, API URL: %s\n", provider.ProviderName, provider.APIURL)
	}
}

func addProvider(cfg *config.Config, jsonData string) {
	var provider service.Provider
	err := json.Unmarshal([]byte(jsonData), &provider)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	err = cfg.Services.Provider.Create(context.Background(), provider)
	if err != nil {
		fmt.Println("Error creating provider:", err)
		return
	}
	log.Printf("Provider created: %+v", provider)
}

func loadTasks(cfg *config.Config) {
	err := cfg.Services.Task.LoadTasks(context.Background())
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	log.Println("Tasks loaded successfully")
}

// printUsage shows available top-level commands.
func printUsage() {
	fmt.Println(`Usage:
  cli list <task|provider>
  cli add <provider> [--json='{"name":"..."}']
  cli load <task>

Examples:
  cli list task
  cli list provider
  cli add provider --json='{"name":"NewProvider","api_url":"/api/new"}'
  cli load task
`)
}
