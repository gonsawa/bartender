package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating session: ", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection: ", err)
		return
	}

	fmt.Println("Bartender is now running. Press CTRL-C to shutdown Bartender.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

//Message listener
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	//Check if the message starts with the prefix
	if strings.HasPrefix(m.Content, "_") {
		if m.Content == "_pong" {
			//Simple pong command message
			s.ChannelMessageSend(m.ChannelID, "Pong!")
		} else {
			//Command not found
			s.ChannelMessageSend(m.ChannelID, "Comando não encontrado!")
		}
	}

}
