package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var buffer = make([][]byte, 0)

var webhook = os.Getenv("DISCORD_WEBHOOK")
var token = os.Getenv("DISCORD_BOT_TOKEN")

type JokeApiResponse struct {
	Value JokeBody `json:"value`
	Type  string   `json:"type`
}

type JokeBody struct {
	Id         int      `json:"id"`
	Joke       string   `json:"joke"`
	Categories []string `json:"categories"`
}

func main() {

	if token == "" {
		fmt.Println("No token provided. Please run: discord-chuck-norris -t <bot token>")
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// Register ready as a callback for the ready events.
	dg.AddHandler(ready)

	// Register messageCreate as a callback for the messageCreate events.
	dg.AddHandler(messageCreate)

	// Open the websocket and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Discord Chuck Norris bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func ready(s *discordgo.Session, event *discordgo.Ready) {

	// Set the playing status.
	s.UpdateStatus(0, "!joke")
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("Recieved Message: ")
	fmt.Println(m.Content)
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// check if the message is "!joke"
	if strings.HasPrefix(m.Content, "!joke") {
		// Send Joke
		err := sendJoke()
		if err != nil {
			fmt.Println(err)
		}
	}
}

//Send joke as a tts (test to speech) message
func sendJoke() (err error) {
	// Get a Chuck Norris joke
	joke, err := getJoke()
	if err != nil {
		return err
	}
	// Send
	resp, err := http.PostForm(webhook, url.Values{"content": {joke}, "tts": {"true"}})
	fmt.Println(resp)
	if err != nil {
		fmt.Println("Couldn't send message")
		fmt.Println(err)
		return err
	} else {
		fmt.Println(resp)
		return err
	}

	return nil
}

//Fetch Chuck Norris Joke
func getJoke() (string, error) {
	resp, err := http.Get("http://api.icndb.com/jokes/random")
	if err != nil {
		fmt.Println("Could not fetch joke")
		return "nil", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Unknown response body")
		return "nil", err
	}

	var jokeResp JokeApiResponse
	json.Unmarshal(body, &jokeResp)
	fmt.Println(jokeResp)
	return jokeResp.Value.Joke, nil
}
