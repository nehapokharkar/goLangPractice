package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Evenets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5142123571527-5156726605267-wUzmM1HnT9kyV45uFPCdYdVh")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A054E28RGVC-5150081248134-a2644b40b45f81cda887a4221d07cd7f49d247e7e8fe55d8f72ab54a4d27cc7a")

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	bot.Command("my name is <name>", &slacker.CommandDefinition{
		Description: "Greet",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			name := request.Param("name")
			quotestr := strconv.Quote(name)
			str, _ := strconv.Unquote(quotestr)
			fmt.Println(name)
			r := fmt.Sprintf("Welcome to Slack-Age-Bot. Hi %s", str)
			response.Reply(r)
			r = fmt.Sprint("Please enter your yob")
			response.Reply(r)
		},
	})

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			t := time.Now()
			thisYear := t.Year()
			if thisYear > yob {
				age := thisYear - yob
				r := fmt.Sprintf("Your age is %d", age)
				response.Reply(r)
			} else {
				r := fmt.Sprintf("Please enter valid year")
				response.Reply(r)
			}
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
