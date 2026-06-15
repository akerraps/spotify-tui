# Tunectl – Terminal Music Downloader & Metadata Tool

`tunectl` is a personal learning project written in Go.

The goal of this project is to explore the design of a terminal-based tool capable of resolving music metadata and downloading audio tracks using external sources.

It is not intended to be a production-ready application. The codebase is experimental and subject to frequent changes and refactoring.

## Overview

tunectl is a terminal tool for downloading individual songs using yt-dlp, with optional metadata enrichment via external APIs.

It operates in two modes:

### API-enabled mode

When enabled, the input is enriched using music metadata services (such as MusicBrainz). This can normalize artist names, fix typos, and complete missing fields like genre before building the final download query.

The downside is that incorrect metadata resolution **may lead to downloading a different track than intended**, since the final match depends on external sources.

### No-API mode

When disabled, no metadata enrichment is performed. The raw input is used directly to build a search query in the form "<song> <artist> song", which is then passed to yt-dlp to download the first matching result.

## CLI Usage

The main command provided by tunectl is songs, which is used to download tracks directly from the terminal.

Basic usage:

```bash
tunectl songs "song;artist;genre"
```

Each input is a semicolon-separated string with the following structure:

song: track name
artist: artist or group name
genre: optional genre tag used for filtering or enrichment

Multiple songs can be provided in a single command:

```bash
songs tunectl "song1;artist1;genre1" "song2;artist2;genre2"
```

When multiple entries are passed, each one is processed independently.

### Metadata behavior

By default, tunectl runs in API-enabled mode.

In this mode:

- song and artist fields are enriched using external metadata sources
- misspellings and incomplete inputs may be corrected
- additional metadata may be resolved via MusicBrainz

If the --no-api flag is used, metadata enrichment is disabled.

In this case:

- the input is used directly to build the search query
- the query format becomes: "<song> <artist> song"
- genre is not used for filtering but is still appended to the final search query when present

Important note:

> [!WARNING]
> If API mode is enabled, all fields may be replaced by resolved metadata except genre, which is always preserved.

## Download behavior

All downloads are handled through `yt-dlp` as the backend extractor.

If no output directory is specified, files are saved to:

```
/home/<user>/Music
```

The output directory can be overridden via CLI flags.

## Cache management

`tunectl` stores the `yt-dlp` binary in the user’s local cache directory.

A cache management command is provided:

```bash
tunectl cache --clear
```

or

```bash
tunectl cache -c
```

This removes the cached `yt-dlp` binary.

### Purpose

This is intended as a recovery mechanism in cases where:

- the downloaded `yt-dlp` binary becomes corrupted
- upstream changes in `yt-dlp` require a fresh binary
- download failures occur due to outdated cached versions

On the next execution, `tunectl` will automatically re-download the latest available version of `yt-dlp` from the official repository.

## Future work

The project is still in early development. The following features are planned:

- [ ] Terminal UI (TUI): A terminal interface that becomes the default behavior when running tunectl without arguments.
- [ ] File input support: CSV, JSON
- [ ] Persistent cache (SQLite): The system will prefer local data over API calls when available, reducing external dependencies over time.
- [ ] Playlist downloads: Support downloading full playlists from YouTube URLs.
- [ ] Album downloads: Support resolving and downloading full albums using

## Notes

This project is primarily a learning exercise in Go, CLI design, and external tool integration.
Behavior is intentionally simple and favors direct execution over strict correctness guarantees.
Metadata resolution is heuristic and may produce incorrect matches in some cases.
