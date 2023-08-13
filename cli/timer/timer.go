package timer

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Timer struct {
	current int
	max     int
}

func (tt Timer) Init() tea.Cmd {
	return tick()
}

func (tt Timer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return tt, tea.Quit
		default:
		}
	case tickMsg:
		if tt.max != 0 && tt.current >= tt.max {
			return tt, tea.Quit
		}
		tt.current++
		return tt, tick()
	}
	return tt, nil
}

func (tt Timer) View() string {
	if tt.current >= tt.max {
		return "finish counting.\nprogram end.\n\n"
	}
	return fmt.Sprintf("start counting to %d\n\n%d\n", tt.max, tt.current)
}

func StartTimer() {
	p := tea.NewProgram(Timer{max: 10}, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
