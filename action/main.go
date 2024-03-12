package main

import (
	"github.com/jaydee029/issue_assignment_bot/bot"
	"github.com/sethvargo/go-githubactions"
)

func main() {

	Token := githubactions.GetInput("github_token")

	if Token == "" {
		githubactions.Fatalf("missing Github Token")
	}

	val := githubactions.GetInput("commands")
	if val == "" {
		githubactions.Fatalf("missing commands")
	}

	input := bot.Input{
		Token:    Token,
		Commands: val,
	}

	go bot.Bot(input)

}
