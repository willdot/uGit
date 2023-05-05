package cli

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/willdot/uGit/pkg/input"
)

func askUserToSelectOptions(availableOptions []string, message string, addSelectAll bool) ([]string, error) {
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
		return nil, err
	}

	var results []string
	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(input.MultiChoiceModel); ok && len(m.Selected) > 0 {
		for _, v := range m.Selected {

			if v == "**Select all**" {
				return availableOptions, nil
			}
			if v == "**Exit and ignore selections**" {
				return nil, nil
			}

			results = append(results, v)
		}
	}

	return results, nil
}

func askUserToSelectSingleOption(availableOptions []string, message string) (string, error) {
	p := tea.NewProgram(input.InitSingleChoiceModel(availableOptions, message))

	// Run returns the model as a tea.Model.
	model, err := p.Run()
	if err != nil {
		return "", err
	}

	// Assert the final tea.Model to our local model and print the choice.
	m, ok := model.(input.SingleChoiceModel)
	if ok {
		return m.Selected, nil
	}

	return "", nil
}
