# Scurl


**Scurl (Simple-curl)** is a command-line tool for making HTTP requests. It provides a simplified interface for sending GET, HEAD, and POST requests with support for custom headers and data.

---

## Installation

### From Source
1. Clone the repository:
   ```bash
   git clone https://github.com/davidonlinearchive/scurl.git
Build the binary:
go mod download
go build
Run the tool:
./scurl [url]
 Usage
Basic Usage
scurl [url]
This will send a GET request to the provided URL. If no subcommand is specified, it defaults to get.

Commands

get [url] - Send a GET request to the specified URL.

Example:
scurl get https://example.com
head [url]
Send a HEAD request to the specified URL. This retrieves only the headers.

Example:
scurl head https://example.com
post [url]
Send a POST request with custom data and headers.

Example:
scurl post https://example.com/post -d '{"key":"value"}' -H "Content-Type: application/json"
    	Description
-d, --data	HTTP POST data (e.g., {"key":"value"})
-H, --header	Set request headers (e.g., -H "Content-Type: application/json")

Features

GET/HEAD/POST support

Custom headers for requests

Pretty-printed JSON/XML responses

Notes

URLs without http:// or https:// are automatically prefixed with http://.
The tool uses the github.com/spf13/cobra library for CLI handling.
All responses are printed to stdout with proper formatting.

Contributing

Feel free to open issues or submit pull requests!
For feature requests or bug reports, please check the GitHub repository.
