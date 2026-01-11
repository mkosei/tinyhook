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
```
## ⚡ Quick Start

1️⃣ Run server with live tail
```
go run . serve --addr :8080 --tail
```


- --addr : Listen address (default :8080)

- --tail : Real-time output of incoming webhooks in the terminal

This starts the webhook receiver and displays incoming events live, like tail -f logs.

2️⃣ Send a test webhook locally
```
curl -X POST http://localhost:8080/hooks/github \
  -H "X-GitHub-Event: push" \
  -d '{"ref":"main"}'
```

Expected output in terminal (if --tail is enabled):

[15:42:10] github push

3️⃣ Run server without live tail
go run . serve --addr :8080


- Webhooks are still received and stored in memory

- No live output will appear

- Can later inspect events via CLI or implement replay

## 🌐 Testing with real GitHub Webhooks

GitHub cannot directly send webhooks to `localhost`. To test with actual GitHub events:

1. **Install ngrok:** [https://ngrok.com/](https://ngrok.com/)

2. **Expose your local server:**

```bash
ngrok http 8080
```

1. **Copy the Forwarding URL** (e.g., `https://abcd1234.ngrok.io`)

2. **In your GitHub repository:**
   - Go to **Settings → Webhooks → Add webhook**
   - **Payload URL:** `https://abcd1234.ngrok.io/hooks/github`
   - **Content type:** `application/json`
   - Choose which events to trigger (e.g., `push`)
   - Click **Add webhook**

3. **Trigger a webhook:**  
   Now push commits or trigger events in GitHub → they will appear in your tinyhook terminal (if `--tail` is enabled).

---

### 📝 Notes

- ngrok URLs are temporary; they change each time you start ngrok  
- For production-like testing, consider using a fixed domain or tunnel  
- Secrets can be used for verifying GitHub webhook payloads


