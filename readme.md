# 🚀 Terminal Portfolio

A beautiful, interactive terminal-based portfolio application built with Go and Charm libraries. Features the gorgeous Catppuccin Mocha color scheme and particle explosion effects.

## ✨ Features

- **Interactive Portfolio** - Navigate through About, Experience, Skills, Projects, Contact, and Live Demo sections
- **Particle Explosions** - Press `x` for colorful fireworks using physics-based particles
- **Beautiful Theme** - Catppuccin Mocha color palette for comfortable viewing
- **SSH Server** - Access remotely via SSH or run locally
- **Live Animations** - Real-time clock, progress bars, and system stats
- **Responsive Design** - Adapts to any terminal size

## 🎮 Controls

| Key | Action |
|-----|--------|
| `Tab` / `Shift+Tab` | Navigate sections |
| `h` | Toggle help (right-aligned tab) |
| `x` | Trigger particle explosion |
| `e` | Toggle effects on/off |
| `↑` `↓` | Scroll content |
| `q` | Quit |

## 🚀 Quick Start

```bash
# Clone and setup
git clone <repo-url>
cd portfolio-terminal

# Install dependencies
go mod tidy

# Run locally
go run ./cmd
```

## 📁 Project Structure

```
├── main.go           # SSH server and entry point
├── model.go          # UI logic and particle effects
├── content.go        # Portfolio content
├── styles.go         # Catppuccin Mocha theme
├── keybindings.go    # Key mappings
└── go.mod           # Dependencies
```

## 🎨 Customization

- **Content**: Edit `content.go` to update your information
- **Colors**: Modify `styles.go` for different themes
- **Effects**: Adjust particle settings in `model.go`

## 🛠️ Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components
- [Wish](https://github.com/charmbracelet/wish) - SSH server

## 📝 License

MIT License - feel free to use for your own portfolio!

---

**Built with ❤️ using Go and Charm libraries**
