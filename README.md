# gib

**gib** is a terminal command suggestion tool powered by AI.
Describe what you want to do in plain English, and `gib` will respond with the most likely terminal command.
Think of it as your helpful shell companion for when memory fails or syntax escapes you.

## Features

- Natural language to terminal command translation
- AI-powered suggestions
- Shell-agnostic: works in any terminal
- Lightweight and fast (written in Go)
- Outputs commands without running them — you're always in control

## Roadmap

- [x] Copy to clipboard on all platforms
- [ ] Model choosing
- [ ] Support other AI providers
- [ ] Explain commands

## Installation

Via ```go install```:

```bash
go install github.com/yourusername/gib@latest
```

From source:

```bash
git clone https://github.com/yourusername/gib.git
cd gib
go build -o gib
./gib
```

## Configuration

An OpenAI API key must be set like so

```bash
export OPENAI_API_KEY=<your key here>
```

## Usage

```bash
gib <message>
```
Examples:

```bash
gib list all files including hidden ones
> ls -a

gib run a http server on 8080
> python3 -m http.server 3000

gib create a tarball of all .log files
> tar -cvf logs.tar *.log
```

## License

gib is licensed under the [MIT License](LICENSE.md)
