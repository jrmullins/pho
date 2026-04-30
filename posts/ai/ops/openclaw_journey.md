---
title: "My AI Perspective Shift"
date: 2026-04-30T12:00:00Z
tags: ["ai", "openclaw", "self-hosted", "ollama", "openrouter", "picoclaw"]
slug: "openclaw-journey"
published: true
---

# My AI Perspective Shift

Back in January, two friends and I started experimenting with OpenClaw. We each took a different path. One ran it on a rack-mounted Mac Mini. Another set it up on Proxmox. I went with a Digital Ocean VPS.

That first week taught me a lot about the ecosystem. I built custom tools for notifications and email. The real win was something I called NightShift, a program that would pick up projects defined in Linear and push pull requests to GitHub overnight.

Then I ran the OpenClaw update command and everything crashed.

I spent the next week rebuilding. Docker Compose, a local LLM on Ollama, and several forks of OpenClaw later, I had a clear goal: bring that $130 weekly bill down.

I landed on devstral-small-2:24b with Picoclaw. Running locally on my M1 MacBook meant roughly one minute per request. Fast enough to be useful, slow enough to notice.

Eventually I settled on a shared CPU Vultr instance running Arch Linux, Picoclaw, and OpenRouter models. A small Go binary I felt comfortable modifying myself. The whole setup runs on hardware you wouldn't blink at.

That's the real lesson here. You don't need a powerful machine to run an AI agent. You just need reliable access to an LLM. OpenRouter's dynamic routing handles model selection automatically, and my weekly costs sit around $5.

A few months ago I was skeptical. Now I run my own stack, spend almost nothing, and ship code while I sleep.
