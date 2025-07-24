# Claude Session Notes

## Project: pho - Simple Blog Generator
**Goal**: Create a simple, future-proof static blog generator in Go

## Key Design Decisions Made
- **Publishing**: Use frontmatter `published: false` instead of filename prefixes
- **Organization**: Topic-based folders (code/, gaming/, work/) instead of date-based
- **Auto-tagging**: Folder names automatically become tags
- **Templates**: Inline CSS until 5+ templates, then extract to shared file
- **Publishing Strategy**: GitHub Pages with Makefile (not Go CLI for deployment)
- **Architecture**: Clean package separation (posts, site, main CLI)

## Architecture Overview
```
pho/
├── main.go              # CLI with "cook" and "serve" commands
├── internal/posts/      # Post processing logic
│   ├── posts.go        # Core post parsing and HTML generation
│   └── cmd.go          # Deleted - moved logic to main
├── internal/site/       # Site-wide logic (index, etc.)
│   └── index.go        # Index page generation with featured post
├── templates/           # HTML templates
│   ├── index.html      # Homepage with featured post + list
│   └── post.html       # Individual post template
├── posts/               # Content organization
│   ├── code/
│   ├── gaming/
│   └── work/
└── output/              # Generated HTML (kept for simple GitHub Pages)
```

## Current Status
- ✅ Core functionality working (markdown → HTML)
- ✅ Beautiful Tokyo Night theme
- ✅ Auto-tagging from folder structure
- ✅ Featured post on index page
- ✅ Navigation menu (home only for launch)
- 🚀 Ready for launch tonight/tomorrow!

## Next Session Priorities
1. Clean navigation (remove unfinished links)
2. Rename to "pho" with cute command names
3. Create Makefile for GitHub Pages deployment
4. Write first blog post
5. Launch publicly!