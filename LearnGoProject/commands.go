package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
)

//help command
func OnHelp(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//build string to display help/info/commands
	embed := buildEmbed(config.CommandTitle, config.Commands)
	s.ChannelMessageSendEmbed(msg.ChannelID, &embed)
}

//jokeThere command
func OnJokeThere(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//confirm channel ID exists
	if len(arg[:]) > 1 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(jokes))
		s.ChannelMessageSend(arg[1], jokes[randomIndex])
	} else {
		s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
	}
}

//jokeHere command
func OnJokeHere(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//creates a random seed
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(jokes))
	s.ChannelMessageSend(msg.ChannelID, jokes[randomIndex])
}

//factsThere command
func OnFactsThere(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//confirm channel ID exists
	if len(arg[:]) > 1 {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(facts))
		s.ChannelMessageSend(arg[1], facts[randomIndex])
	} else {
		s.ChannelMessageSend(msg.ChannelID, "Invalid Channel. Use /help to see commands")
	}
}

//factsHere command
func OnFactsHere(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//creates a random seed
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(facts))
	s.ChannelMessageSend(msg.ChannelID, facts[randomIndex])
}

//+say command
func OnSet(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//set the channel topic
	//turns out there is a long cool down on this so I cant use this
	//s.ChannelEditComplex(msg.ChannelID, &discordgo.ChannelEdit{
	//	Topic: ":red_circle:  Currently sending all messages to " + arg[1],
	//})
	//checks to see if user is already active

	m[msg.Author.Username] = arg[1]

	//confirm channel id is in list and print id
	if len(arg[:]) > 1 {
		s.ChannelMessageSend(arg[1], arg[2])
		s.ChannelMessageSend(msg.ChannelID, "You are now sending messages to <#"+arg[1]+">")
	} else {
		s.ChannelMessageSend(msg.ChannelID, "You are currently sending messages to <#"+arg[1]+">")
	}
	fmt.Println(m)
}

func OnText(s *discordgo.Session, msg *discordgo.MessageCreate) {
	//posts all messages to other channel
	s.ChannelMessageSend(m[msg.Author.Username], msg.Content)

}

//-say command
func OnUnset(s *discordgo.Session, msg *discordgo.MessageCreate, arg []string) {
	//clear the channel topic
	//turns out there is a long cool down on this so I cant use this
	//s.ChannelEditComplex(msg.ChannelID, &discordgo.ChannelEdit{
	//	Topic: "",
	//})
	if _, exists := m[msg.Author.Username]; exists {
		//if user exists delete record in map
		s.ChannelMessageSend(msg.ChannelID, "You are no longer sending messages to channel <#"+m[msg.Author.Username]+">")
		delete(m, msg.Author.Username)
	} else {
		//if user doesnt exist return
		s.ChannelMessageSend(msg.ChannelID, "Say is not currently active for "+msg.Author.Username)
	}
	fmt.Println(m)
}