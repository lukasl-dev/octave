package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// ErrorResponse returns discordgo.InteractionResponseData with the error message
// as content.
func ErrorResponse(err error) *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Content: fmt.Sprintf(":name_badge: Failed to run command: **%s**", err),
	}
}
