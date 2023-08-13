package session

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SessionModel struct {
}

func initSession() *SessionModel {
	return &SessionModel{}
}

func (s *SessionModel) Init() tea.Cmd {
	return nil
}

func (s *SessionModel) Update(m tea.Msg) (tea.Model, tea.Cmd) {
	lipgloss.NewStyle()
	return s, nil
}

func (s *SessionModel) View() string {
	return ""
}

func SaveSession() {
	p := tea.NewProgram(&SessionModel{})
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
