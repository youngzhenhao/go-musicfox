package ui

import (
	tea "github.com/anhoder/bubbletea"
	"time"
)

// update main ui
func updateMainUI(msg tea.Msg, m NeteaseModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		return keyMsgHandle(msg, m)

	case tickMainUIMsg:
		// every second update ui
		return m, tickMainUI(time.Second)

	}

	return m, nil
}

// get main ui view
func mainUIView(m NeteaseModel) string {
	return "test"
}

func keyMsgHandle(msg tea.KeyMsg, m NeteaseModel) (tea.Model, tea.Cmd) {
	return m, nil
}