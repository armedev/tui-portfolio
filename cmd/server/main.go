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

func NewServer(host string, port uint, sshKeyPath string) (*ssh.Server, error) {
	log.Printf("Starting SSH server on %s:%d", host, port)
	log.Printf("Connect with: ssh %s -p %d", host, port)

	return wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithHostKeyPath(sshKeyPath),
		wish.WithMiddleware(
			bubbletea.Middleware(teaHandler),
			logging.Middleware(),
		),
	)

}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// Get terminal dimensions
	pty, _, _ := s.Pty()

	model := NewPortfolioModel(int(pty.Window.Width), int(pty.Window.Height))

	return model, []tea.ProgramOption{
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	}
}
