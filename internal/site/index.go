package site

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"github.com/jrmullins/pho/internal/posts"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
)

type IndexPost struct {
	Title   string
	Date    time.Time
	Slug    string
	AllTags []string
	Content template.HTML
}

type IndexData struct {
	SiteTitle  string
	AuthorName string
	AuthorBio  string
	Posts      []IndexPost
}

func GenerateIndex(postList []posts.Post) error {
	var indexPosts []IndexPost

	// Convert posts to index format and filter published only
	for _, post := range postList {
		if !post.Metadata.Published {
			continue
		}

		allTags := posts.DeduplicateTags(append(post.Metadata.Tags, post.FolderTags...))

		slug := post.Metadata.Slug
		if slug == "" {
			slug = generateSlug(post.Metadata.Title)
		}

		// Convert markdown to HTML for content
		var buf bytes.Buffer
		if err := goldmark.New().Convert([]byte(post.Content), &buf); err != nil {
			return err
		}

		indexPosts = append(indexPosts, IndexPost{
			Title:   post.Metadata.Title,
			Date:    post.Metadata.Date,
			Slug:    slug,
			AllTags: allTags,
			Content: template.HTML(buf.String()),
		})
	}

	// Sort posts by date (newest first)
	sort.Slice(indexPosts, func(i, j int) bool {
		return indexPosts[i].Date.After(indexPosts[j].Date)
	})

	// Create template data
	data := IndexData{
		SiteTitle:  "idm.life",
		AuthorName: "John Mullins", // TODO: Make configurable
		AuthorBio:  "Welcome to my blog where I share thoughts on code, hacking, and life.",
		Posts:      indexPosts,
	}

	// Parse template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		return err
	}

	// Generate index.html
	outputPath := filepath.Join("docs", "index.html")
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}

func GenerateAbout() error {
	// Create template data
	data := IndexData{
		SiteTitle:  "idm.life",
		AuthorName: "John Mullins", // TODO: Make configurable
		AuthorBio:  "Welcome to my blog where I share thoughts on code, hacking, and life.",
	}

	// Parse template
	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		return err
	}

	// Generate about.html
	outputPath := filepath.Join("docs", "about.html")
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}

func generateSlug(title string) string {
	// Simple slug generation - same logic as in posts package
	return strings.ReplaceAll(strings.ToLower(title), " ", "-")
}

