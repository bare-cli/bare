package ui

import (
	"bare/styles"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

func PromptString(label string, defval string) string {
	fmt.Print(styles.PromptStyle.Render(label + " [" + defval + "]" + " > "))
	input := bufio.NewReader(os.Stdin)
	line, err := input.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSuffix(line, "\n")
}

func PromptSelect(label string, items []string) string {

	prompt := promptui.Select{
		Label:        label,
		Items:        items,
		HideHelp:     true,
		HideSelected: true,
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatal("Error encountered in Prompt.")
	}
	fmt.Printf("Template : %s\n", result)
	return result
}
