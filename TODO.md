# pho - Simple Blog Generator TODO List

## 🚀 LAUNCH CHECKLIST (TONIGHT/TOMORROW)
- [x] Clean navigation menu - remove unfinished links
- [x] Rename project from "bloggr" to "pho" 🍜
- [x] Update CLI commands: "cook" (generate) and "serve" (publish)  
- [x] Add Makefile with GitHub Pages publish strategy
- [x] Write first solid blog post
- [x] Create public GitHub repo  
- [x] Test `make serve` deployment
- [ ] Point idm.life domain to GitHub Pages
- [ ] 🎉 **GO LIVE!**

## High Priority (Core Features)
- [x] Analyze project structure and create initial Go module setup
- [x] Design markdown file structure with metadata handling
- [x] Create CLI command structure (generate, publish, serve)
- [x] Update folder structure to topic-based organization
- [x] Implement recursive markdown file discovery
- [x] Implement auto-tagging from folder structure
- [x] Implement basic HTML generation from markdown
- [ ] **Implement publish command** - Deploy generated blog to hosting platform

## Medium Priority (Essential Features)
- [x] Implement markdown parser with frontmatter support
- [x] Refactor code into proper Go packages (posts package)
- [x] Create simple HTML template
- [x] Create minimal CSS theme (Tokyo Night inspired)
- [x] Fix subfolder discovery and tag deduplication
- [x] Generate index page - List all posts with links, dates, and tags
- [x] Add navigation menu - Home link and basic site navigation to templates
- [ ] Create About page - Personal introduction, bio, and background
- [ ] Create Projects page - Links and descriptions to personal projects
- [ ] Create Categories page - Dynamic tag-based post grouping and navigation
- [ ] Post search functionality - Search posts by title, tags, or content
- [ ] Add index search functionality - Filter posts by title or tag on index page

## Low Priority (Future Enhancements)
- [ ] Enhanced personal introduction section - Picture, bio, links to music/code/projects
- [ ] Add footer - Social links and donation buttons
- [ ] Decide on tag ordering (folder tags vs frontmatter tags priority)
- [ ] Consider using slices package for tag deduplication optimization
- [ ] Add static asset handling and Bootstrap integration
- [ ] Implement serve command with local development server
- [ ] Add RSS feed generation

## Long-term Scaling Solutions (10+ Year Future-Proofing)
- [ ] Incremental builds - Only regenerate changed posts for performance
- [ ] Index pagination - Split large index pages for better load times
- [ ] CDN migration strategy - Move from GitHub Pages to Cloudflare Pages for performance
- [ ] SQLite metadata layer - Enable fast search and complex queries on post metadata
- [ ] Asset optimization pipeline - Image processing, minification, compression
- [ ] CI/CD pipeline - GitHub Actions for automated builds instead of local generation
- [ ] Content staging system - Preview workflows and draft management
- [ ] Multi-author collaboration - Branching strategies and content review processes
- [ ] Caching strategy - Template and content caching for faster regeneration
- [ ] Performance monitoring - Track build times and site performance metrics