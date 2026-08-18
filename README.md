# tunectl

Terminal music downloader written in Go.

Uses `yt-dlp` for downloading and MusicBrainz for optional metadata lookup.

## Installation

Build the binary:

```bash
    go build -o tunectl ./cmd/tunectl
```

Or run it directly:

```bash
    go run ./cmd/tunectl
```

## Configuration

Configuration can be provided through a `.env` file or CLI options.

Create a `.env` file in the working directory:

```env
TUNECTL_OUTPUT_DIR=/home/user/Music
TUNECTL_NO_API=false
TUNECTL_PARALLELISM=1
```

CLI options override environment variables.

### Options

| Option         | Alias | Description                         | Default   |
| -------------- | ----- | ----------------------------------- | --------- |
| `--output`     | `-o`  | Output directory                    | `~/Music` |
| `--no-api`     |       | Disable MusicBrainz metadata lookup | `false`   |
| `--paralelism` | `-n`  | Number of parallel downloads        | `1`       |

## Download songs

The main command is `songs`.

```bash
    tunectl songs "Title;Artist"
```

A genre can optionally be specified:

```bash
    tunectl songs "Title;Artist;Genre"
```

Multiple artists can be specified as a comma-separated list:

```bash
    tunectl songs "Title;Artist 1, Artist 2"
```

Multiple songs can be downloaded in a single command:

```bash
    tunectl songs \
      "Song 1;Artist 1" \
      "Song 2;Artist 2" \
      "Song 3;Artist 3"
```

### Examples

Download a song:

```bash
    tunectl songs "Alchemy;Philip Sayce"
```

Specify a genre:

```bash
    tunectl songs "Alchemy;Philip Sayce;Blues"
```

Specify an output directory:

```bash
    tunectl songs -o ~/Music "Alchemy;Philip Sayce"
```

Download multiple tracks in parallel:

```bash
    tunectl songs -n 4 \
      "Song 1;Artist 1" \
      "Song 2;Artist 2"
```

## Metadata lookup

By default, `tunectl` uses the MusicBrainz API to resolve track metadata before downloading.

For example:

```bash
    tunectl songs "Alchemy;Philip sayce"
```

MusicBrainz can resolve the artist and track information before generating the download query.

To disable metadata lookup:

```bash
    tunectl songs --no-api "Alchemy;Philip Sayce"
```

When the API is disabled, the original input is used directly.

> [!WARNING]
> MusicBrainz resolution is not guaranteed to return the intended recording. An incorrect match can result in downloading a different track.

## File input

Tracks can also be loaded from CSV or JSON files using the `file` command.

### CSV

```bash
    tunectl file --csv --data songs.csv
```

Example:

```csv
    Title,Artists,Album,Genre
    What the Funk,"Gustavo Mota, Naizon",What the Funk,Electronic
```

Additional columns are ignored.

### JSON

    tunectl file --json --data songs.json

Example:

```json
[
  {
    "Title": "Title",
    "Artists": ["Artist 1", "Artist 2"],
    "Album": "Album",
    "Genre": "Rock"
  }
]
```

Genre is optional and is used as part of the track metadata.

## Cache

`tunectl` automatically downloads and caches the `yt-dlp` binary.

Clear the cached binary with:

```bash
    tunectl cache --clear
```

The next execution will download `yt-dlp` again.

This can be useful if the cached binary is corrupted or outdated.

## Debugging

Enable debug logs with:

```bash
    tunectl --debug songs "Alchemy;Philip Sayce"
```

This displays information about configuration, metadata lookup, downloads, and other internal operations.

## Command reference

### `songs`

Download one or more songs.

```
    tunectl songs [OPTIONS] SONG...
```

### `file`

Download songs from a CSV or JSON file.

```
    tunectl file --csv --data FILE
    tunectl file --json --data FILE
```

### `cache`

Manage the local `yt-dlp` cache.

```
    tunectl cache --clear
```
