package root

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// func askUserToSelectOptions(availableOptions []string, message string, addSelectAll bool) []string {

// 	options := append([]string{"**Exit and ignore selections**"}, availableOptions...)

// 	if addSelectAll {
// 		options = append([]string{"**Select all**"}, options...)
// 	}

// 	result := []string{}
// 	prompt := &survey.MultiSelect{
// 		Message: message,
// 		Options: options,
// 	}

// 	survey.AskOne(prompt, &result, nil)

// 	for i := 0; i < len(result); i++ {
// 		if result[i] == "**Select all**" {
// 			return availableOptions
// 		}
// 		if result[i] == "**Exit and ignore selections**" {
// 			return nil
// 		}
// 	}

// 	return result
// }

type model struct {
	choices  []string
	cursor   int
	selected map[int]string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = m.choices[m.cursor]
			}
		case "enter":
			return m, tea.Quit
		}
	default:
		fmt.Println("yo")
	}

	return m, nil
}

func (m model) View() string {
	s := "What would you like?\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.|n"

	return s
}

func initModel(choices []string) model {
	return model{
		choices:  choices,
		selected: make(map[int]string),
	}
}

func askUserToSelectOptions(availableOptions []string, message string, addSelectAll bool) []string {
	options := make([]string, 0, len(availableOptions))
	for _, opt := range availableOptions {
		options = append(options, opt)
	}
	options = append([]string{"**Exit and ignore selections**"}, options...)

	if addSelectAll {
		options = append([]string{"**Select all**"}, options...)
	}

	p := tea.NewProgram(initModel(options))

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	var results []string
	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && len(m.selected) > 0 {
		for _, v := range m.selected {

			if v == "**Select all**" {
				return availableOptions
			}
			if v == "**Exit and ignore selections**" {
				return nil
			}

			results = append(results, v)
		}
	}

	return results
}
