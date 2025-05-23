# 🆔 UUIDGen CLI

A tiny and lightning-fast CLI tool to generate one or more **UUIDs (v4)** from your terminal. Built with Go and powered by the official [`github.com/google/uuid`](https://pkg.go.dev/github.com/google/uuid) package.

![Go Version](https://img.shields.io/badge/Go-1.24+-blue?logo=go)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🔢 Generate one or multiple UUID v4s
- ⚡ Super fast — returns instantly
- 🧪 Secure and RFC-compliant
- 🛠️ Clean, minimal, and dependency-light
- 📦 Uses the official Google UUID library

---

## 📦 Installation

### Option 1: Go Install

```bash
go install github.com/hadzicni/uuidgen-cli/cmd/uuidgen@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Manual Build (Windows, Linux, macOS)

#### 🪟 Windows (PowerShell oder CMD)

```powershell
git clone https://github.com/hadzicni/uuidgen-cli.git
cd uuidgen-cli/cmd/uuidgen
go build -o uuidgen.exe
```

#### 🐧 Linux / 🍏 macOS

```bash
git clone https://github.com/hadzicni/uuidgen-cli.git
cd uuidgen-cli/cmd/uuidgen
go build -o uuidgen
```

---

## 🚀 Usage

```bash
uuidgen [flags]
```

### Available Flags

| Flag       | Description                 | Default |
| ---------- | --------------------------- | ------- |
| `-n <int>` | Number of UUIDs to generate | `1`     |

---

## 🔧 Examples

Generate 1 UUID (default):

```bash
uuidgen
```

Generate 5 UUIDs:

```bash
uuidgen -n 5
```

---

## 👨‍💻 Author

Made by **Nikola Hadzic**  
GitHub: [@hadzicni](https://github.com/hadzicni)

---

## 📄 License

This project is licensed under the Apache License 2.0. See the [LICENSE](./LICENSE) file for details.
