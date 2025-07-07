package server

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

// Server configuration
type ServerConfig struct {
	Host       string
	Port       uint
	SSHKeyPath string
	DataLoader *DataLoader
}

func NewServer(host string, port uint, sshKeyPath, dataPath string) (*ssh.Server, error) {
	log.Printf("Starting SSH server on %s:%d", host, port)
	log.Printf("Connect with: ssh %s -p %d", host, port)
	log.Printf("Loading portfolio data from: %s", dataPath)

	// Initialize data loader
	dataLoader := NewDataLoader(dataPath)
	if err := dataLoader.LoadData(); err != nil {
		return nil, fmt.Errorf("Warning: Failed to load portfolio data: %v", err)
	}

	// Validate loaded data
	if dataLoader.IsLoaded() {
		if err := dataLoader.ValidateData(); err != nil {
			return nil, fmt.Errorf("Warning: Data validation failed: %v", err)
		}
	}

	// Create a server config to pass around
	config := &ServerConfig{
		Host:       host,
		Port:       port,
		SSHKeyPath: sshKeyPath,
		DataLoader: dataLoader,
	}

	return wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(sshKeyPath),
		wish.WithMiddleware(
			bubbletea.Middleware(func(s ssh.Session) (tea.Model, []tea.ProgramOption) {
				return teaHandler(s, config)
			}),
			logging.Middleware(),
		),
	)
}

func teaHandler(s ssh.Session, config *ServerConfig) (tea.Model, []tea.ProgramOption) {
	// Get terminal dimensions
	pty, _, _ := s.Pty()

	model := NewPortfolioModel(int(pty.Window.Width), int(pty.Window.Height), config.DataLoader)

	return model, []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}
}
