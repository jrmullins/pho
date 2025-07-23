# pho 🍜 - Simple Blog Generator

## Goals
- Make blogging easier
- I am very good at making markdown every day, I am not good at remembering various frameworks like hugo every time I want to publish an idea
- Make a .MD file, generate and publish through CLI
- Have an empty metadata section at the top of each file to help interpret what to do later (future proofing)
- jekyl/octa/hugo without all the bs

## Rough ideas
- Generate only 2 types of files first - index and posts
- Have some templating in place for a header and a body
- Use some extremely basic javascript libs like twitter bootstrap
- Some very basic CSS that looks nice and professional - hacker like. Maybe mono fonts with a minimal set of colors - Maybe Tokyo Nights. The goal here is post code and ideas, not sell anything. 
- Some light organization - maybe a file structure like posts/YYYY/MM/DD. We can use timestamps from the actual MD files (created_at and updated_at) for the time stamps in the posts that get generated.
- Maybe an optional yaml file with the same filename for metadata
- We need some mechanism to publish/unpublish. Maybe something idiomatic like appending dot (.) to make it a "hidden" file or underscore (_).

## Design Decisions

### ✅ Resolved
- **Publishing mechanism**: Use frontmatter `published: false` (more explicit and future-proof than filename prefixes)
- **Folder structure**: Freeform topic-based organization (code/, gaming/, work/, etc.) instead of rigid date folders
- **Auto-tagging**: Folder names automatically become tags (posts/code/go-tips.md gets "code" tag)
- **File discovery**: Recursive .md file walking through entire posts/ directory

### 🤔 Under Consideration

#### Publishing Platform Options
**GitHub Pages (Recommended - Simplicity First)**
- ✅ Zero configuration required
- ✅ Free forever with generous limits
- ✅ Custom domain support
- ✅ Integrates with existing Git workflow
- ✅ `bloggr publish` → push output/ to gh-pages branch
- ⚠️ Public repos only (unless GitHub Pro)

**Other Options (Future Consideration)**
- **Netlify**: Drop folder or Git integration (familiar, works well)
- **Cloudflare Pages**: Fast global CDN, excellent performance
- **Vercel**: Great developer experience, similar to Netlify
- **Cloud Functions**: AWS Lambda/DO Functions (overkill for static files)

#### Publishing Implementation Strategy
**Makefile vs Go CLI Decision**
- **Go CLI (`bloggr generate`)**: Complex logic (markdown parsing, templating, file discovery)
- **Makefile**: Simple operations (clean, build, publish, serve) - shell commands are more transparent
- Benefits: Standard tool, easier to customize, no recompiling for deployment tweaks

**Flexible Publishing Strategy Pattern**
Use environment variable to switch hosting providers:
```makefile
PUBLISH_STRATEGY ?= github-pages

publish: generate
	@$(MAKE) publish-$(PUBLISH_STRATEGY)

publish-github-pages:
	git checkout --orphan gh-pages 2>/dev/null || git checkout gh-pages
	cp -r output/* .
	git add .
	git commit -m "Deploy blog $$(date)"
	git push origin gh-pages
	git checkout main

publish-netlify:
	netlify deploy --prod --dir=output

# Future: publish-s3, publish-vercel, etc.
```

**Usage Examples:**
```bash
make publish                                    # Uses default (github-pages)
PUBLISH_STRATEGY=netlify make publish          # Override for specific deploy
```

This follows the Strategy Pattern - simplicity now, flexibility later without changing workflow.

## Quick Start 🚀

1. **Write posts** in `posts/` folder organized by topic:
   ```
   posts/
   ├── code/my-golang-post.md
   ├── gaming/game-review.md
   └── work/productivity-tips.md
   ```

2. **Cook your blog**:
   ```bash
   make cook
   ```

3. **Serve to the world**:
   ```bash
   make serve  # Deploys to GitHub Pages
   ```

## Commands

- `make cook` - Generate HTML from markdown posts 🍜
- `make serve` - Deploy to GitHub Pages 🥢  
- `make clean` - Clean generated files 🧹
- `make dev` - Start local development server
- `make help` - Show all commands

## Features

✅ **Simple** - Just markdown files and a Makefile  
✅ **Future-proof** - Your content stays portable  
✅ **Auto-tagging** - Folder names become tags automatically  
✅ **Tokyo Night theme** - Beautiful dark theme with code highlighting  
✅ **Featured posts** - Latest post displays prominently on homepage  
✅ **Zero config** - Works out of the box  

---

*TODO items moved to TODO.md for development tracking*
