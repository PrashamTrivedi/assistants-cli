# My Go CLI Project

This is a CLI project written in Go. It includes the following commands:

- `create`: creates a new resource
- `remove`: removes an existing resource
- `start`: starts a resource
- `update`: updates an existing resource

## Configuration

The project includes configuration files in both YAML and TOML formats. These files can be found in the `config` directory.

## Dependencies

This project uses the following dependencies:

- [Go Releaser](https://goreleaser.com/)
- [Go OpenAI](https://github.com/sashabaranov/go-openai)

## Development

To develop this project, use the included devcontainer for Go. This container includes all necessary tools and dependencies.

## Project Structure

```
my-go-cli-project
├── cmd
│   ├── config.go
│   ├── create.go
│   ├── remove.go
│   ├── start.go
│   └── update.go
├── config
│   ├── config.yaml
│   └── config.toml
├── internal
│   ├── assistant
│   │   ├── assistant.go
│   │   ├── prompt.go
│   │   └── remove.go
│   └── chat
│       ├── chat.go
│       └── start.go
├── go.mod
├── go.sum
├── main.go
├── README.md
├── .devcontainer
│   ├── devcontainer.json
│   └── Dockerfile
└── .vscode
    ├── settings.json
    └── launch.json
```