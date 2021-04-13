package discordGoLog;

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"github.com/bwmarrin/discordgo"
)

const maxMessageLen = 2000;

func sendLogToDiscordWebhookInternal(url string, data []byte) error {
	_, err := http.Post(url, "application/json", bytes.NewBuffer(data));

	if err != nil {
		return err;
	}

	return nil;
}

type body struct {
	Embeds []discordgo.MessageEmbed `json:"embeds"`
	Username string `json:"username"`
	Content string `json:"content"`
}

func SendLogToDiscordEmbed(url string, embedType discordgo.MessageEmbed, username string) error  {
	embeds := []discordgo.MessageEmbed{embedType};

	bdy := body{
		Embeds:   embeds,
		Username: username,
	}

	embedJson, err := json.Marshal(bdy);

	if err != nil {
		return err;
	}

	return sendLogToDiscordWebhookInternal(url, embedJson);
}

func SendLogToDiscordString(url, message, username string) error {
	bdy := body{
		Username: username,
		Content: message,
	}

	if len(message) > maxMessageLen {
		return errors.New(fmt.Sprintf("message length must be less than %d characters", maxMessageLen))
	}

	embedJson, err := json.Marshal(bdy);

	if err != nil {
		return err;
	}

	return sendLogToDiscordWebhookInternal(url, embedJson);
}