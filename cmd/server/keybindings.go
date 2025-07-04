package server

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Quit     key.Binding
	Help     key.Binding
	Next     key.Binding
	Prev     key.Binding
	Tab      key.Binding
	ShiftTab key.Binding
	Up       key.Binding
	Down     key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
		Help: key.NewBinding(
			key.WithKeys("h", "?"),
			key.WithHelp("h/?", "help"),
		),
		Next: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "next"),
		),
		Prev: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "prev"),
		),
		Tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "next section"),
		),
		ShiftTab: key.NewBinding(
			key.WithKeys("shift+tab"),
			key.WithHelp("shift+tab", "prev section"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "scroll up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "scroll down"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Tab, k.ShiftTab, k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Tab, k.ShiftTab, k.Next, k.Prev},
		{k.Up, k.Down},
		{k.Help, k.Quit},
	}
}
