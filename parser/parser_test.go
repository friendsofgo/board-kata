package parser_test

import (
	"testing"
	"github.com/friendsofgo/board-kata/parser"
)

func TestParser(t * testing.T) {
	tests := map[string]struct{ code, expected string } {
		"Mention example": {"@CodelyTV", "<a href=\"https://fogo-parser.dev/\">@CodelyTV</a>"},
		"Hastag example": {"#golang", "<a href=\"https://fogo-parser.dev/hash/\">#golang</a>"},
		"HTTP example": {"(link: http://uflare.io)", "(link: <a href=\"http://uflare.io\">http://uflare.io</a>)"},
	}

	for name, tc := range tests {
		t.Run(name, func (t *testing.T) {
			got := parser.Parse(tc.code)
			if got != tc.expected {
				t.Fatalf(
					"wrong codification - expected: %s, got: %s\n",
					tc.expected,
					got,
				)
			}
		})
	}
}