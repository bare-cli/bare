package prompt

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/manifoldco/promptui"
)

func Prompt() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func PromptString(label string) string {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatal("Error encountered in Prompt.")
	}

	return result
}
