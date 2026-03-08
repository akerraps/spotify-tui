# tunectl – Terminal Music Downloader & Metadata Tool

**tunectl** is a personal learning project written in Go.

The goal of this project is to experiment with building a terminal-based tool that can fetch music metadata, download tracks, and organize a local music library.

It is designed as a sandbox to explore:

- Go project architecture
- CLI and terminal UX
- working with public APIs
- metadata processing
- integrating external tools like `yt-dlp`

This is **not a production-ready application**. Expect rough edges, experimental code, and frequent refactors.

---

# Project Direction

Originally, `tunectl` started as a Spotify playlist manager using the Spotify Web API.

Due to API limitations and access restrictions, the project has pivoted toward a more open approach:

- using **public music metadata databases**
- removing the dependency on Spotify
- allowing users to **manually query albums, artists, and tracks**
