# VK-Finance

A simple Go console app for tracking personal finance data. Runs as an interactive command-line shell with data persisted locally.

## Tech Stack

- [Go](https://go.dev/)
- Local file/JSON-based storage (no external database required)

## Project Structure

```
├── DATABASES/FINANCE/   # Local data storage for finance records
├── pkg/
│   ├── cmd/               # Command handling / interactive shell logic
│   ├── config/             # Configuration (e.g. local storage file path)
│   ├── db/                 # Finance data model and load/save logic
│   └── util/                # Utility helpers (local storage init, etc.)
├── main.go                # Entry point
├── vk-finance.json        # Sample/default data file
├── go.mod
└── go.sum
```

## How It Works

On startup, `VK-Finance`:

1. Initializes local storage (`util.InitLocalStorage`).
2. Loads existing finance data from a local file (`db.Finance.LoadFromFile`).
3. Starts an interactive command shell (`cmd.Start`) where you can enter commands to manage your finance data.

## Getting Started

### Prerequisites

- Go installed

### Installation

```bash
git clone https://github.com/VkHyperNova/VK-FINANCE.git
cd VK-FINANCE
go build
```

> **Note:** Build with `go build` (not `go build main.go`) so the generated Windows resources (icon/manifest, see below) are included correctly.

### Usage

Run the built binary:

```bash
./VK-FINANCE
```

Or run directly with Go:

```bash
go run main.go
```

This launches an interactive shell where you can log and manage your finance data, which is persisted locally between sessions.

## Building a Windows .exe with a Custom Icon

This project can be built into a Windows executable with a custom icon and manifest using [`goversioninfo`](https://github.com/josephspurrier/goversioninfo).

1. Install `goversioninfo`:

   ```bash
   go install github.com/josephspurrier/goversioninfo/cmd/goversioninfo
   ```

2. Copy `testdata/resource/versioninfo.json` into your working directory and modify it with your own settings.

3. If needed, regenerate `go.mod`:

   ```bash
   go mod init github.com/VkHyperNova/VK-FINANCE
   go mod tidy
   ```

4. Copy your manifest file into the `resource` folder.

5. Add the following `go:generate` directive at the top of the relevant Go file:

   ```go
   //go:generate goversioninfo -icon=testdata/resource/icon.ico -manifest=testdata/resource/goversioninfo.exe.manifest
   ```

6. Generate the `resource.syso` file:

   ```bash
   go generate
   ```

7. Build the project:

   ```bash
   go build
   ```

   > Just run `go build` — building a single file directly (e.g. `go build main.go`) will not pick up the generated resources.

## License

No license specified — all rights reserved by the author unless stated otherwise.
