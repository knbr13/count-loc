# count-loc

`count-loc` is a fast and efficient Lines of Code (LOC) counter written in Go. It traverses directories concurrently and provides detailed statistics about code, comments, and blank lines across many programming languages.

## Features

- **Blazing Fast**: Uses a worker pool to process files concurrently.
- **Detailed Statistics**: Categorizes lines into Code, Comments, and Blank lines.
- **Extensive Language Support**: Supports over 40 programming languages.
- **Flexible Exclusions**: Exclude directories by name or files/directories by glob patterns.
- **Multiple Output Formats**: Supports default table, JSON, compact summary, and formatted table outputs.
- **Hidden File Support**: Optionally include hidden files and directories in the count.

## Installation

To install `count-loc`, you need to have Go installed on your system.

```bash
go install github.com/knbr13/count-loc@latest
```

Alternatively, you can clone the repository and build it manually:

```bash
git clone https://github.com/knbr13/count-loc.git
cd count-loc
go build -o count-loc .
```

## Usage

```bash
count-loc [options] [path]
```

### Options

- `-p, --path <path>`: Path to the directory to analyze (default: current directory).
- `-w, --workers <n>`: Number of worker goroutines (default: number of CPUs).
- `-H, --hidden`: Include hidden files and directories.
- `-f, --format <format>`: Output format: `default`, `json`, `compact`, `formatted`.
- `-x, --exclude <dirs>`: Comma-separated list of directories to exclude.
- `-i, --ignore <patterns>`: Comma-separated list of patterns to exclude files (e.g., `"*_test.go,*.log"`).
- `-e, --errors`: Show detailed error messages.
- `-v, --verbose`: Enable verbose output.
- `-q, --quiet`: Suppress non-essential output.
- `-V, --version`: Print version information.
- `-h, --help`: Print help message.

### Examples

```bash
# Count LOC in current directory
count-loc

# Count LOC in specified directory
count-loc /path/to/project

# Output results in JSON format
count-loc -f json .

# Use 8 workers and include hidden files
count-loc -w 8 -H .

# Exclude test and docs directories
count-loc -x "test,docs" .

# Exclude files matching patterns
count-loc -i "users_*.go,*log" .
```

## Supported Languages

`count-loc` supports a wide range of languages, including:

Go, JavaScript, TypeScript, Python, Java, C, C++, C#, Ruby, PHP, Swift, Kotlin, Rust, Scala, HTML, CSS, SCSS, SQL, Shell, YAML, JSON, Markdown, XML, Vue, Svelte, Lua, R, Perl, Elixir, Erlang, Haskell, Clojure, TOML, INI, Terraform, Protocol Buffers, GraphQL, Assembly, and more.
