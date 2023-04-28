package root

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/willdot/uGit/app/pkg/input"
)

func askUserToSelectOptions(availableOptions []string, message string, addSelectAll bool) []string {
	options := make([]string, 0, len(availableOptions))
	for _, opt := range availableOptions {
		options = append(options, opt)
	}
	options = append([]string{"**Exit and ignore selections**"}, options...)

	if addSelectAll {
		options = append([]string{"**Select all**"}, options...)
	}

	p := tea.NewProgram(input.InitMultiChoiceModel(options, message))

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	var results []string
	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(input.MultiChoiceModel); ok && len(m.Selected) > 0 {
		for _, v := range m.Selected {

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

func askUserToSelectSingleOption(availableOptions []string) string {
	p := tea.NewProgram(input.InitSingleChoiceModel(availableOptions))

	// Run returns the model as a tea.Model.
	model, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	m, ok := model.(input.SingleChoiceModel)
	if ok {
		return m.Selected
	}

	return ""
}
