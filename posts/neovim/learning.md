---
title: "Learning Neovim"
date: 2025-07-24T11:00:00Z
tags: ["neovim"]
slug: "learning-neovim"
published: true
---

I'm slowly learning how to neovim. I'm going to try to put some gotchas in here.

## Commenting multiple lines of code
1. select the first caracter of your block
2. press Ctrl+V ( this is rectangular visual selection mode)
3. type j for each line more you want to be commented
4. type Shift-i (like I for "insert at start")
5. type // (or # or " or ...)
6. you will see the modification appearing only on the first line
7. IMPORTANT LAST STEP: type Esc key, and there you see the added character appear on all lines

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
