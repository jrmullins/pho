package main

import (
	"github.com/jrmullins/pho/internal/posts"
	"github.com/jrmullins/pho/internal/site"
	"fmt"
	"os"

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

		err = os.MkdirAll("output", 0755)
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

func init() {
	rootCmd.AddCommand(cookCmd)
	rootCmd.AddCommand(serveCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
