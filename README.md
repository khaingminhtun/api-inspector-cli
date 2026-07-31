"# API Inspector CLI

A simple Go CLI using Cobra for inspecting HTTP APIs from the terminal.

## What it does

- Send `GET` and `POST` requests
- Add custom headers
- Print responses in `table`, `json`, or `yaml` format
- Save, load, list, and delete named requests for quick reuse

## What I learned

- How to build a CLI with `Cobra` and organize subcommands cleanly
- How `Viper` can load and manage configuration settings for the app
- How to define persistent flags, command flags, and shared header handling
- How shell completion works in a Go CLI and how to expose it through Cobra
- How to structure a small Go project with commands, config, storage, and formatting layers

## Run it

From the project root:

```bash
go run .
```

Or build it:

```bash
go build
```

## Main commands

```bash
apispy get https://jsonplaceholder.typicode.com/todos/1
apispy post https://example.com/api -d '{"name":"demo"}'
apispy save my-request https://jsonplaceholder.typicode.com/todos/1
apispy load my-request
apispy list
apispy delete my-request
```

## Useful flags

```bash
apispy get https://example.com/api --output json
apispy get https://example.com/api --timeout 10 -H "Accept: application/json"
```

## Output formats

Use `--output` with one of:

- `table`
- `json`
- `yaml`

## License

This project is licensed under the MIT License." 
