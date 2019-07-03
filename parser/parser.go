package parser

import (
	"regexp"
	"strings"
)

const (
	baseURL    = "https://fogo-parser.dev/"
	hashTagURL = "https://fogo-parser.dev/hash/"
)

var orderedRegexps = []string{
	`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&amp;:/~\+#]*[\w\-\@?^=%&amp;/~\+#])?`,
	`(?:#)([a-zA-Z\d]+)?`,
	`(?:@)([a-zA-Z\d]+)?`,
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Parse fomating the text into a valid output text
func Parse(msg string) (output string) {
	for i, value := range orderedRegexps {
		re := regexp.MustCompile(value)
		matches := re.FindAllString(msg, -1)
		uniqueMatches := unique(matches)
		for _, match := range uniqueMatches {
			if i == 1 {
				msg = strings.Replace(msg, match, "<a href=\""+hashTagURL+"\">"+match+"</a>", -1)
			} else if i == 2 {
				msg = strings.Replace(msg, match, "<a href=\""+baseURL+"\">"+match+"</a>", -1)
			} else {
				msg = strings.Replace(msg, match, "<a href=\""+match+"\">"+match+"</a>", -1)
			}
		}
	}
	return msg
}
