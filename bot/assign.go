package bot

import (
	"github.com/sethvargo/go-githubactions"
)

type Input struct {
	Token    string
	Commands string
}

func Bot(input Input) {

	githubactions.Infof("Hello")

	//event := githubactions.GitHubContext{}
	//if event.EventName == "issue_comment" {

	//githubactions.Fatalf("This action is only supported for issue_comment event")
}

/*
	client := github.NewClient(nil)
	// list all organizations for user "willnorris"
	orgs, _, _ := client.Organizations.List(context.Background(), "hydai", nil)
	fmt.Println("orgs:", orgs)
*/
