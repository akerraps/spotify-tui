# `tunectl` – Terminal Music Downloader & Metadata Tool

`tunectl` is a personal learning project written in Go.

The goal of this project is to explore the design of a terminal-based tool capable of resolving music metadata and downloading audio tracks using external sources.

It is not intended to be a production-ready application. The codebase is experimental and subject to frequent changes and refactoring.

## Overview

`tunectl` allows downloading individual songs from the terminal, optionally using a metadata API to improve search accuracy.

The tool operates in two modes:

### 1. API-enabled mode

When API usage is enabled, `tunectl` attempts to enrich the user input using external music metadata sources.

This includes:

- Normalizing artist names
- Attempting to correct typos in track or artist names
- Filling missing metadata such as genres

After resolution, the final query is passed to the downloader.

Important limitation:

> If the metadata resolution produces an incorrect match, the tool may download a different track than the one intended by the user, since the final resolution is delegated to `yt-dlp`.

### 2. No-API mode

When API usage is disabled, `tunectl` bypasses all metadata enrichment.

The input provided by the user is used directly to construct a search query in the form:

```
"<song> <artist> song"
```

The resulting query is passed directly to `yt-dlp`, which downloads the first matching result found.

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

## Future work / Pending features

The project is still in early development and several major features are planned but not yet implemented.

### Terminal UI (TUI)

A terminal user interface is planned, which will be launched when running tunectl without any arguments.

Although this is part of the long-term direction of the project, it is currently not the main focus. There are several core features that must be improved before the TUI becomes a priority.

### File-based input

Support will be added for reading track data from external files, including:

- CSV files
- JSON files

This will allow batch processing of songs and better integration with external tools or playlists.

### Cache improvements

A more advanced caching system is planned, replacing the current lightweight cache approach.

The goal is to introduce a local SQLite database that will store:

- artist metadata
- track metadata
- previously resolved queries

This will significantly reduce repeated API calls over time.

The intended behavior is:

- initial requests rely on external API lookups
- results are stored locally in SQLite
- subsequent requests prefer local data over API calls
- API usage gradually decreases as the local database becomes more complete

This component is considered a key part of the long-term architecture of the project.

### Playlist and batch downloads

Support will be added for downloading full playlists from YouTube via links.

This will allow batch ingestion of multiple tracks in a single operation.

### Album downloads

The tool will also support downloading full albums by providing:

- album name
- artist name

The system will attempt to resolve the album and download all associated tracks automatically.

## Notes

This project is primarily a learning exercise in Go, CLI design, and external tool integration.
Behavior is intentionally simple and favors direct execution over strict correctness guarantees.
Metadata resolution is heuristic and may produce incorrect matches in some cases.
