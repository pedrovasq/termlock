# Termlock

**Termlock** is a simple terminal-based password manager built in Go, using the Bubble Tea TUI framework. It allows you to browse saved credentials, view details, search entries using fuzzy search, and securely copy passwords to your system clipboard.

---

## Features

- 📂 Browse and manage multiple saved password entries
- 🔍 Live fuzzy search using `/`
- 🔑 Copy passwords to clipboard with `Enter`
- ⬇️ Import CSV files dynamically with `i`
- 🔐 Secure storage using BoltDB with encrypted password fields
- 🎨 Styled terminal interface with a custom theme

---

## Controls

| Key             | Action                         |
| --------------- | ------------------------------ |
| `j` / `↓`       | Move cursor down               |
| `k` / `↑`       | Move cursor up                 |
| `Enter`         | Copy password to clipboard     |
| `/`             | Enter fuzzy search mode        |
| `Esc`           | Exit search or import mode     |
| `i`             | Import entries from a CSV file |
| `q` or `Ctrl+C` | Quit the application           |

---

## CSV Format

Your CSV file should be structured as follows:

```csv
name,url,username,password,note
GitHub,https://github.com,raider54,ghp_abc123,Personal token
```

- `url` can contain multiple sites separated by commas.

---

## Usage

### Clone and Run

```bash
git clone https://github.com/yourusername/termlock.git
cd termlock
go run cmd/termlock/main.go
```

---

## Project Structure

```
termlock/
├── cmd/termlock/        # Entry point
├── internal/
│   ├── models/          # PasswordEntry struct
│   ├── storage/         # DB import/export logic
│   └── styles/          # UI styling
├── go.mod
└── README.md
```

---

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lip Gloss](https://github.com/charmbracelet/lipgloss)
- [Clipboard](https://github.com/atotto/clipboard)
- [BoltDB](https://github.com/etcd-io/bbolt)
- [Fuzzy Search](https://github.com/sahilm/fuzzy)

---

## License

MIT License

---

## Author

Pedro Vasquez

---

**Enjoy using Termlock and keeping your secrets safe!**

