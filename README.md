# Scurl

**Scurl (Simple-curl)** is a command-line tool for making HTTP requests. It provides a simplified interface for sending GET, HEAD, and POST requests with support for custom headers and data.

---

## Installation

### From Source
1. Clone the repository:
   ```bash
   git clone https://github.com/davidonlinearchive/scurl.git
   ```
2. Build the binary:
   ```bash
   go mod download
   go build
   ```
3. Run the tool:
   ```bash
   ./scurl [url]
   ```

---

## Usage

### Basic Usage
```bash
scurl [url]
```
This will send a **GET** request to the provided URL.
If no subcommand is specified, it defaults to **get**.

---

### Commands

#### get [url]
Send a **GET** request to the specified URL.

**Example:**
```bash
scurl get https://example.com
```

#### head [url]
Send a **HEAD** request to the specified URL.
This retrieves only the headers.

**Example:**
```bash
scurl head https://example.com
```

#### post [url]
Send a **POST** request with custom data and headers.

**Example:**
```bash
scurl post https://example.com/post -d '{"key":"value"}' -H "Content-Type: application/json"
```

---

### Description

| Flag | Long Form | Description | Example |
|------|------------|-------------|----------|
| `-d` | `--data` | HTTP POST data | `-d '{"key":"value"}'` |
| `-H` | `--header` | Set request headers | `-H "Content-Type: application/json"` |

---

## Features

- GET / HEAD / POST support
- Custom headers for requests
- Pretty-printed JSON/XML responses

---

## Notes

- URLs without `http://` or `https://` are automatically prefixed with `http://`.
- The tool uses the [`github.com/spf13/cobra`](https://github.com/spf13/cobra) library for CLI handling.
- All responses are printed to stdout with proper formatting.

---

## Contributing

Feel free to open issues or submit pull requests!
For feature requests or bug reports, please check the [GitHub repository](https://github.com/davidonlinearchive/scurl).
