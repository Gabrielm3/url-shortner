# Go URL Shortener

This is a project that showcases how to create a URL shortener in Go.

Utilizing the following technologies and libraries:

- [Go](https://golang.org/): The programming language used.
- [net/http](https://golang.org/pkg/net/http/): Go's standard package for building HTTP servers.
- [html/template](https://golang.org/pkg/html/template/): Go's standard package for HTML template rendering.

## Project Structure

The project is organized as follows:

- `main.go`: The main code for the API server.
- `handlers`: Folder containing HTTP request handlers.
- `services`: Folder containing business logic.
- `templates`: Folder containing HTML templates.

## Prerequisites

Before getting started, you must have the following prerequisites installed:

- Go: [Go Installation](https://golang.org/doc/install)

## How to Set Up and Run

Follow these steps to set up and run the API:

1. Clone the repository:

```
git clone https://github.com/Gabrielm3/url-shortner.git
```

2. Navigate to the project folder:

```
cd url-url-shortner
```

3. Run the project:

```
go run main.go
```

## Usage

After successful execution, you can use the **URL Shortner**. Simply open `http://localhost:8080` in your browser and enter the URL you want to shorten.

## Contributing

Contributions are welcome! If you find a bug or have an enhancement, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
