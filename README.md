# tunectl – Terminal Spotify Playlist Manager (Learning Project)

TuneCtl is a personal learning project written in Go.  
The goal is to explore different ways of interacting with the Spotify Web API from the terminal, using both a classic CLI and (eventually) a full TUI.

This is not a production-ready tool. It’s an experimental sandbox where I learn Go, project structure, APIs, and terminal UX. The code may be messy, incomplete, or broken in places — that’s intentional and part of the process.

---

## Project Goal

The main purpose of this project is learning and experimentation. In particular, I want to:

- [x] Authenticate with the Spotify Web API
- [x] Fetch user playlists (with pagination support)
- [x] Fetch all tracks from large playlists
- [x] Download songs locally (yt-dlp)
- [x] Export playlist data to CSV / JSON
- [ ] Add a terminal UI (TUI) using Bubble Tea
- [ ] Improve error handling and UX
- [ ] Experiment with clean project architecture in Go

> Note: All interactions are done through official Spotify Web API endpoints and within Spotify’s terms of service. This project is for personal use and learning only.

---

## Music Downloading (yt-dlp Integration)

Song downloading is implemented using **yt-dlp**:

- Repository: https://github.com/yt-dlp/yt-dlp

TuneCtl acts as a thin wrapper around `yt-dlp` to fetch audio from YouTube based on track metadata.

### Why yt-dlp?

- Actively maintained
- Adapts quickly to YouTube changes
- Avoids reinventing a complex downloader
- Keeps TuneCtl focused on orchestration and learning

---

## Project Structure


```
.
├── cmd
│ └── tunctl
│ └── main.go # Entry point
├── internal
│ ├── cache # Cache management (yt-dlp, future API cache)
│ │ └── cache.go
│ ├── cli # CLI layer (urfave/cli)
│ │ └── cli.go
│ ├── core # Core logic
│ │ ├── authenticate.go
│ │ ├── core.go
│ │ ├── export.go # CSV / JSON export
│ │ └── playlists.go # Spotify API logic
│ ├── fetcher # yt-dlp wrapper
│ │ └── fetcher.go
│ ├── tui # Future TUI (Bubble Tea)
│ │ └── tui.go
│ └── types # Shared types
│ └── types.go
├── go.mod
├── go.sum
└── README.md
```

---

## Architecture Philosophy

- **core** → business logic (Spotify, data processing)
- **cli** → commands, flags, argument parsing
- **fetcher** → audio downloading (yt-dlp)
- **types** → shared structs
- **cache** → cache handling (extensible)
- **tui** → future UI layer

This separation allows reuse of core logic across CLI and TUI.

---

## Getting Started

1. Go to Spotify Developer Dashboard: https://developer.spotify.com/
2. Create an app
3. Get your credentials

Create a `.env` file:

```
SPOTIFY_ID=your_client_id_here
SPOTIFY_SECRET=your_client_secret_here
```

4. Run tunectl from your terminal. The application will use these credentials to connect to your Spotify account.

>Note: This project only accesses your own Spotify data for personal purposes. Use it at your own risk.

## Usage

### General help

```bash
go run cmd/tunctl/main.go -h
```


### `playlists` command

Manage Spotify playlists.

>Note: Every search by playlist name it is case sensitive

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

### Export playlist to file

```bash
# CSV
go run cmd/tunctl/main.go songs -l "<playlist>" -o songs.csv

# JSON
go run cmd/tunctl/main.go songs -l "<playlist>" -o songs.json
```
- Format is auto-detected by extension
- Supports large playlists (pagination handled internally)

**Download specific songs by name and artist**

```bash
go run cmd/tunctl/main.go songs --download "Song Name 1" "Song Name 2;Artist 2" --output ./music
```

- Downloads only the specified songs
- Matches songs by name
- If ";" used, uses the second part to search for the artist
- Requires an output directory

Flags:

- `--list` → List songs in a playlist
- `--download` → Download specific songs by name
- `--output, -o` → Target directory for downloads (required for download)

### `cache` command

Clears cache.

```bash
go run cmd/tunctl/main.go cache --clear
```

- Deletes previously downloaded yt-dlp
- TODO: When API calls are cached, clear them

## Pending Work (Minimum TODO)

There is still a lot to do. At minimum:

- [x] Validate user inputs
- [ ] Implement the TUI
- [ ] Allow choosing audio download format
- [x] Check if a song already exists before downloading
- [x] Improve Spotify permissions (currently limited scopes)
- [x] Add metadata to downloaded files (artist, album, etc.)
- [x] Organize downloads into folders (artist / album / playlist)
- [x] Improve error handling and overall CLI syntax

## Disclaimer

This is a learning project. Expect:

- Messy commits.
- Bugs are likely present.
- Nothing is stable or production-ready.

If you’re looking for a reliable Spotify manager, this is not it.
