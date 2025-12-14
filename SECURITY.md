# Security Policy

## Supported Versions

We actively support the latest version of the go-dependency-injector library.

| Version | Supported          |
| ------- | ------------------ |
| Latest  | :white_check_mark: |
| < 1.0   | :x:                |

## Reporting a Vulnerability

If you discover a security vulnerability, please do **NOT** open a public issue.

Instead, please email security concerns to the maintainers or use GitHub's private vulnerability reporting feature.

### What to Include

When reporting a vulnerability, please include:

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fix (if any)

### Response Timeline

- We will acknowledge receipt of your vulnerability report within 48 hours
- We will provide a more detailed response within 7 days
- We will work on a fix and keep you updated on the progress
- Once a fix is ready, we will coordinate disclosure

## Security Best Practices

When using this library:

1. Always use the latest version
2. Keep your Go version up to date
3. Review your dependency registrations carefully
4. Be cautious with factory functions that have side effects
5. Use appropriate lifetimes for your dependencies

## Known Security Considerations

- This library uses reflection for type resolution - ensure your factory functions are trusted
- Circular dependency detection helps prevent infinite loops
- Thread-safety is provided via mutexes - be aware of potential deadlocks in complex scenarios

Thank you for helping keep go-dependency-injector secure!


