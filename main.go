package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	//slack-bot-token changed
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-1234-3949201490449-5389446008610-185A5SeUPF8M01XIbvPD8BWL")
	//channel-ID changed
	os.Setenv("CHANNEL_ID", "C123TJGP1SSW")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"ZIPL.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}