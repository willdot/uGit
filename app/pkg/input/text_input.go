package input

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TextInputModel struct {
	TextInput textinput.Model
	Err       error
}

func InitialTextInputModel(question string) TextInputModel {
	ti := textinput.New()
	ti.Placeholder = question
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return TextInputModel{
		TextInput: ti,
		Err:       nil,
	}
}

func (m TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case error:
		m.Err = msg
		return m, nil
	default:
	}

	m.TextInput, cmd = m.TextInput.Update(msg)

	return m, cmd
}

func (m TextInputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.TextInput.View(),
		"(esc to quit)",
	)
}
