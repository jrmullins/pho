# pho 🍜 - Simple Blog Generator

Checkout the live version at [https://idm.life](https://idm.life)

## Goals
- Make blogging easier
- I am very good at making markdown every day, I am not good at remembering various frameworks like hugo every time I want to publish an idea
- Make a .MD file, generate and publish through CLI
- Have an empty metadata section at the top of each file to help interpret what to do later (future proofing)
- jekyl/octa/hugo without all the bs


## Quick Start 🚀

1. **Write posts** in `posts/` folder organized by topic:
   ```
   posts/
   ├── code/my-golang-post.md
   ├── gaming/rpg/game-review.md
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

