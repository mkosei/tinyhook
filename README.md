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
