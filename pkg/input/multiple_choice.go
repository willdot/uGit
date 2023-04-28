package input

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MultiChoiceModel struct {
	choices  []string
	cursor   int
	Selected map[int]string
	question string
}

func (m MultiChoiceModel) Init() tea.Cmd {
	return nil
}

func (m MultiChoiceModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			_, ok := m.Selected[m.cursor]
			if ok {
				delete(m.Selected, m.cursor)
			} else {
				m.Selected[m.cursor] = m.choices[m.cursor]
			}
		case "enter":
			return m, tea.Quit
		}
	default:
	}

	return m, nil
}

func (m MultiChoiceModel) View() string {
	s := fmt.Sprintf("%s\n", m.question)

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.Selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit"

	return s
}

func InitMultiChoiceModel(choices []string, question string) MultiChoiceModel {
	return MultiChoiceModel{
		choices:  choices,
		Selected: make(map[int]string),
		question: question,
	}
}
