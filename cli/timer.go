package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type tryTea struct {
	current int
	max     int
}
type tickMsg time.Time

var count int

func (tt tryTea) Init() tea.Cmd {
	return tick()
}

func (tt tryTea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (tt tryTea) View() string {
	if tt.current >= tt.max {
		return "finish counting.\nprogram end.\n\n"
	}
	s := fmt.Sprintf("start counting to %d\n\n%d\n", tt.max, tt.current)
	return s
}

func runTryTea() {
	p := tea.NewProgram(tryTea{max: 10})
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
