# `gonet`

`gonet` is a cross-platform CLI for real-time monitoring of network bandwidth usage by process name. It provides detailed network usage statistics for each application, supporting both macOS and Linux.

## Features

- Real-time monitoring of network usage by application.
- Cross-platform support for macOS and Linux.
- Lightweight and simple to use.

## Installation

### macOS

#### Option 1: Homebrew (Recommended)

If you have Homebrew installed, you can install `gonet` using the following command:

```bash
brew install https://github.com/elC0mpa/gonet
```

#### Option 2: Download Binary from Releases

1. Go to the [Releases page](https://github.com/elC0mpa/gonet/releases).
2. Download the latest release binary for macOS.
3. Make the file executable:

   ```bash
   chmod +x gonet-macos-amd64
   ```

4. (Optional) Move the file to `/usr/local/bin` for easy access:

   ```bash
   sudo mv gonet-macos-amd64 /usr/local/bin/gonet
   ```

5. Run `gonet` from any terminal:

   ```bash
   gonet
   ```

### Linux

#### Option 1: Download Binary from Releases

1. Go to the [Releases page](https://github.com/elC0mpa/gonet/releases).
2. Download the latest release binary for Linux (`gonet-linux-amd64`).
3. Make the file executable:

   ```bash
   chmod +x gonet-linux-amd64
   ```

4. (Optional) Move the file to `/usr/local/bin` for easy access:

   ```bash
   sudo mv gonet-linux-amd64 /usr/local/bin/gonet
   ```

5. Run `gonet` from any terminal:

   ```bash
   gonet
   ```

#### Option 2: Build from Source

To build from source, make sure you have [Go](https://golang.org/dl/) installed:

1. Clone the repository:

   ```bash
   git clone https://github.com/elC0mpa/gonet.git
   cd gonet-cli
   ```

2. Build the binary:

   ```bash
   # For macOS
   GOOS=darwin GOARCH=amd64 go build -o gonet-darwin-amd64

   # For Linux
   GOOS=linux GOARCH=amd64 go build -o gonet-linux-amd64
   ```

3. Make the file executable and move it to a directory in your `$PATH`:

   ```bash
   chmod +x gonet-linux-amd64
   sudo mv gonet-linux-amd64 /usr/local/bin/gonet
   ```

4. Run `gonet`:

   ```bash
   gonet
   ```

## Usage

After installation, you can run `gonet` to start monitoring network usage.

```bash
gonet [options]
```

### Options

- `-a, --all`: Show all applications.
- `-f, --filter <name>`: Filter network usage by application name.

### Example

```bash
gonet -f chrome
```

## Contributing

We welcome contributions to improve `gonet`! To ensure smooth collaboration, please follow these guidelines:

1. **Branching and Pull Requests**:

   - All pull requests (PRs) should be made to the `develop` branch, not `main`.
   - Make sure to create a new feature or bug-fix branch from `develop` and submit PRs from there (e.g., `feature/your-feature` or `fix/your-bug`).
   - Use descriptive names for branches to indicate the purpose of the branch.

2. **Code Quality**:

   - Write clear, concise, and well-documented code.
   - Run tests before submitting a PR to ensure that your changes don’t introduce any issues.
   - If you’re adding a new feature, consider including tests to cover your changes.

3. **Commit Messages**:

   - Follow conventional commit messages, e.g., `feat: add new monitoring filter` or `fix: resolve bug in CLI parsing`.
   - Write clear and concise messages that describe what your commit does.

4. **Review Process**:

   - Once your PR is submitted, it will undergo a review. Please be responsive to any feedback and make requested changes as necessary.
   - Avoid merging your own PRs; a maintainer will handle the merging once the PR is approved.

Thank you for contributing to `gonet` and helping to make it better!

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE.md) file for details.
