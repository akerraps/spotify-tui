# tunectl – Terminal Spotify Playlist Manager (Learning Project)

tunectl is a personal project I’m building to learn Go.
The goal is to create a terminal-based user interface (TUI) that lets me interact with the Spotify Web API to explore and organize my personal music library in local.

This is not a production-ready tool. It’s an experimental sandbox where I try out Go, APIs, and different ways of handling data. The code is likely messy, incomplete, or broken in places — that’s part of the learning process.

## Project Goal

The main purpose is learning and experimentation. Specifically, I want to:

[x] Connect and authenticate with the Spotify Web API.
[x] Retrieve information about my account: playlists, tracks, artists, etc.
[] Export personal data (e.g., to CSV) for analysis or organization.
[] Provide tools in the TUI to sort, browse, and manage my music library in a convenient way.

>Note: All interactions are done through official Spotify API endpoints and within their terms of service. This project is intended for personal use and learning only.

## Usage / Getting Started

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

## Disclaimer

This is a learning project. Expect:

- Messy commits.
- Bugs are likely present.
- Nothing is stable or production-ready.

If you’re looking for a reliable Spotify manager, this is not it.

If you’re interested in following along as I learn Go and experiment with the Spotify API, you’re welcome to explore.
