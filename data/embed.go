package data

import "embed"

var (
	//go:embed botPolicies.yaml botPolicies.json all:apps all:bots all:common all:crawlers
	BotPolicies embed.FS
)
