package posts

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"gopkg.in/yaml.v3"
)

// This is Metadata about the post
type Metadata struct {
	Title     string    `yaml:"title"`
	Date      time.Time `yaml:"date"`
	Tags      []string  `yaml:"tags"`
	Slug      string    `yaml:"slug"`
	Published bool      `yaml:"published"`
}

type Post struct {
	Metadata   Metadata
	Content    string
	FilePath   string
	FolderTags []string
}

func Discover(postsDir string) ([]Post, error) {
	var posts []Post

	err := filepath.Walk(postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !strings.HasSuffix(path, ".md") {
			return nil
		}

		post, err := Parse(path, postsDir)
		if err != nil {
			fmt.Printf("Warning: Failed to parse %s: %v\n", path, err)
			return nil
		}

		posts = append(posts, post)
		return nil
	})

	return posts, err
}

func Parse(filePath, postsDir string) (Post, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var frontmatter []string
	var content []string
	inFrontmatter := false
	frontmatterEnded := false

	for scanner.Scan() {
		line := scanner.Text()

		if line == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				frontmatterEnded = true
				continue
			}
		}

		if inFrontmatter && !frontmatterEnded {
			frontmatter = append(frontmatter, line)
		} else if frontmatterEnded {
			content = append(content, line)
		}
	}

	var metadata Metadata
	if len(frontmatter) > 0 {
		yamlContent := strings.Join(frontmatter, "\n")
		if err := yaml.Unmarshal([]byte(yamlContent), &metadata); err != nil {
			return Post{}, fmt.Errorf("failed to parse frontmatter: %v", err)
		}
	}

	folderTags := ExtractFolderTags(filePath, postsDir)

	return Post{
		Metadata:   metadata,
		Content:    strings.Join(content, "\n"),
		FilePath:   filePath,
		FolderTags: folderTags,
	}, nil
}

func ExtractFolderTags(filePath, postsDir string) []string {
	relPath, err := filepath.Rel(postsDir, filePath)
	if err != nil {
		return nil
	}

	dir := filepath.Dir(relPath)
	if dir == "." {
		return nil
	}

	parts := strings.Split(dir, string(filepath.Separator))
	var tags []string
	for _, part := range parts {
		if part != "" {
			tags = append(tags, part)
		}
	}

	return tags
}

func DeduplicateTags(tags []string) []string {
	var result []string
	for _, tag := range tags {
		if !slices.Contains(result, tag) {
			result = append(result, tag)
		}
	}
	return result
}

type TemplateData struct {
	SiteTitle string
	Title     string
	Date      time.Time
	Content   template.HTML
	AllTags   []string
}

func GenerateHTML(post Post) error {
	var buf bytes.Buffer
	if err := goldmark.New().Convert([]byte(post.Content), &buf); err != nil {
		return fmt.Errorf("failed to convert markdown: %v", err)
	}

	tmpl, err := template.ParseFiles("templates/post.html")
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	allTags := DeduplicateTags(append(post.Metadata.Tags, post.FolderTags...))
	data := TemplateData{
		SiteTitle: "idm.life",
		Title:     post.Metadata.Title,
		Date:      post.Metadata.Date,
		Content:   template.HTML(buf.String()),
		AllTags:   allTags,
	}

	fileName := post.Metadata.Slug
	if fileName == "" {
		fileName = strings.ReplaceAll(strings.ToLower(post.Metadata.Title), " ", "-")
	}
	fileName += ".html"

	outputPath := filepath.Join("docs", fileName)
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer file.Close()

	return tmpl.Execute(file, data)
}

