package server

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Section int

const (
	AboutSection Section = iota
	ExperienceSection
	SkillsSection
	ContactSection
	HelpSection // Add Help as a section but not in navigation
)

type PortfolioModel struct {
	sections       []Section
	currentSection Section
	viewport       viewport.Model
	width          int
	height         int
	styles         *PortfolioStyles
	ready          bool
	animationTick  int

	// Explosion effects only
	particles      []Particle
	effectsEnabled bool
	startTime      time.Time
}

// Particle system for explosions
type Particle struct {
	x, y   float64
	vx, vy float64
	life   float64
	char   string
	color  lipgloss.Color
}

type tickMsg time.Time

func NewPortfolioModel(width, height int) *PortfolioModel {
	vp := viewport.New(width-4, height-8)

	model := &PortfolioModel{
		sections: []Section{
			AboutSection,
			ExperienceSection,
			SkillsSection,
			ContactSection,
			// Help section excluded from normal navigation
		},
		currentSection: AboutSection,
		viewport:       vp,
		width:          width,
		height:         height,
		styles:         NewPortfolioStyles(),
		animationTick:  0,
		particles:      make([]Particle, 0),
		effectsEnabled: true,
		startTime:      time.Now(),
	}

	model.updateContent()
	return model
}

func (m *PortfolioModel) Init() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m *PortfolioModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.viewport.Width = msg.Width - 4
		m.viewport.Height = msg.Height - 8
		m.updateContent()
		if !m.ready {
			m.ready = true
		}

	case tickMsg:
		if m.effectsEnabled {
			m.animationTick++
			m.updateParticles()
		}

		// Update content for sections with real-time data (About section)
		if m.currentSection == AboutSection {
			m.updateContent()
		}

		return m, tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
			return tickMsg(t)
		})

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultKeyMap().Quit):
			return m, tea.Quit
		case key.Matches(msg, DefaultKeyMap().Help):
			// Toggle help section - special navigation
			if m.currentSection == HelpSection {
				m.currentSection = AboutSection // Return to About when leaving help
			} else {
				m.currentSection = HelpSection
			}
			m.updateContent()
			return m, nil
		case msg.String() == "e":
			m.effectsEnabled = !m.effectsEnabled
			return m, nil
		case msg.String() == "x":
			m.addExplosion(m.viewport.Width/2, m.viewport.Height/2)
			return m, nil
		case key.Matches(msg, DefaultKeyMap().Next):
			// Only navigate through normal sections, not help
			if m.currentSection != HelpSection {
				m.nextSection()
				m.updateContent()
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap().Prev):
			// Only navigate through normal sections, not help
			if m.currentSection != HelpSection {
				m.prevSection()
				m.updateContent()
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap().Tab):
			// Only navigate through normal sections, not help
			if m.currentSection != HelpSection {
				m.nextSection()
				m.updateContent()
			}
			return m, nil
		case key.Matches(msg, DefaultKeyMap().ShiftTab):
			// Only navigate through normal sections, not help
			if m.currentSection != HelpSection {
				m.prevSection()
				m.updateContent()
			}
			return m, nil
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m *PortfolioModel) addExplosion(x, y int) {
	chars := []string{"*", "★", "✦", "✧", "●", "◉", "◎", "○", "◯", "◦", "•", "+", "×", "▪", "▫"}

	// Catppuccin Mocha explosion colors
	colors := []lipgloss.Color{
		"#f5c2e7", // Pink
		"#cba6f7", // Mauve
		"#b4befe", // Lavender
		"#89b4fa", // Blue
		"#74c7ec", // Sapphire
		"#89dceb", // Sky
		"#94e2d5", // Teal
		"#a6e3a1", // Green
		"#f9e2af", // Yellow
		"#fab387", // Peach
		"#eba0ac", // Maroon
		"#f38ba8", // Red
		"#f2cdcd", // Flamingo
		"#f5e0dc", // Rosewater
	}

	// Create explosion particles
	particleCount := 20 + rand.Intn(10)
	for range particleCount {
		angle := rand.Float64() * 2 * math.Pi
		speed := 1.0 + rand.Float64()*3.0

		// Add randomness to initial position
		offsetX := rand.Float64()*4 - 2
		offsetY := rand.Float64()*4 - 2

		m.particles = append(m.particles, Particle{
			x:     float64(x) + offsetX,
			y:     float64(y) + offsetY,
			vx:    math.Cos(angle) * speed,
			vy:    math.Sin(angle) * speed,
			life:  0.8 + rand.Float64()*0.4,
			char:  chars[rand.Intn(len(chars))],
			color: colors[rand.Intn(len(colors))],
		})
	}
}

func (m *PortfolioModel) updateParticles() {
	for i := len(m.particles) - 1; i >= 0; i-- {
		p := &m.particles[i]
		p.x += p.vx
		p.y += p.vy
		p.vy += 0.05 // Gravity
		p.vx *= 0.99 // Air resistance
		p.life -= 0.015

		// Remove dead or out-of-bounds particles
		if p.life <= 0 || p.x < -5 || p.x >= float64(m.viewport.Width+5) || p.y >= float64(m.viewport.Height+5) {
			m.particles = append(m.particles[:i], m.particles[i+1:]...)
		}
	}
}

func (m *PortfolioModel) renderParticles() string {
	if len(m.particles) == 0 {
		return ""
	}

	width := m.viewport.Width
	height := m.viewport.Height

	if width <= 0 || height <= 0 {
		return ""
	}

	canvas := make([][]string, height)
	for i := range canvas {
		canvas[i] = make([]string, width)
		for j := range canvas[i] {
			canvas[i][j] = " "
		}
	}

	// Render particles
	for _, p := range m.particles {
		x, y := int(p.x), int(p.y)
		if x >= 0 && x < width && y >= 0 && y < height {
			style := lipgloss.NewStyle().Foreground(p.color)

			// Fade out based on life
			if p.life < 0.3 {
				style = style.Faint(true)
			}

			canvas[y][x] = style.Render(p.char)
		}
	}

	var result strings.Builder
	for _, row := range canvas {
		result.WriteString(strings.Join(row, ""))
		result.WriteString("\n")
	}

	return result.String()
}

func (m *PortfolioModel) overlayParticles(base string) string {
	if len(m.particles) == 0 {
		return base
	}

	particlesOverlay := m.renderParticles()
	baseLines := strings.Split(base, "\n")
	overlayLines := strings.Split(particlesOverlay, "\n")

	maxLines := max(len(overlayLines), len(baseLines))

	result := make([]string, maxLines)

	for i := range maxLines {
		var baseLine, overlayLine string

		if i < len(baseLines) {
			baseLine = baseLines[i]
		}
		if i < len(overlayLines) {
			overlayLine = overlayLines[i]
		}

		// Combine lines - overlay non-space characters
		result[i] = m.combineLine(baseLine, overlayLine)
	}

	return strings.Join(result, "\n")
}

func (m *PortfolioModel) combineLine(base, overlay string) string {
	baseRunes := []rune(base)
	overlayRunes := []rune(overlay)

	maxLen := max(len(overlayRunes), len(baseRunes))

	result := make([]rune, maxLen)

	for i := range maxLen {
		if i < len(overlayRunes) && overlayRunes[i] != ' ' {
			result[i] = overlayRunes[i]
		} else if i < len(baseRunes) {
			result[i] = baseRunes[i]
		} else {
			result[i] = ' '
		}
	}

	return string(result)
}

func (m *PortfolioModel) View() string {
	if !m.ready {
		return m.renderLoadingScreen()
	}

	var content strings.Builder

	// Header
	content.WriteString(m.renderHeader())
	content.WriteString("\n")

	// Navigation tabs
	content.WriteString(m.renderTabs())
	content.WriteString("\n")

	// Main content area with particle overlay
	mainContent := m.viewport.View()
	if len(m.particles) > 0 && m.effectsEnabled {
		mainContent = m.overlayParticles(mainContent)
	}

	content.WriteString(m.styles.ContentBox.Render(mainContent))
	content.WriteString("\n")

	// Footer - simplified since help is now a tab
	content.WriteString(m.renderFooter())

	return content.String()
}

func (m *PortfolioModel) renderLoadingScreen() string {
	loadingFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	frame := loadingFrames[m.animationTick%len(loadingFrames)]

	loading := fmt.Sprintf(`
    %s Loading Portfolio Terminal...
    
    🚀 Initializing components...
    📊 Loading portfolio data...
    🎨 Setting up interface...
    
    Please wait...
    `, frame)

	return m.styles.Header.Render(loading)
}

func (m *PortfolioModel) renderHeader() string {
	title := "🚀 Portfolio Terminal"
	subtitle := "Interactive Developer Portfolio"

	sparkles := []string{"✨", "⭐", "🌟", "💫"}
	sparkle := sparkles[m.animationTick%len(sparkles)]

	titleWithEffects := sparkle + " " + title + " " + sparkle
	headerContent := titleWithEffects + "\n" + m.styles.Subtitle.Render(subtitle)

	statusLine := m.renderStatusIndicators()
	headerContent += "\n" + statusLine

	return m.styles.Header.Render(headerContent)
}

func (m *PortfolioModel) renderStatusIndicators() string {
	var indicators []string

	if m.effectsEnabled {
		indicators = append(indicators, "🎬 Effects: ON")
	} else {
		indicators = append(indicators, "📺 Effects: OFF")
	}

	if len(m.particles) > 0 {
		indicators = append(indicators, fmt.Sprintf("⚡ Particles: %d", len(m.particles)))
	}

	uptime := time.Since(m.startTime).Truncate(time.Second)
	indicators = append(indicators, fmt.Sprintf("⏱️ Uptime: %v", uptime))

	return m.styles.FooterLeft.Render(strings.Join(indicators, " │ "))
}

func (m *PortfolioModel) renderTabs() string {
	var leftTabs []string
	var rightTabs []string

	sectionNames := map[Section]string{
		AboutSection:      "About",
		ExperienceSection: "Experience",
		SkillsSection:     "Skills",
		ContactSection:    "Contact",
		HelpSection:       "Help",
	}

	icons := map[Section]string{
		AboutSection:      "👋",
		ExperienceSection: "💼",
		SkillsSection:     "🚀",
		ContactSection:    "📞",
		HelpSection:       "❓",
	}

	// Render normal navigation tabs on the left
	for _, section := range m.sections {
		name := sectionNames[section]
		tabText := icons[section] + " " + name

		if section == m.currentSection {
			leftTabs = append(leftTabs, m.styles.ActiveTab.Render(tabText))
		} else {
			leftTabs = append(leftTabs, m.styles.InactiveTab.Render(tabText))
		}
	}

	// Render help tab on the right
	helpTabText := icons[HelpSection] + " " + sectionNames[HelpSection]
	if m.currentSection == HelpSection {
		rightTabs = append(rightTabs, m.styles.ActiveTab.Render(helpTabText))
	} else {
		rightTabs = append(rightTabs, m.styles.InactiveTab.Render(helpTabText))
	}

	// Join left tabs
	leftSide := lipgloss.JoinHorizontal(lipgloss.Bottom, leftTabs...)
	rightSide := lipgloss.JoinHorizontal(lipgloss.Bottom, rightTabs...)

	// Fallback if not enough space
	return leftSide + " " + rightSide
}

func (m *PortfolioModel) renderFooter() string {
	help := "Tab/Shift+Tab: Navigate • h: Help • e: Toggle Effects • x: Explosion • q: Quit"

	status := ("💻 Portfolio on Interactive Terminal 🎮")

	left := m.styles.FooterLeft.Render(help)
	right := m.styles.FooterRight.Render(status)

	gap := m.width - lipgloss.Width(left) - lipgloss.Width(right)
	if gap > 0 {
		return left + strings.Repeat(" ", gap) + right
	}
	return left
}

func (m *PortfolioModel) nextSection() {
	current := int(m.currentSection)
	m.currentSection = Section((current + 1) % len(m.sections))
}

func (m *PortfolioModel) prevSection() {
	current := int(m.currentSection)
	if current == 0 {
		m.currentSection = Section(len(m.sections) - 1)
	} else {
		m.currentSection = Section(current - 1)
	}
}

func (m *PortfolioModel) updateContent() {
	content := m.getSectionContent(m.currentSection)
	m.viewport.SetContent(content)
	m.viewport.GotoTop()
}

func (m *PortfolioModel) getSectionContent(section Section) string {
	switch section {
	case AboutSection:
		return m.renderAbout()
	case ExperienceSection:
		return m.renderExperience()
	case SkillsSection:
		return m.renderSkills()
	case ContactSection:
		return m.renderContact()
	case HelpSection:
		return m.renderHelp()
	default:
		return "Section not found"
	}
}

func (m *PortfolioModel) renderHelp() string {
	help := `
🎮 Navigation & Controls:
  Tab              Next section
  Shift+Tab        Previous section  
  ↑ ↓              Scroll content
  h                Toggle help (this section)
  q / Ctrl+C       Quit

🎬 Effects:
  e                Toggle particle effects on/off
  x                Trigger explosion at center

📋 Sections:
  👋 About         Personal introduction, current time, and tech facts
  💼 Experience    Professional work history
  🚀 Skills        Technical expertise and proficiency
  📞 Contact       Get in touch information

💡 Tips:
  • Press 'h' anytime to access help
  • This runs entirely in your terminal!

🌈 Theme:
  • Using beautiful Catppuccin Mocha color palette

🚀 Getting Started:
  • Use Tab/Shift+Tab to navigate between main sections
  • Press 'h' to return here anytime
  • Try 'x' for particle explosions
  • Toggle effects with 'e' if needed
`

	return m.styles.ContentText.Render(help)
}
