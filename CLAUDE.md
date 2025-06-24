# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`slp` (sleep with progress) is a command-line tool that provides a visual progress bar while sleeping for a specified duration. It's a modern replacement for the traditional `sleep` command with a rich terminal UI.

## Development Commands

### Environment Setup
```bash
# Install development tools using mise
mise install
```

### Build
```bash
go build .
```

### Lint
```bash
golangci-lint run
```

### Check Release Configuration
```bash
goreleaser check
```

### Update Dependencies
```bash
go mod tidy
```

### Test Release Locally
```bash
goreleaser release --snapshot --clean
```

## Architecture

### Project Structure
- `cmd/`: CLI command implementation
  - `root.go`: Command-line parsing and flag definitions using Cobra
  - `slp.go`: Progress bar implementation using Bubble Tea TUI framework
- `main.go`: Entry point

### Key Dependencies
- **Bubble Tea** (`github.com/charmbracelet/bubbletea`): Terminal UI framework for the progress bar
- **Bubbles** (`github.com/charmbracelet/bubbles`): Pre-built progress bar component
- **Cobra** (`github.com/spf13/cobra`): CLI framework for command parsing
- **Lipgloss** (`github.com/charmbracelet/lipgloss`): Terminal styling

### CLI Design
- Uses Cobra for command-line argument parsing
- Supports sleep-compatible time duration formats (seconds, minutes, hours)
- Progress bar customization through flags (colors, gradients, emojis, beep)

## CI/CD Pipeline

### GitHub Actions Workflows
- **CI** (`ci.yml`): Runs on every push
  - Build verification: `go build .`
  - Linting: `golangci-lint run`
  - Release config check: `goreleaser check`
- **Release** (`release.yml`): Runs on main branch pushes
  - Uses Release Please for automated semantic versioning
  - Builds cross-platform binaries with goreleaser
  - Publishes to Homebrew tap

### Release Process
1. Develop features in feature branches
2. Merge to main triggers Release Please to create/update a release PR
3. Merging the release PR creates a tag and triggers goreleaser
4. Goreleaser builds binaries for multiple platforms and updates Homebrew tap

## Development Notes

- Go version: 1.20
- No test files currently exist in the project
- When adding new features:
  - Define flags in `cmd/root.go`
  - Implement logic in `cmd/slp.go`
  - Follow Bubble Tea's Model-Update-View pattern for UI changes
- Always run `golangci-lint run` before committing
- The project uses mise for tool version management