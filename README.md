# Go Url Shortener

This is a simple URL shortener written in Go using the Gin web framework. It allows users to create short URLs that redirect to longer URLs.

## Installation
1. Clone the repository: git clone https://github.com/ghostdev-labs/GoUrlShortener.git
2. Install dependencies: go mod download

## Usage
1. Start the server: go run main.go
2. Navigate to http://localhost:8080 in your web browser
3. Enter a long URL and click "Shorten"
4. Copy the shortened URL and use it to redirect to the original long URL

## Configuration
The application uses environment variables for configuration. You can set these variables in a .env file in the root directory of the project. Here are the available variables:

* DB_CONNECTION_STRING: The connection string for the MySQL database. Example: DB_CONNECTION_STRING=db_user:db_password@tcp(db_host:db_port)/db_name?charset=utf8&parseTime=True&loc=Local
* PORT: The port number for the server to listen on. Example: PORT=8080

## Contributing
Contributions are welcome! Please open an issue or pull request if you find a bug or have a feature request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
