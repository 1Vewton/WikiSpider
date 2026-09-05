package operations

import (
	"fmt"
	"strings"
)

// YesOrNo asks user yes or no question
func YesOrNo(
	defaultChoice bool,
	prompt string,
) bool {
	var yesOrNoPrompt string
	if defaultChoice {
		yesOrNoPrompt = " (Y/n) "
	} else {
		yesOrNoPrompt = " (y/N) "
	}
	fmt.Print(prompt)
	fmt.Println(yesOrNoPrompt)
	var answer string
	fmt.Scan(&answer)
	answerNoSpace := strings.TrimSpace(answer)
	if answerNoSpace == "Y" || answerNoSpace == "y" {
		return true
	}
	if answerNoSpace == "N" || answerNoSpace == "n" {
		return false
	}
	return defaultChoice
}
