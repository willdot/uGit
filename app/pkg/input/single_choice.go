package input

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type SingleChoiceModel struct {
	choices  []string
	cursor   int
	Selected string
}

func (m SingleChoiceModel) Init() tea.Cmd {
	return nil
}

func (m SingleChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case " ":
			m.Selected = m.choices[m.cursor]
		case "enter":
			return m, tea.Quit
		}
	default:
	}

	return m, nil
}

func (m SingleChoiceModel) View() string {
	s := strings.Builder{}
	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func InitSingleChoiceModel(choices []string) SingleChoiceModel {
	return SingleChoiceModel{
		choices: choices,
	}
}
