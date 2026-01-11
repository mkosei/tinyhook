# tinyhook

A lightweight, local-first webhook receiver built with **Go + Echo**.  
Designed for **development, debugging, and learning**, without any frontend or cloud dependency.

> Think of it as a minimal, self-hosted alternative to webhook.site — focused on clarity, extensibility, and clean architecture.

---

## ✨ Features

- 🚀 Receive webhooks locally from any provider
- 🧩 Provider-agnostic (`/hooks/:provider`)
- 🧠 Clean separation using Echo **Middleware**
- 📦 Capture raw request body & headers
- 🕒 Timestamped event records
- 🔧 No frontend, no database, no cloud required
- 📚 Great for learning real-world Go backend patterns

---

## 📦 Use Cases

- Test webhooks from GitHub, Stripe, Slack, etc.
- Inspect raw webhook payloads during development
- Learn Echo middleware & context patterns
- Build your own webhook debugging tools
- Use as a base for replay, filtering, or persistence features

---

## 🏗 Architecture Overview

```text
Request
  ↓
[ Echo Middleware ]
  - Read raw body
  - Collect headers
  - Extract provider
  - Build Event model
  - Store in Context
  ↓
[ Handler ]
  - Minimal logic
  - Just consume Event
  ↓
Response

## ⚡ Quick Start

### 1️⃣ Run server with live tail

```bash
go run . serve --addr :8080 --tail

- `--addr` : Listen address (default `:8080`)  
- `--tail` : Real-time output of incoming webhooks in the terminal  

This starts the webhook receiver **and** displays incoming events live, like `tail -f` logs.

---

### 2️⃣ Send a test webhook

```bash
curl -X POST http://localhost:8080/hooks/github \
  -H "X-GitHub-Event: push" \
  -d '{"ref":"main"}'

### 3️⃣ Run server without live tail

```bash
go run . serve --addr :8080

- Webhooks are still received and stored in memory
- No live output will appear
- Can later inspect events via CLI or implement replay


