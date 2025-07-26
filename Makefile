# pho - Simple Blog Generator 🍜
# Makefile for building and deploying your blog

.PHONY: cook serve clean help

# Default publishing strategy - can be overridden
PUBLISH_STRATEGY ?= github-pages

# Cook the blog (generate HTML from markdown)
cook:
	@echo "🍜 Cooking your blog..."
	go run pho.go cook

# Serve/Deploy the blog to hosting platform
serve: cook
	@echo "🥢 Serving your blog..."
	@$(MAKE) serve-$(PUBLISH_STRATEGY)

# GitHub Pages deployment strategy
serve-github-pages:
	@echo "📦 Committing changes to main branch..."
	git add .
	git commit -m "Update blog: $$(date)" || echo "No changes to commit"
	git push origin main
	@echo "📤 Deploying to GitHub Pages..."
	git checkout pages || git checkout -b pages
	git merge main --no-edit
	make cook
	git add output/
	git commit -m "Deploy blog: $$(date)" || echo "No changes to deploy"
	git push origin pages
	git checkout main
	@echo "✅ Blog served at: https://jrmullins.github.io/pho"

# Future deployment strategies can be added here
serve-netlify:
	@echo "🌐 Deploying to Netlify..."
	netlify deploy --prod --dir=output
	@echo "✅ Blog served on Netlify"

# Clean generated files
clean:
	@echo "🧹 Cleaning up..."
	rm -rf output/

# Development server (future implementation)
dev: cook
	@echo "🥢 Starting development server..."
	cd output && python3 -m http.server 8080

# Show available commands
help:
	@echo "🍜 pho - Simple Blog Generator"
	@echo ""
	@echo "Available commands:"
	@echo "  cook              Cook HTML from markdown posts"
	@echo "  serve             Serve/deploy blog (default: GitHub Pages)"
	@echo "  clean             Clean generated files"
	@echo "  dev               Start local development server"
	@echo "  help              Show this help message"
	@echo ""
	@echo "Deployment strategies:"
	@echo "  PUBLISH_STRATEGY=github-pages make serve  (default)"
	@echo "  PUBLISH_STRATEGY=netlify make serve"
	@echo ""
	@echo "🍜 Happy blogging!"