package plugin

import (
	"context"
	"strings"

	rayleabot "github.com/RayleaBot/RayleaBot/sdk/go"
)

func Run(ctx context.Context) error {
	return rayleabot.Run(ctx, rayleabot.Options{
		PluginID:      "raylea.echo",
		Subscriptions: []string{"message.group", "message.private"},
	}, rayleabot.HandlerFunc(handleEvent))
}

func handleEvent(_ context.Context, event *rayleabot.EventContext) error {
	if event.Event.Command() != "echo" {
		return event.Result(map[string]any{"handled": false})
	}
	text := strings.TrimSpace(strings.Join(event.Event.Args(), " "))
	if text == "" {
		text = "（空消息）"
	}
	return event.SendText(text)
}
