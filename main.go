package main

import (
  "os"
  "log"
  "github.com/nlopes/slack"
)

func main() {

  SLACK_API_TOKEN, exists := os.LookupEnv("SLACK_API_TOKEN")

  if !exists {
    log.Fatal("SLACK_API_TOKEN not set")
  }

  api := slack.New(SLACK_API_TOKEN)

  rtm := api.NewRTM()

  go rtm.ManageConnection()

  for msg := range rtm.IncomingEvents {
    switch ev := msg.Data.(type) {
    case *slack.ConnectedEvent:
      log.Println("Connected!")
    case *slack.UserTypingEvent:
      rtm.SendMessage(rtm.NewTypingMessage(ev.Channel))
    }
  }
}
