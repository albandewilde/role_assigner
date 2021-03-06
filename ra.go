package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	dgo "github.com/bwmarrin/discordgo"
)

var discordToken string
var channelsId []string
var messagesId []string

func init() {
	// Read the discord token
	discordToken = os.Getenv("discord")

	// Read the channel and message id
	channelsId = strings.Split(os.Getenv("channels"), ", ")
	messagesId = strings.Split(os.Getenv("messages"), ", ")

	if len(channelsId) != len(messagesId) {
		log.Fatal("Not the same number of channelsId and messagesId")
	}
}

func main() {
	// Create discord bot instance
	bot, err := dgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal(err)
	}

	// Register function callback
	bot.AddHandler(manageRolesAdd)
	bot.AddHandler(manageRolesRemove)

	// Open bot connection
	err = bot.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("I'm logged in ! (Press CTRL-C to exit.)\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func manageRolesAdd(s *dgo.Session, m *dgo.MessageReactionAdd) {
	if rolesAssignerMessage(m.MessageID, m.ChannelID) {
		guildRoles, _ := s.GuildRoles(m.GuildID)
		askedRoleId, _ := roleIdFromEmote(m.Emoji, guildRoles)

		// Add the role
		s.GuildMemberRoleAdd(m.GuildID, m.UserID, askedRoleId)
	}

}

func manageRolesRemove(s *dgo.Session, m *dgo.MessageReactionRemove) {
	if rolesAssignerMessage(m.MessageID, m.ChannelID) {
		guildRoles, _ := s.GuildRoles(m.GuildID)
		askedRoleId, _ := roleIdFromEmote(m.Emoji, guildRoles)

		// Remove the role
		s.GuildMemberRoleRemove(m.GuildID, m.UserID, askedRoleId)
	}
}

func roleIdFromEmote(e dgo.Emoji, guildRoles []*dgo.Role) (string, error) {
	// Get the role id
	roles, err := readRoles()
	if err != nil {
		return "", err
	}
	askedRoleName := roles[e.Name]

	var askedRoleId string
	for _, r := range guildRoles {
		if r.Name == askedRoleName {
			askedRoleId = r.ID
			break
		}
	}

	// Chech if the role is found in the guild
	if askedRoleId == "" {
		return "", fmt.Errorf("Role `%s` not found in the guild", askedRoleName)
	}

	return askedRoleId, nil
}

func readRoles() (roles map[string]string, err error) {
	rolesFile, err := ioutil.ReadFile("./roles.json")
	if err != nil {
		return
	}

	err = json.Unmarshal(rolesFile, &roles)
	if err != nil {
		return
	}
	return
}

func rolesAssignerMessage(messageId, channelId string) bool {
	for idx := 0; idx < len(messagesId); idx += 1 {
		// Compare message and channel id have the same index
		if messageId == messagesId[idx] && channelId == channelsId[idx] {
			return true
		}
	}
	return false
}
