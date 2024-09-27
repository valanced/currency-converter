# Currency Converter Project

This project is a simple Currency Converter CLI application built in Go, designed to interact with external services for fetching conversion rates and performing currency conversions. The project utilizes dependency injection and clean architecture principles to maintain separation of concerns and ensure testability.

### Requirements

    •	Go 1.18 or later
    •	gomock
    •	golangci-lint

### Setup

1. Clone the repository:
```bash
git clone https://github.com/valanced/currency-converter.git
cd currency-converter 
```
2. Install dependencies:
```bash
go mod tidy
```
3. Run the application:
```bash 
COINMARKETCAP_API_KEY=xxx /bin/currency-converter 12 BTC USD
```

This will execute all unit tests in the project and provide a detailed output of the results.

Future Improvements

1.	Linter Setup:
It’s important to integrate a linter like golangci-lint to enforce code quality and style guidelines across the codebase.
2.	Configuration:
The application currently uses a basic .env file for configuration, but it should be replaced with a more robust configuration management system (e.g., viper or envconfig) to support different environments and settings.
3.	Makefile:
The current Makefile is a preliminary version, providing basic commands for building and testing. It should be expanded and refined to better suit a production environment, including commands for deployment, linting, and advanced testing scenarios.
