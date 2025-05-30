# CKB Go Integration Test

This repository contains integration tests for the Nervos CKB blockchain using the Go SDK.

## Overview

The CKB Go Integration Test project provides a comprehensive testing framework for validating the functionality of the [Nervos CKB](https://www.nervos.org/) blockchain using the Go programming language. It leverages the CKB SDK for Go to interact with the CKB blockchain and perform various tests.

## Requirements

- Go 1.18 or higher
- Git

## Installation

```bash
# Clone the repository
git clone https://github.com/cryptape/ckb-go-integration-test.git
cd ckb-go-integration-test

# Install dependencies
go mod tidy
```

## Project Structure

```
├── ckb-sdk-go/       # Local copy of the CKB SDK for Go
├── sdk/              # Integration test files
│   ├── common.go     # Common utilities for tests
│   ├── newMultiSign_test.go  # Multisig transaction tests
│   ├── v106/         # Tests for CKB v1.06
│   ├── v107/         # Tests for CKB v1.07
│   └── v108/         # Tests for CKB v1.08
└── .github/workflows/  # CI/CD configuration
```

## Running Tests

### Standard Tests

```bash
cd sdk/
go test ./... -short
```

### Multisig Tests

To run the multisig transaction tests:

```bash
# Clone the required SDK repository with multisig support
git clone https://github.com/15168316096/ckb-sdk-go.git -b exec/upgrade-multisig

# Setup dependencies
cd ckb-sdk-go
go mod tidy
cd ..
go mod tidy

# Run the multisig tests
cd sdk/
go test -v ./newMultiSign_test.go
```

## CI/CD Workflows

This project uses GitHub Actions for continuous integration and testing.

### Standard Test Workflow

The standard test workflow runs on push and pull requests to the main branch. It:

1. Sets up Go 1.18
2. Checks out the code
3. Runs the standard test suite

### Multisig Test Workflow

The multisig test workflow runs on push, pull requests to the main branch, or can be triggered manually. It:

1. Sets up Go 1.18
2. Checks out the code
3. Clones the CKB SDK Go repository with multisig support
4. Sets up dependencies
5. Runs the multisig tests

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the MIT license.