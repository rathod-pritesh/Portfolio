# AGENTS.md - AI Coding Agent Guidelines for Pritesh Rathod's Portfolio

Welcome, AI coding agent! This document guides you on how to understand, develop, test, and contribute to this repository.

## Project Overview

This repository hosts the personal developer portfolio and backend services for **Pritesh Rathod** (Python & Go Backend Developer).

- **Frontend (`/frontend`)**: SvelteKit 2 + Vite + TailwindCSS deployed on Vercel (`https://priteshrathod.vercel.app`).
- **Backend (`/portfolio-backend`)**: Golang + Gin framework + MongoDB Atlas deployed on Render (`https://portfolio-i5x9.onrender.com`).

---

## Repository Structure

```
NewPortfolio/
├── frontend/                     # SvelteKit client application
│   ├── src/
│   │   ├── app.html              # HTML shell with meta tags & JSON-LD
│   │   ├── lib/                  # Reusable Svelte components & configs
│   │   └── routes/               # SvelteKit routes & API endpoints
│   ├── static/                   # Static assets, specifications, discovery manifests
│   │   ├── .well-known/          # ARD, MCP, A2A, Agent Skills catalogs
│   │   ├── openapi.json          # OpenAPI 3.1.0 API specification
│   │   ├── auth.md               # WorkOS-compliant agent auth guide
│   │   ├── llms.txt              # Standard LLM navigation index
│   │   ├── llms-full.txt         # Full developer portfolio details
│   │   ├── robots.txt            # AI crawler & answer engine policies
│   │   └── sitemap.xml           # XML sitemap with lastmod timestamps
│   ├── package.json
│   └── svelte.config.js
├── portfolio-backend/            # Golang REST API backend
│   ├── controllers/              # Gin route handlers
│   ├── models/                   # Data models
│   ├── db/                       # MongoDB database client
│   ├── main.go                   # Server entrypoint
│   └── go.mod
├── AGENTS.md                     # This file
├── .cursorrules                  # Cursor agent instructions
├── SKILL.md                      # Agent skill definition
└── plugin.json                   # Agent Plugin manifest (agent-plugins.org)
```

---

## Development Workflows

### 1. Frontend (SvelteKit)
```bash
cd frontend
npm install
npm run dev        # Runs local dev server on http://localhost:5173
npm run build      # Builds production bundle
npm run preview    # Previews production build
```

### 2. Backend (Golang)
```bash
cd portfolio-backend
go mod tidy
go run main.go     # Runs API server on http://localhost:8080
```

---

## Agent Conventions & Rules

1. **Agent Discovery Files**:
   - Maintain `robots.txt` AI crawler rules. Always ensure answer engine bots (`GPTBot`, `ClaudeBot`, `PerplexityBot`, `OAI-SearchBot`) remain allowed.
   - Do not break standard paths: `/.well-known/ard.json`, `/.well-known/agent-skills/index.json`, `/.well-known/mcp/server-card.json`, `/openapi.json`, `/auth.md`, and `/llms.txt`.

2. **Code Style**:
   - Frontend: Modern Svelte 5 / SvelteKit syntax, TailwindCSS v4 utility classes.
   - Backend: Idiomatic Go, error handling with explicit `if err != nil`, Gin JSON response formatting.

3. **API Contracts**:
   - Keep `openapi.json` and `openapi.yaml` in sync whenever backend endpoints in `portfolio-backend/main.go` change.
