# Examples

Runnable examples demonstrating `go-xgen` features.

## Table of Contents

| Example | Description | Run |
| ------- | ----------- | --- |
| [generator](./generator/) | ID generation, UUID, API keys | `cd generator && go run main.go` |
| [hash](./hash/) | Password hashing with HMAC-SHA256 + bcrypt | `cd hash && go run main.go` |
| [signature](./signature/) | HMAC-SHA256 request signing & verification | `cd signature && go run main.go` |

## Quick Start

```bash
# Clone the repository
git clone https://github.com/hotfixfirst/go-xgen.git
cd go-xgen/_examples

# Run a specific example
cd generator && go run main.go
```

## Adding New Examples

When adding a new feature example:

1. Create a new directory: `_examples/{feature}/`
2. Add `main.go` with runnable code
3. Add `README.md` with documentation
4. Update this file's table of contents
