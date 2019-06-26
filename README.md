# Board Kata

A simple kata to emulate Twitter text transformation behaviour.

## Exercise details

This excersise want to simulate the Twitter text transformation behaviour, it does it means that you'll recieve a `csv file` with multiples texts, you need to have in consideration that these texts could contains "," into them.

You need to create a `parse` function who is responsible of all transformations.

```go
// Parse fomating the text into a valid output text
func Parse(msg string) (output string) {
	// code here
	return output
}
```

The rules of transformations are:

* Each url must be parsed into a html form -> https://friendsofgo.tech (`<a href="https://friendsofgo.tech">https://friendsofgo.tech</a>`)
* Each mention must be parsed into a author link -> @FriendsOfGoTech (`<a href="https://fogo-parser.dev/FriendsOfGoTech">@FriendsOfGoTech</a>`)
* Each hashtag must be parsed into a hashtag link -> #friendsofgo (`<a href="https://fogo-parser.dev/hash/friendsofgo">#friendsofgo</a>`)

Furthermore you'll have to detect on you can apply [errors by behaviour](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully) and create the test for the package `parser`

## Boilerplate

In this excersise we give you the `csv reader` that read the `data/input.csv` file and we give you the `main.go` and you need the complete with all necessary

```go
func main() {
	msg, err := board.ReadInput("data/input.csv")
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Parse each message
	fmt.Println(msg)

	//TODO: Print output into an html file

	fmt.Println("done!")
}
```

## Run

First of all we use `go modules` for this boilerplate

```sh
$ go mod tidy
```

And then you can run your application

```sh
$ go run cmd/board/main.go
```

## Bonus
* Use only errors by behaviour
* Use table driven tests
* Get >90% coverage into `parser` package
* Print output of all files into a `html file`

