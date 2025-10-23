package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n') // read input until \n

	text = strings.TrimSuffix(text, "\n") // remove \n at the end of the text
	text = strings.TrimSuffix(text, "\r") // on windows there's an \r at the end

	if error != nil {
		return ""
	}

	return text
}
