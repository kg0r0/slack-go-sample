package main

import (
	"flag"
	"fmt"

	"github.com/kg0r0/slack-go-sample/slackbot"
	"github.com/nlopes/slack"
)

func main() {
	config, err := slackbot.NewConfig("conf/slack_conf.json")
	if err != nil {
		panic("conf read error")
	}
	if config.SlackConfig.AccessToken == "" {
		panic("no access token")
	}
	t := flag.String("t", config.SlackConfig.AccessToken, "access token")
	flag.Parse()
	if len(*t) == 0 {
		panic("no token")
	}

	api := slack.New(*t)
	user, err := api.GetUserInfo("U023BECGF")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
}
