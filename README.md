# tunectl – Terminal Spotify Playlist Manager (Learning Project)

TuneCtl is a personal learning project written in Go.
The goal is to explore different ways of interacting with the Spotify Web API from the terminal, using both a classic CLI and (eventually) a full TUI.

This is not a production-ready tool. It’s an experimental sandbox where I learn Go, project structure, APIs, and terminal UX. The code may be messy, incomplete, or broken in places — that’s intentional and part of the process.

## Project Goal

The main purpose of this project is learning and experimentation. In particular, I want to:

- [x] Authenticate with the Spotify Web API
- [x] Fetch user playlists
- [x] Fetch tracks from a specific playlist
- [ ] Add a terminal UI (TUI) using Bubble Tea
- [ ] Export personal data (CSV / JSON)
- [ ] Add local caching to avoid repeated API calls
- [ ] Improve error handling and UX
- [ ] Experiment with clean project architecture in Go
- [ ] Be able to download songs locally

>Note: All interactions are done through official Spotify Web API endpoints and within Spotify’s terms of service. This project is for personal use and learning only.
>
## Project Structure (High-Level)

```
cmd/
  tunctl/        → Application entrypoint (main)
internal/
  core/          → Business logic (Spotify auth, playlists, tracks)
  cli/           → CLI layer (urfave/cli commands)
```

- core: does the real work (Spotify API, data handling)
- cli: wires commands and arguments to the core
- cmd/tunctl: decides how the app runs (CLI vs TUI)

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

### Available Commands

**List playlists**

```bash
go run cmd/tunctl/main.go playlists
```

Outputs all playlists associated with your Spotify account.

**List songs from a playlist**

```bash
go run cmd/tunctl/main.go songs "<playlist name>"
```

## Planned TUI Behavior

In the future:

- Running tunectl with no arguments will launch a TUI
- Running it with arguments will keep using the CLI

## Disclaimer

This is a learning project. Expect:

- Messy commits.
- Bugs are likely present.
- Nothing is stable or production-ready.

If you’re looking for a reliable Spotify manager, this is not it.
