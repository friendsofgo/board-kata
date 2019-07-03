package parser

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	baseURL    = "https://fogo-parser.dev/"
	hashTagURL = "https://fogo-parser.dev/hash/"
)

// Parse fomating the text into a valid output text
func Parse(msg string) (output string, err error) {

	if msg == "" {
		return "", errors.New("empty string")
	}
	output = ""
	words := strings.Split(msg, " ")
	for i := 0; i < len(words); i++ {
		output += parseWord(words[i]) + " "
	}
	return output[:len(output)-1], nil
}

func parseWord(msg string) (output string) {
	re := regexp.MustCompile(`@`)
	if re.MatchString(msg) {
		stringifiedHash := msg[1:]
		return fmt.Sprintf("<a href='https://fogower.com/%s'>@%s</a>", stringifiedHash, stringifiedHash)
	}

	re = regexp.MustCompile(`#`)
	if re.MatchString(msg) {
		stringifiedHash := msg[1:]
		return fmt.Sprintf("<a href='https://fogower.com/hash/%s'>#%s</a>", stringifiedHash, stringifiedHash)
	}

	re = regexp.MustCompile(`https*`)
	if re.MatchString(msg) {
		return fmt.Sprintf("<a href='%s'>%s</a>", msg, msg)
	}

	return msg
}
