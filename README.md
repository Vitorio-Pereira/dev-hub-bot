# dev-hub-bot

🇧🇷 [Leia em Português](./README.pt-br.md)

A Discord bot that turns a server into a personal hub for organization: project tracking, job/freelance opportunities, backlog of ideas/studies, and a base of technical resources — with each project's tech stack automatically detected via the GitHub API.

Written in Go, using [`discordgo`](https://github.com/bwmarrin/discordgo), SQLite for storage, and a fully customizable server structure via `config.yaml`.

## Motivation

Keeping track of multiple personal projects (status, next steps, stack, priority) scattered across notes, GitHub issues, and memory doesn't scale. This bot centralizes that inside a Discord server, with a pinned embed per project acting as its "current state" — always synced with real data (GitHub) and updated through commands, never edited by hand.

## Features

- **Automatic server setup**: a single command applies a `config.yaml` and creates/updates categories, channels, and roles — idempotently (running it again won't duplicate anything).
- **Project tracking**: each project has a channel with a pinned embed showing current phase, progress, priority, difficulty, stack, done/next steps/blockers.
- **Auto-detected stack**: when creating a project from a repository link, the bot queries the GitHub API and classifies the stack into Languages / Infra / Other.
- **Multi-server**: all data is isolated by `guild_id` — the same bot can be invited to any server, each with its own configuration and projects.

## Tech stack

| Component | Choice |
|---|---|
| Language | Go |
| Discord | [`discordgo`](https://github.com/bwmarrin/discordgo) |
| Storage | SQLite |
| External integration | GitHub REST API |
| Config | YAML (`config.yaml`, customizable per server) |
| Style guide | [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) |

## Project structure

```
dev-hub-bot/
├── cmd/
│   └── bot/
│       └── main.go              # bootstrap: config, discordgo session, graceful shutdown
├── internal/
│   ├── discord/
│   │   ├── session.go           # thin wrapper over *discordgo.Session
│   │   ├── router.go            # dispatches InteractionCreate to the right handler
│   │   └── commands/             # implementation of each slash command
│   ├── github/
│   │   └── client.go            # fetches repo + languages, stack classification
│   ├── project/
│   │   ├── project.go           # domain structs (Project, Stack, Priority, Difficulty)
│   │   ├── store.go             # Store interface
│   │   └── sqlite_store.go      # SQLite implementation
│   ├── config/
│   │   └── config.go            # config.yaml parser and validation
│   └── embed/
│       └── render.go            # Project -> *discordgo.MessageEmbed
├── config.example.yaml
├── .env.example
├── go.mod
└── go.sum
```

## Running locally

### Prerequisites

- Go 1.22+
- A Discord application created in the [Discord Developer Portal](https://discord.com/developers/applications), with a bot enabled
- A [GitHub Personal Access Token](https://github.com/settings/tokens) (fine-grained, `Contents: Read-only` scope is enough)

### Steps

```bash
git clone https://github.com/<your-username>/dev-hub-bot.git
cd dev-hub-bot

cp .env.example .env
# fill in DISCORD_TOKEN, DISCORD_APP_ID, GITHUB_TOKEN, and the other variables

go mod tidy
go run ./cmd/bot
```

Invite the bot to a test server with the `bot` and `applications.commands` scopes, then run `/setup` attaching a `config.yaml` (see `config.example.yaml` as a starting point) to generate the category/channel/role structure.

## Server configuration (`config.yaml`)

The server's category, channel, and role structure isn't hardcoded — it's defined by a YAML file applied via `/setup`. This lets each server customize names, emojis, and organization, while the bot keeps working through stable internal keys (`key`).

See `config.example.yaml` for the full format and a ready-to-use example structure.

## Project status

In early development. Follow progress and upcoming milestones on the repository's issues/projects board.

## License

MIT — see [`LICENSE`](./LICENSE).