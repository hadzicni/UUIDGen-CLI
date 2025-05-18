# 🔑 UUIDGen CLI

A fast, minimal and developer-friendly command-line tool to generate UUIDs — built with Go.

![Go Version](https://img.shields.io/badge/Go-1.20+-blue?logo=go)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Platform](https://img.shields.io/badge/platform-macOS%20%7C%20Linux%20%7C%20Windows-lightgrey)

---

## ✨ Features

- 🔁 Generate one or multiple UUIDs  
- 🧪 Supports versions 1, 3, 4, and 5  
- 📋 Copy UUIDs to clipboard (optional)  
- 📦 Lightweight & dependency-free (only uses the Go standard library)  
- 🛠️ Easy to install and use on any platform  

---

## 📦 Installation

### Option 1: Go Install (recommended)

```bash
go install github.com/nikolahadzic/uuidgen-cli@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

### Option 2: Build Manually

Clone the repo and build it:

```bash
git clone https://github.com/nikolahadzic/uuidgen-cli.git
cd uuidgen-cli
go build -o uuidgen ./cmd/uuidgen
```

---

## 🚀 Usage

```bash
uuidgen [flags]
```

### Available Flags

| Flag             | Description                                 | Example                         |
|------------------|---------------------------------------------|----------------------------------|
| `--version`      | Show version information                    | `uuidgen --version`             |
| `--help`         | Show help message                           | `uuidgen --help`                |
| `--count`, `-c`  | Number of UUIDs to generate (default: 1)    | `uuidgen -c 5`                  |
| `--type`, `-t`   | UUID type: `v1`, `v3`, `v4`, `v5` (default: v4) | `uuidgen -t v1`              |

---

## 🔧 Examples

Generate a single UUID:

```bash
uuidgen
```

Generate 5 UUIDs of version 1:

```bash
uuidgen -c 5 -t v1
```

Get help:

```bash
uuidgen --help
```

---

## 🧪 Development

To run the CLI locally during development:

```bash
go run ./cmd/uuidgen
```

To run tests:

```bash
go test ./...
```

---

## 📁 Project Structure

```
.
├── cmd/
│   └── uuidgen/       # CLI command implementation
├── go.mod             # Go module definition
├── LICENSE            # License file
└── README.md          # Project readme
```

---

## 👨‍💻 Author

Made with ❤️ by **Nikola Hadzic**

- GitHub: [@nikolahadzic](https://github.com/nikolahadzic)
- Website: _coming soon_

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
