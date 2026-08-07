package plugin

import (
	"testing"

	rayleabot "github.com/RayleaBot/RayleaBot/sdk/go"
)

func TestEchoManifestCommandInput(t *testing.T) {
	event := rayleabot.Event{Payload: map[string]any{"command": "echo", "args": []any{"hello", "go"}}}
	if event.Command() != "echo" || len(event.Args()) != 2 {
		t.Fatalf("unexpected command input: %#v", event)
	}
}

func TestEchoTextPreservesCommandArgumentsAndEmptyFallback(t *testing.T) {
	if got := echoText([]string{" hello ", "go"}); got != " hello  go" {
		t.Fatalf("echoText preserved output = %q", got)
	}
	if got := echoText([]string{"", " "}); got != "(空消息)" {
		t.Fatalf("echoText empty output = %q", got)
	}
}
