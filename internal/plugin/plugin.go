package plugin

import (
	"context"
	"strings"

	rayleabot "github.com/RayleaBot/RayleaBot/sdk/go"
)

func Run(ctx context.Context) error {
	return rayleabot.Run(ctx, rayleabot.Options{}, rayleabot.HandlerFunc(handleEvent))
}

func handleEvent(_ context.Context, event *rayleabot.EventContext) error {
	if event.Event.Command() != "echo" {
		return event.Result(map[string]any{"handled": false})
	}
	return event.SendText(echoText(event.Event.Args()))
}

func echoText(args []string) string {
	text := strings.Join(args, " ")
	if strings.TrimSpace(text) == "" {
		return "(空消息)"
	}
	return text
}
