package root

import survey "gopkg.in/AlecAivazis/survey.v1"

func askUserToSelectOptions(availableOptions []string, message string, addSelectAll bool) []string {

	options := append([]string{"**Exit and ignore selections**"}, availableOptions...)

	if addSelectAll {
		options = append([]string{"**Select all**"}, options...)
	}

	result := []string{}
	prompt := &survey.MultiSelect{
		Message: message,
		Options: options,
	}

	survey.AskOne(prompt, &result, nil)

	for i := 0; i < len(result); i++ {
		if result[i] == "**Select all**" {
			return availableOptions
		}
		if result[i] == "**Exit and ignore selections**" {
			return nil
		}
	}

	return result
}
