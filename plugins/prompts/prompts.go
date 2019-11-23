package prompts

import (
	"strconv"

	"github.com/manifoldco/promptui"
)

// String -
func String(name string) (string, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateEmptyStringInput,
	}

	return prompt.Run()
}

// Integer -
func Integer(name string) (int64, error) {
	prompt := promptui.Prompt{
		Label:    name,
		Validate: validateIntegerNumberInput,
	}

	res, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	parse, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return 0, err
	}
	return parse, nil
}
