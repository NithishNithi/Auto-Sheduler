package main

import (
	"context"
	"log"
	"time"

	"github.com/google/go-github/github"
	"github.com/robfig/cron"
	"golang.org/x/oauth2"
)

func main() {
	// Create a GitHub client using your personal access token
	ctx := context.Background()
	token := ""
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Define your cron schedule (e.g., run every day at midnight)
	c := cron.New()
	c.AddFunc("0 0 23 * * *", func() {

		// Read new content from a data source (e.g., a file or an API)

		// Define the file to update in your GitHub repository
		repoOwner := "NithishNithi"
		repoName := "Auto-Sheduler"
		fileName := "example.txt"
		commitMessage := "Update example.txt"

		// Read the current content of the file from GitHub
		fileContent, _, _, err := client.Repositories.GetContents(ctx, repoOwner, repoName, fileName, nil)
		if err != nil {
			log.Printf("Error getting file content: %s\n", err)
			return
		}

		// Extract the current content and SHA
		currentContent, err := fileContent.GetContent()
		if err != nil {
			log.Printf("Error getting content: %s\n", err)
			return
		}
		sha := fileContent.GetSHA()

		// Append new content
		currentDate := time.Now().Format("2006-01-02 15:04:05")
		newContent := currentContent + "\n" + currentDate + "\n"

		// Update the file content on GitHub
		_, _, err = client.Repositories.UpdateFile(ctx, repoOwner, repoName, fileName, &github.RepositoryContentFileOptions{
			Message: &commitMessage,
			Content: []byte(newContent),
			SHA:     &sha,
		})
		if err != nil {
			log.Printf("Error updating file: %s\n", err)
			return
		}

		log.Println("File updated successfully")
	})

	// Start the cron scheduler
	c.Start()

	// Keep the program running
	select {}
}
