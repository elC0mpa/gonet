# NetStats CLI

`NetStats CLI` is a cross-platform command-line tool for monitoring network usage statistics by application. This tool supports both macOS and Linux and is designed to be easy to build and extend.

## Features

- Real-time monitoring of network usage by application.
- Cross-platform support for macOS and Linux.
- Lightweight and simple to use.

## Installation

### macOS

#### Option 1: Homebrew (Recommended)

If you have Homebrew installed, you can install `NetStats CLI` using the following command:

```bash
brew install yourusername/tap/netstats-cli
```

#### Option 2: Download Binary from Releases

1. Go to the [Releases page](https://github.com/elC0mpa/netstats/releases).
2. Download the latest release binary for macOS (`netstats-darwin-amd64`).
3. Make the file executable:

   ```bash
   chmod +x netstats-darwin-amd64
   ```

4. (Optional) Move the file to `/usr/local/bin` for easy access:

   ```bash
   sudo mv netstats-darwin-amd64 /usr/local/bin/netstats
   ```

5. Run `netstats` from any terminal:

   ```bash
   netstats
   ```

### Linux

#### Option 1: Download Binary from Releases

1. Go to the [Releases page](https://github.com/elC0mpa/netstats/releases).
2. Download the latest release binary for Linux (`netstats-linux-amd64`).
3. Make the file executable:

   ```bash
   chmod +x netstats-linux-amd64
   ```

4. (Optional) Move the file to `/usr/local/bin` for easy access:

   ```bash
   sudo mv netstats-linux-amd64 /usr/local/bin/netstats
   ```

5. Run `netstats` from any terminal:

   ```bash
   netstats
   ```

#### Option 2: Build from Source

To build from source, make sure you have [Go](https://golang.org/dl/) installed:

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/netstats-cli.git
   cd netstats-cli
   ```

2. Build the binary:

   ```bash
   # For macOS
   GOOS=darwin GOARCH=amd64 go build -o netstats-darwin-amd64

   # For Linux
   GOOS=linux GOARCH=amd64 go build -o netstats-linux-amd64
   ```

3. Make the file executable and move it to a directory in your `$PATH`:

   ```bash
   chmod +x netstats-linux-amd64
   sudo mv netstats-linux-amd64 /usr/local/bin/netstats
   ```

4. Run `netstats`:

   ```bash
   netstats
   ```

## Usage

After installation, you can run `netstats` to start monitoring network usage.

```bash
netstats [options]
```

### Options

- `-a, --all`: Show all applications.
- `-f, --filter <name>`: Filter network usage by application name.

### Example

```bash
netstats -f chrome
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

```markdown
MIT License

Copyright (c) 2023 Jose Gabriel Companioni Benitez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
