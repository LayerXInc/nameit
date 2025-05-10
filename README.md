# nameit

A command-line tool for generating memorable names with different styles.

## Description

`nameit` is a name generator CLI tool that creates memorable names by combining adjectives and nouns. It supports different naming styles and allows customization through various options.

## Features

- Generate names in different styles/modes:
  - Modern style
  - Heroku-like style
- Customize with your own word lists
- Load custom adjectives and nouns from files
- Configuration via config file or environment variables

## Installation

```bash
go install github.com/layerxinc/nameit@latest
```

## Usage

Basic usage:

```bash
nameit
```

With options:

```bash
# Generate name
nameit

# Specify a count
nameit --count 5

# Specify a prefix
nameit --prefix my-corp

# Specify a random suffix of 5 hex characters
nameit --append-random --random-length 5 --random-chars 0123456789abcdef

# Use custom word lists
nameit --adjectives-file /path/to/adjectives.txt --nouns-file /path/to/nouns.txt
```

## Configuration

`nameit` supports configuration via:

1. Command-line flags
2. Environment variables
3. Config file: `.nameit.yaml` in your home directory
