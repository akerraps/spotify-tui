# tunectl – Terminal Spotify Playlist Manager (Learning Project)

TuneCtl is a personal learning project written in Go.
The goal is to explore different ways of interacting with the Spotify Web API from the terminal, using both a classic CLI and (eventually) a full TUI.

This is not a production-ready tool. It’s an experimental sandbox where I learn Go, project structure, APIs, and terminal UX. The code may be messy, incomplete, or broken in places — that’s intentional and part of the process.

## Project Goal

The main purpose of this project is learning and experimentation. In particular, I want to:

- [x] Authenticate with the Spotify Web API
- [x] Fetch user playlists
- [x] Fetch tracks from a specific playlist
- [x] Be able to download songs locally
- [ ] Add a terminal UI (TUI) using Bubble Tea
- [ ] Improve error handling and UX
- [ ] Experiment with clean project architecture in Go

>Note: All interactions are done through official Spotify Web API endpoints and within Spotify’s terms of service. This project is for personal use and learning only.
>
## Music Downloading (yt-dlp Integration)

Song downloading is implemented using yt-dlp as a backend:

- Repository: <https://github.com/yt-dlp/yt-dlp>

TuneCtl acts as a thin wrapper around `yt-dlp` to fetch audio from YouTube based on track metadata.

### Why yt-dlp?

Using yt-dlp provides several advantages:

- It is actively maintained and adapts quickly to YouTube changes
- It significantly increases the long-term durability of the project
- It avoids reinventing a complex and constantly changing downloader
- It keeps TuneCtl focused on orchestration and learning, not scraping

>This project is for personal use and learning only. All Spotify interactions are done through the official Spotify Web API and within Spotify’s terms of service.

## Project Structure (High-Level)

```
.
├── cmd
│   └── tunctl
│       └── main.go          # Application entrypoint
├── internal
│   ├── cli                  # CLI layer (urfave/cli)
│   │   └── cli.go
│   ├── core                 # Central logic (Spotify auth, playlists, tracks)
│   │   ├── authenticate.go
│   │   ├── core.go
│   │   └── playlists.go
│   ├── fetcher              # yt-dlp wrapper for audio downloads
│   │   └── fetcher.go
│   ├── tui                  # TUI (planned, currently empty)
│   │   └── tui.go
│   └── types                # Shared structs and types
│       └── types.go
├── go.mod
├── go.sum
└── README.md
```

### Architecture Philosophy

- core: does the real work (Spotify API, logic)
- cli: wires commands, flags, and arguments to the core
- fetcher: wraps yt-dlp for downloading audio
- types: shared data structures used across the project
- tui: reserved for future Bubble Tea implementation

This separation is intentional so the same core logic can be reused by both CLI and TUI.

## Getting Started

To use tunectl, you need to have a Spotify Developer account and create an application to obtain your credentials.

1. Go to Spotify [Developer Dashboard](https://developer.spotify.com/) and create a new application.
2. Copy your Client ID and Client Secret.
3. Create a .env file in the project root with the following variables:

```
SPOTIFY_ID=your_client_id_here
SPOTIFY_SECRET=your_client_secret_here
```

4. Run tunectl from your terminal. The application will use these credentials to connect to your Spotify account.

>Note: This project only accesses your own Spotify data for personal organization and learning purposes.

## Usage

### General help

```bash
go run cmd/tunctl/main.go -h
```

Output:

```
NAME:
   tunectl - Manage your playlists and songs

USAGE:
   tunectl [global options] command [command options]

COMMANDS:
   playlists  Manage playlists
   songs      Manage songs
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### `playlists` command

Manage Spotify playlists.

**List playlists**

```bash
go run cmd/tunctl/main.go playlists --list
```

Lists all playlists associated with your Spotify account.

**Download an entire playlist**

```bash
go run cmd/tunctl/main.go playlists --download "<playlist name>" --output ./music
```

- Downloads all songs from the specified playlist
- Uses `yt-dlp` under the hood
- Requires an output directory

Flags:

- `--list` → List all playlists
- `--download` → Download all songs from a playlist
- `--output, -o` → Target directory for downloads (required for download)

### `songs` command

Manage songs inside a playlist.

**List songs in a playlist**

```bash
go run cmd/tunctl/main.go songs --list "<playlist name>"
```

Displays all tracks in the given playlist.

**Download specific songs by name**

```bash
go run cmd/tunctl/main.go songs --download "<playlist>" "Song Name 1" "Song Name 2" --output ./music
```

- Downloads only the specified songs
- Matches songs by name
- Requires an output directory

Flags:

- `--list` → List songs in a playlist
- `--download` → Download specific songs by name
- `--output, -o` → Target directory for downloads (required for download)

## Pending Work (Minimum TODO)

There is still a lot to do. At minimum:

- [ ] Validate user inputs
- [ ] Implement the TUI
- [ ] Allow choosing audio download format
- [ ] Check if a song already exists before downloading
- [ ] Improve Spotify permissions (currently limited scopes)
- [ ] Add metadata to downloaded files (artist, album, genres, etc.)
- [ ] Organize downloads into folders (artist / album / playlist)
- [ ] Improve error handling and overall CLI syntax

## Disclaimer

This is a learning project. Expect:

- Messy commits.
- Bugs are likely present.
- Nothing is stable or production-ready.

If you’re looking for a reliable Spotify manager, this is not it.
