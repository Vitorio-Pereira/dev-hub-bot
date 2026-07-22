# dev-hub-bot

🇺🇸 [Read in English](./README.md)

Bot de Discord para transformar um servidor em um hub pessoal de organização: tracking de projetos, vagas, propostas de freela, backlog de ideias/estudos e uma base de recursos técnicos — com stack de cada projeto detectada automaticamente via GitHub API.

Escrito em Go, usando [`discordgo`](https://github.com/bwmarrin/discordgo), storage em SQLite, e estrutura de servidor totalmente personalizável via `config.yaml`.

## Motivação

Manter tracking de múltiplos projetos pessoais (status, próximos passos, stack, prioridade) espalhado entre anotações, issues do GitHub e memória não escala. Este bot centraliza isso num servidor Discord, com um embed fixado por projeto que serve como "estado atual" — sempre sincronizado com dados reais (GitHub) e atualizado via comandos, nunca editado manualmente.

## Funcionalidades

- **Setup automático de servidor**: um comando aplica um `config.yaml` e cria/atualiza categorias, canais e roles — de forma idempotente (rodar de novo não duplica nada).
- **Tracking de projetos**: cada projeto tem um canal com embed fixado mostrando fase atual, progresso, prioridade, dificuldade, stack, feito/próximos passos/bloqueios.
- **Stack detectada automaticamente**: ao criar um projeto a partir de um link de repositório, o bot consulta a GitHub API e classifica a stack em Linguagens / Infra / Outros.
- **Multi-servidor**: todo dado é isolado por `guild_id` — o mesmo bot pode ser convidado para qualquer servidor, cada um com sua própria configuração e projetos.

## Stack técnica

| Componente | Escolha |
|---|---|
| Linguagem | Go |
| Discord | [`discordgo`](https://github.com/bwmarrin/discordgo) |
| Storage | SQLite |
| Integração externa | GitHub REST API |
| Config | YAML (`config.yaml`, personalizável por servidor) |
| Guia de estilo | [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md) |

## Estrutura do projeto

```
dev-hub-bot/
├── cmd/
│   └── bot/
│       └── main.go              # bootstrap: config, sessão discordgo, graceful shutdown
├── internal/
│   ├── discord/
│   │   ├── session.go           # wrapper sobre *discordgo.Session
│   │   ├── router.go            # despacha InteractionCreate pros handlers
│   │   └── commands/             # implementação de cada slash command
│   ├── github/
│   │   └── client.go            # busca repo + languages, classificação de stack
│   ├── project/
│   │   ├── project.go           # structs de domínio (Project, Stack, Priority, Difficulty)
│   │   ├── store.go             # interface Store
│   │   └── sqlite_store.go      # implementação SQLite
│   ├── config/
│   │   └── config.go            # parser e validação do config.yaml
│   └── embed/
│       └── render.go            # Project -> *discordgo.MessageEmbed
├── config.example.yaml
├── .env.example
├── go.mod
└── go.sum
```

## Como rodar localmente

### Pré-requisitos

- Go 1.22+
- Uma aplicação Discord criada no [Discord Developer Portal](https://discord.com/developers/applications), com bot habilitado
- Um [Personal Access Token do GitHub](https://github.com/settings/tokens) (fine-grained, escopo `Contents: Read-only` é suficiente)

### Passos

```bash
git clone https://github.com/<seu-usuario>/dev-hub-bot.git
cd dev-hub-bot

cp .env.example .env
# preencha DISCORD_TOKEN, DISCORD_APP_ID, GITHUB_TOKEN e demais variáveis

go mod tidy
go run ./cmd/bot
```

Convide o bot para um servidor de teste com os escopos `bot` e `applications.commands`, e rode `/setup` anexando um `config.yaml` (veja `config.example.yaml` como ponto de partida) para gerar a estrutura de categorias, canais e roles.

## Configuração do servidor (`config.yaml`)

A estrutura de categorias, canais e roles do servidor não é fixa no código — é definida por um arquivo YAML aplicado via `/setup`. Isso permite que cada servidor customize nomes, emojis e organização, mantendo o bot funcional através de chaves (`key`) estáveis internas.

Veja `config.example.yaml` para o formato completo e um exemplo de estrutura pronta para uso.

## Status do projeto

Em desenvolvimento inicial. Acompanhe o progresso e os próximos marcos no board de issues/projects do repositório.

## Licença

MIT — veja [`LICENSE`](./LICENSE).