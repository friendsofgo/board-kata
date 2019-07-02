package parser_test

import (
	"testing"

	"github.com/friendsofgo/board-kata/parser"
)

func TestParser(t *testing.T) {

	tests := map[string]struct{ textToParse, expected string }{
		"plainText":          {"cucu", "cucu"},
		"testOnlyHashtag":    {"#golang", "<a href='https://fogower.com/hash/golang'>#golang</a>"},
		"testAnotherHashtag": {"#goGoPowerRangers", "<a href='https://fogower.com/hash/goGoPowerRangers'>#goGoPowerRangers</a>"},
		"testOnlyMention":    {"@friendsofgotech", "<a href='https://fogower.com/friendsofgotech'>@friendsofgotech</a>"},
		"testUrlToLink":      {"http://www.google.com", "<a href='http://www.google.com'>http://www.google.com</a>"},
		"testHttpsUrlToLink": {"https://www.google.com", "<a href='https://www.google.com'>https://www.google.com</a>"},
		"testMultipleWords":  {"hola http://www.google.com #goGoPowerRangers", "hola <a href='http://www.google.com'>http://www.google.com</a> <a href='https://fogower.com/hash/goGoPowerRangers'>#goGoPowerRangers</a>"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			text, _ := parser.Parse(tc.textToParse)
			if text != tc.expected {
				t.Fatalf("wrong parse expected %s got %s", tc.expected, text)
			}
		})
	}
}

func TestEmptyString(t *testing.T) {
	_, error := parser.Parse("")
	if error == nil {
		t.Errorf("wrong")
	}
}
