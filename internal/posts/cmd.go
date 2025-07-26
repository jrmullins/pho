package posts

import (
	"fmt"
	"os"
)

func DiscoverAndGenerate() {
	fmt.Println("Discovering posts...")
	postList, err := Discover("posts")
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

	for _, post := range postList {
		if !post.Metadata.Published {
			fmt.Printf("Skipping draft: %s\n", post.Metadata.Title)
			continue
		}

		err := GenerateHTML(post)
		if err != nil {
			fmt.Printf("Error generating HTML for %s: %v\n", post.Metadata.Title, err)
			continue
		}

		fmt.Printf("Generated: %s\n", post.Metadata.Title)
	}

}
