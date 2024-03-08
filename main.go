package main

import (
	"context"
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"github.com/google/go-github/v60/github"
)

func main(){	
	client := github.NewClient(nil)
	// list all organizations for user "willnorris"
	orgs, _, _ := client.Organizations.List(context.Background(), "hydai", nil)
	fmt.Println("orgs:", orgs)

	val := githubactions.GetInput("val")
  if val == "" {
    githubactions.Fatalf("missing 'val'")
  }

}
