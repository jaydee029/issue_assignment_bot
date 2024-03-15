package main

import (
	"github.com/sethvargo/go-githubactions"
)

func main() {
	//ctx := context.Background()

	action := githubactions.New()

	Token := action.GetInput("github_token")

	if Token == "" {
		githubactions.Fatalf("missing Github Token")
	}

	val := action.GetInput("commands")
	if val == "" {
		githubactions.Fatalf("missing commands")
	}

	action.Infof(Token, val)
	/*
		input := bot.Input{
			Token:    Token,
			Commands: val,
		}

		action.Infof("hello")
		//go bot.Bot(input)
	*/
}
