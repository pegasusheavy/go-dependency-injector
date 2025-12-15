# Contributing to Go Dependency Injector

Thank you for your interest in contributing! We welcome contributions from the community.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/go-dependency-injector.git`
3. Create a feature branch: `git checkout -b feature/my-new-feature`
4. Make your changes
5. Run tests: `go test ./...`
6. Run linter: `golangci-lint run`
7. Commit your changes: `git commit -am 'Add some feature'`
8. Push to the branch: `git push origin feature/my-new-feature`
9. Create a Pull Request

## Development Setup

### Prerequisites

- Go 1.22 or later
- golangci-lint (optional, for linting)

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# Run tests with race detection
go test -race ./...

# Run benchmarks
go test -bench=. -benchmem ./...
```

### Code Style

- Follow standard Go conventions
- Run `gofmt` and `goimports` on your code
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and concise

## Contribution Guidelines

### Bug Reports

When filing a bug report, please include:

- A clear description of the issue
- Steps to reproduce
- Expected behavior
- Actual behavior
- Go version
- Any relevant code samples

### Feature Requests

For feature requests, please:

- Describe the feature and its use case
- Explain why it would be valuable
- Provide examples of how it would be used

### Pull Requests

Good pull requests include:

- A clear description of what the PR does
- Tests for new functionality
- Documentation updates for new features
- No breaking changes without discussion
- Clean, focused commits

### Testing

- All new code should have tests
- Aim for high test coverage
- Include both positive and negative test cases
- Test edge cases

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Help others learn and grow

## Questions?

Feel free to open an issue for any questions or discussions!

## License

By contributing, you agree that your contributions will be licensed under the MIT License.



