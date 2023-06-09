package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"os"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5180972964100-5176258582550-ZhomMy1GGDvQ9FefUZxWkdCw")
	os.Setenv("CHANNEL_ID", "C055M3HPZPB")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"rahul.pdf"}

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
