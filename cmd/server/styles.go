package server

import "github.com/charmbracelet/lipgloss"

type PortfolioStyles struct {
	Header             lipgloss.Style
	Subtitle           lipgloss.Style
	ActiveTab          lipgloss.Style
	InactiveTab        lipgloss.Style
	ContentBox         lipgloss.Style
	FooterLeft         lipgloss.Style
	FooterRight        lipgloss.Style
	HelpBox            lipgloss.Style
	SectionTitle       lipgloss.Style
	ContentText        lipgloss.Style
	ExperienceTitle    lipgloss.Style
	ExperienceMeta     lipgloss.Style
	ExperienceDetail   lipgloss.Style
	SkillCategory      lipgloss.Style
	SkillBar           lipgloss.Style
	ProjectTitle       lipgloss.Style
	ProjectDescription lipgloss.Style
	ProjectLabel       lipgloss.Style
	LiveTitle          lipgloss.Style
	LiveValue          lipgloss.Style
	LiveSubtitle       lipgloss.Style
	StatsBox           lipgloss.Style
	FactBox            lipgloss.Style
	AsciiArt           lipgloss.Style
}

func NewPortfolioStyles() *PortfolioStyles {
	// Catppuccin Mocha color palette
	var (
		// Base colors
		base   = lipgloss.Color("#1e1e2e") // Dark background
		mantle = lipgloss.Color("#181825") // Darker background

		// Surface colors
		surface0 = lipgloss.Color("#313244") // Light surface
		surface1 = lipgloss.Color("#45475a") // Medium surface

		// Text colors
		text     = lipgloss.Color("#cdd6f4") // Main text
		subtext1 = lipgloss.Color("#bac2de") // Secondary text
		subtext0 = lipgloss.Color("#a6adc8") // Muted text
		overlay2 = lipgloss.Color("#9399b2") // Overlays
		overlay1 = lipgloss.Color("#7f849c") // Darker overlays

		// Accent colors
		lavender = lipgloss.Color("#b4befe") // Light purple
		blue     = lipgloss.Color("#89b4fa") // Blue
		sapphire = lipgloss.Color("#74c7ec") // Light blue
		sky      = lipgloss.Color("#89dceb") // Sky blue
		teal     = lipgloss.Color("#94e2d5") // Teal
		green    = lipgloss.Color("#a6e3a1") // Green
		yellow   = lipgloss.Color("#f9e2af") // Yellow
		peach    = lipgloss.Color("#fab387") // Orange
		mauve    = lipgloss.Color("#cba6f7") // Purple
		pink     = lipgloss.Color("#f5c2e7") // Pink
		flamingo = lipgloss.Color("#f2cdcd") // Light pink
	)

	return &PortfolioStyles{
		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(text).
			Background(base).
			Padding(1, 2).
			MarginBottom(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lavender).
			Align(lipgloss.Center),

		Subtitle: lipgloss.NewStyle().
			Foreground(subtext1).
			Italic(true).
			Align(lipgloss.Center),

		ActiveTab: lipgloss.NewStyle().
			Bold(true).
			Foreground(base).
			Background(mauve).
			Padding(0, 2).
			MarginRight(1),

		InactiveTab: lipgloss.NewStyle().
			Foreground(overlay1).
			Background(surface0).
			Padding(0, 2).
			MarginRight(1),

		ContentBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(surface1).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1),

		FooterLeft: lipgloss.NewStyle().
			Foreground(overlay2).
			Background(surface0).
			Padding(0, 1),

		FooterRight: lipgloss.NewStyle().
			Foreground(sapphire).
			Background(surface0).
			Padding(0, 1).
			Bold(true),

		HelpBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(yellow).
			Background(base).
			Foreground(text).
			Padding(1, 2).
			MarginTop(1),

		SectionTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(mauve).
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(mauve).
			PaddingBottom(1).
			MarginBottom(1),

		ContentText: lipgloss.NewStyle().
			Foreground(text).
			MarginBottom(1),

		ExperienceTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(blue).
			Background(surface0).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(blue),

		ExperienceMeta: lipgloss.NewStyle().
			Foreground(subtext0).
			Italic(true),

		ExperienceDetail: lipgloss.NewStyle().
			Foreground(text).
			MarginLeft(2),

		SkillCategory: lipgloss.NewStyle().
			Bold(true).
			Foreground(teal).
			MarginBottom(1).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(teal).
			PaddingBottom(1),

		SkillBar: lipgloss.NewStyle().
			Foreground(green),

		ProjectTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(peach).
			Background(surface0).
			Padding(0, 1).
			MarginBottom(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(peach),

		ProjectDescription: lipgloss.NewStyle().
			Foreground(subtext1).
			Italic(true),

		ProjectLabel: lipgloss.NewStyle().
			Bold(true).
			Foreground(sky),

		LiveTitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(sapphire).
			Background(surface0).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(sapphire),

		LiveValue: lipgloss.NewStyle().
			Bold(true).
			Foreground(yellow).
			Background(mantle).
			Padding(0, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(yellow).
			Align(lipgloss.Center),

		LiveSubtitle: lipgloss.NewStyle().
			Foreground(subtext0).
			Italic(true).
			Align(lipgloss.Center),

		StatsBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(green).
			Background(surface0).
			Foreground(text).
			Padding(1, 2),

		FactBox: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(pink).
			Background(surface0).
			Foreground(text).
			Padding(1, 2).
			Italic(true),

		AsciiArt: lipgloss.NewStyle().
			Foreground(flamingo).
			Align(lipgloss.Center),
	}
}
