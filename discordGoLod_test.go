package discordGoLog;

import (
	"testing"
	"github.com/bwmarrin/discordgo"
)

const url = ""

func TestSendLogToDiscordEmbed(t *testing.T) {
	if url == "" {
		t.Fatalf("Skipped test. URL set to \"\"");
		return;
	}

	err := SendLogToDiscordEmbed(url, discordgo.MessageEmbed{
		Title:       "TEST",
		Description: "just a quick test",
		Color:       123325,
	}, "Go test");

	if err != nil {
		t.Fatalf("Test failed with error: %s", err.Error())
	}
}

func TestSendLogToDiscordString(t *testing.T) {
	if url == "" {
		t.Fatalf("Skipped test. URL set to \"\"");
		return;
	}

	err := SendLogToDiscordString(url, "Test message", "Go test");

	if err != nil {
		t.Fatalf("Test failed with error: %s", err.Error())
	}
}

func TestSendLogToDiscordString2(t *testing.T) {
	if url == "" {
		t.Fatalf("Skipped test. URL set to \"\"");
		return;
	}

	message := "";

	for i := 0; i < 501; i++ {
		message += "TEST";
	}

	err := SendLogToDiscordString(url, message, "Go test");

	if err == nil {
		t.Fatalf("Test failed! Did not return error with message len higher than 2000:")
	}
}