---
title: "Learning Neovim"
date: 2025-07-24T11:00:00Z
tags: ["neovim"]
slug: "learning-neovim"
published: true
---

I'm slowly learning how to neovim. I'm going to try to put some gotchas in here.

## Identation weirdness
I love the `o` and `O` commands for inserting lines in edit mode before or after the line your on, especially for adding code. However it was slowly driving me mad that it would always end up on the beginning of the line, and when writing code I'd always have to tab over. Turns out I needed this in my init.lua after the plugin section:
```
-- Indentation settings
vim.cmd("filetype plugin indent on")
vim.o.autoindent = true
vim.o.smartindent = true
vim.o.smarttab = true
```

## Simple file manulation
`:Explore`
