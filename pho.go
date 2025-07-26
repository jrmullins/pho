package main

import (
	"github.com/jrmullins/pho/internal/posts"
	"github.com/jrmullins/pho/internal/site"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pho",
	Short: "A simple, future-proof blog generator 🍜",
	Long:  "pho converts markdown files into a static blog site with minimal configuration.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pho v1.0.0 - Simple blog generator 🍜")
		fmt.Println("Use 'pho --help' for available commands")
	},
}

var cookCmd = &cobra.Command{
	Use:   "cook",
	Short: "Cook HTML files from markdown posts 🍜",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Discovering posts...")
		postList, err := posts.Discover("posts")
		if err != nil {
			fmt.Printf("Error discovering posts: %v\n", err)
			return
		}

		fmt.Printf("Found %d posts\n", len(postList))

		err = os.MkdirAll("docs", 0755)
		if err != nil {
			fmt.Printf("Error creating output directory: %v\n", err)
			return
		}

		// Generate individual posts
		for _, post := range postList {
			if !post.Metadata.Published {
				fmt.Printf("Skipping draft: %s\n", post.Metadata.Title)
				continue
			}

			err := posts.GenerateHTML(post)
			if err != nil {
				fmt.Printf("Error generating HTML for %s: %v\n", post.Metadata.Title, err)
				continue
			}

			fmt.Printf("Generated: %s\n", post.Metadata.Title)
		}

		// Generate index page
		fmt.Println("Generating index page...")
		err = site.GenerateIndex(postList)
		if err != nil {
			fmt.Printf("Error generating index: %v\n", err)
			return
		}
		fmt.Println("Generated: index.html")

		// Generate about page
		fmt.Println("Generating about page...")
		err = site.GenerateAbout()
		if err != nil {
			fmt.Printf("Error generating about page: %v\n", err)
			return
		}
		fmt.Println("Generated: about.html")
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the blog locally for development 🥢",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting local server 🥢...")
		// TODO: Implement serve logic
	},
}

var postCmd = &cobra.Command{
	Use:   "post [category] [subcategory] [filename]",
	Short: "Create a new blog post with automatic metadata",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		category := args[0]
		subcategory := args[1] 
		filename := args[2]

		// Create the full directory path
		postDir := filepath.Join("posts", category, subcategory)
		err := os.MkdirAll(postDir, 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			return
		}

		// Create the full file path
		if filepath.Ext(filename) != ".md" {
			filename = filename + ".md"
		}
		filePath := filepath.Join(postDir, filename)

		// Check if file already exists
		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("File already exists: %s\n", filePath)
			return
		}

		// Create metadata template
		metadata := fmt.Sprintf(`---
title: ""
date: %s
tags: []
slug: ""
published: false
---

# Your Title Here

Write your post content here...
`, time.Now().Format(time.RFC3339))

		// Write the file
		err = os.WriteFile(filePath, []byte(metadata), 0644)
		if err != nil {
			fmt.Printf("Error creating post file: %v\n", err)
			return 
		}

		fmt.Printf("Created new post: %s\n", filePath)

		// Open in vim
		vimCmd := exec.Command("vim", filePath)
		vimCmd.Stdin = os.Stdin
		vimCmd.Stdout = os.Stdout
		vimCmd.Stderr = os.Stderr
		err = vimCmd.Run()
		if err != nil {
			fmt.Printf("Warning: Could not open vim: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cookCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(postCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
