package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/friendsofgo/board-kata/parser"
)

func TestPLainText(t *testing.T) {
	text, _ := parser.Parse("cucu")
	assert.Equal(t, "cucu", text, "plain text not valid")
}

func TestOnlyHashtag(t *testing.T) {
	text, _ := parser.Parse("#golang")
	assert.Equal(t, "<a href='https://fogower.com/hash/golang'>#golang</a>", text)
}

func TestAnotherHashtag(t *testing.T) {
	text, _ := parser.Parse("#goGoPowerRangers")
	assert.Equal(t, "<a href='https://fogower.com/hash/goGoPowerRangers'>#goGoPowerRangers</a>", text)
}

func TestOnlyMention(t *testing.T) {
	text, _ := parser.Parse("@friendsofgotech")
	assert.Equal(t, "<a href='https://fogower.com/friendsofgotech'>@friendsofgotech</a>", text)
}

func TestChangeUrlToLink(t *testing.T) {
	text, _ := parser.Parse("http://www.google.com")
	assert.Equal(t, "<a href='http://www.google.com'>http://www.google.com</a>", text)
}

func TestMultipleWords(t *testing.T) {
	text, _ := parser.Parse("hola http://www.google.com #goGoPowerRangers")
	assert.Equal(t, "hola <a href='http://www.google.com'>http://www.google.com</a> <a href='https://fogower.com/hash/goGoPowerRangers'>#goGoPowerRangers</a>", text)
}

func TestEmptyString(t *testing.T) {
	_, error := parser.Parse("")
	if error == nil {
		t.Errorf("wrong")
	}
}
