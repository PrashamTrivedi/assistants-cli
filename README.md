# Assistant CLI
This is a CLI project to create and chat with an OpenAI assistants. 

# About

The Assistant CLI is a project that allows you to create and chat with OpenAI assistants through a Command Line Interface (CLI). It provides a convenient way to interact with chat assistants and utilize their capabilities.


# Usage


To use the Assistant CLI, you can run the following commands:

```bash
assistants [flags]
assistants [command]
```

Available Commands:

- `add`: Create a new assistant
- `chat`: Chat with one of the assistants
- `completion`: Generate the autocompletion script for the specified shell
- `config`: Set OpenAI Key and file paths to store assistants and chats
- `help`: Help about any command
- `list`: List Assistants
- `listModels`: List Available models
- `logs`: Shows your chat history
- `remove`: Remove a resource
- `update`: Update the assistant

Flags:

- `-h, --help`: Show help for the Assistant CLI
- `-t, --toggle`: Help message for a toggle
- `-v, --version`: Show the version of the Assistant CLI

These commands and flags provide various functionalities for creating, managing, and interacting with OpenAI assistants. You can use them to create new assistants, chat with existing assistants, configure settings, list available models, view chat history, and more.

## Configuration

Before using the Assistant CLI, you need to configure it with your OpenAI key and file paths for storing assistants and chats. To do this, use the `config` command.

```bash
assistants config
```

The `config` command has the following flags:

- `-k, --key`: Specifies your OpenAI API key. This is required for authentication with OpenAI services.

- `-a, --assistant-path`: Specifies the file path for storing assistant data. This can be a local file path or a cloud storage provider like AWS S3 or Google Cloud Storage.

- `-c, --chat-log-path`: Specifies the file path for storing chat log data. Similar to the assistant path, this can also be a local file path or a cloud storage provider.

Example usage:

```bash
assistants config -k <YOUR_API_KEY> -a /path/to/assistants -c /path/to/chat_logs
```

Make sure to replace `<YOUR_API_KEY>` with your actual OpenAI API key and `/path/to/assistants` and `/path/to/chat_logs` with the desired file paths for your assistant and chat log data.

Once you have provided the necessary configuration information, the Assistant CLI will be ready to use with your specified settings.
## Assistants

The Assistant CLI allows you to create, update, delete, and list assistants. This section will explain the available commands and flags for managing assistants.

### Create Assistant

To create a new assistant, use the `add` command with the following flags:


- `-m, --model string`: Specifies the default model to use with the assistant.
- `-n, --name string`: Specifies the name of the assistant.
- `-p, --prompt string`: Specifies the prompt for the assistant.

#### Note: Following flags are subject to change in future releases.
- `-c, --allow-commands`: Allows the assistant to run commands.
- `-f, --allow-file-reading`: Allows the assistant to read files.
- `-s, --allow-search`: Allows the assistant to search the web.

Example usage:

```bash
assistants add --allow-commands --allow-file-reading --allow-search --model gpt-3.5-turbo --name MyAssistant --prompt "How can I assist you?"
```

This command will create a new assistant with the specified flags and values.

### Update Assistant

To update an existing assistant, use the `update` command. This command allows you to modify the assistant's configuration, such as changing the prompt, model, or other settings.

Example usage:

```bash
assistants update --name MyAssistant --prompt "What else can I help you with?"
```

This command will update the assistant named "MyAssistant" with the new prompt.

### Delete Assistant

To delete an assistant, use the `remove` command followed by the name of the assistant.

Example usage:

```bash
assistants remove --name MyAssistant
```

This command will delete the assistant named "MyAssistant".

### List Assistants

To list all the assistants that have been created, use the `list` command.

Example usage:

```bash
assistants list
```

This command will display a list of all the assistants along with their names and other details.

Note: The `add`, `update`, `remove`, and `list` commands require authentication and proper configuration of the Assistant CLI. Make sure to configure the CLI with your OpenAI key and file paths before using these commands.

## Chat

The Assistant CLI allows you to chat with the assistants you have created. This section will explain the command and flags for initiating or continuing a chat session.

To start or continue a chat with an assistant, use the `chat` command with the following flags:

- `-a, --assistantId string`: Specifies the name or ID of the assistant you want to chat with.

- `-c, --chatId string`: Specifies the ID of the chat to continue. You can provide the chat ID to resume a previous conversation.

- `--continue`: A flag that indicates to continue with the latest chat ID. This flag will automatically continue the chat with the most recent conversation.

- `-m, --message string`: Specifies the message you want to send to the assistant.

Example usage:

```bash
assistants chat --assistantId MyAssistant --message "Hello, how can I assist you?"
```

This command will start a chat with the assistant named "MyAssistant" and send the initial message.

If you want to continue a chat with a specific chat ID, you can use the following command:

```bash
assistants chat --assistantId MyAssistant --chatId ABC123
```

This command will continue the chat with the specified chat ID.

Additionally, you can use the `logs` command to view the chat history. This command will show you the logs of previous chat sessions, including the messages exchanged with the assistant.

Example usage:

```bash
assistants logs
```

Make sure to use the `logs` command to keep track of your chat conversations and reference them when needed.