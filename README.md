# trim

A lightweight Go application to remove whitespace from stdin, depending on the executable's name (`trim`, `ltrim`, `rtrim`).

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Additional Information](#additional-information)
- [License](#license)

## Installation

### Prerequisites

- [Go](https://golang.org/) (1.16 or newer)

### Building from Source

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/jordiroca/trim.git
cd trim
```

Build the binaries:

```bash
go build -o trim.go
cp trim ltrim
cp trim rtrim

# Copy the programs under your PATH:

sudo mv trim ltrim rtrim /usr/local/bin/
```

## Usage

Use the respective binary to trim input from the command line.

Examples:

### trim

Remove leading and trailing whitespace:

```bash
echo -e "  hello \n world  " | trim
```

### ltrim

Remove leading whitespace:

```bash
echo -e "  hello \n world  " | ltrim
```

### rtrim

Remove trailing whitespace:

```bash
echo -e "  hello \n world  " | rtrim
```

Note: Ensure that the name of the executable binary (`trim`, `ltrim`, `rtrim`) matches the desired trimming operation.

## Additional Information

### Contributing

Feel free to submit PRs, issues, or any other feedback! You'll be welcomed and appreciated.

### Support

For issues using `trim`, please create a [GitHub issue](https://github.com/jordiroca/trim/issues).

## License

See [LICENSE](LICENSE) for more information.
