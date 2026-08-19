# LedgerLite

LedgerLite is a lightweight banking ledger CLI built in Go for recording financial transactions, calculating balances, and enforcing basic ledger invariants.

The project models a small transaction-based ledger with an emphasis on clear code, explicit business rules, predictable behaviour, and maintainable application design.

## Features

LedgerLite will progressively support:

* Recording deposits and withdrawals
* Maintaining transaction history
* Calculating account balances from ledger entries
* Validating transaction input
* Preventing invalid financial operations
* Searching for transactions
* Displaying account and ledger information through a command-line interface
* Handling errors explicitly and predictably

## Design Principles

* **Transactions are the source of truth** - balances are derived from ledger activity.
* **Financial rules are explicit** - invalid operations are rejected through validation.
* **Keep the design simple** - abstractions are introduced only when they provide a clear benefit.
* **Prefer the standard library** - external dependencies are added only when justified.

## Technology

* Go
* Go standard library
* Git

LedgerLite currently has no third-party runtime dependencies.

## Project Structure

```text
ledgerlite/
├── .gitignore
├── README.md
├── go.mod
└── main.go
```

The structure will evolve as application responsibilities become clearer.

## Requirements

A supported version of Go must be installed.

Check your installation with:

```bash
go version
```

## Running the Application

From the project root:

```bash
go run .
```

## Building the Application

Build the executable with:

```bash
go build
```

On Linux or macOS:

```bash
./ledgerlite
```

On Windows PowerShell:

```powershell
.\ledgerlite.exe
```

## Code Formatting

Format the project using the standard Go formatter:

```bash
go fmt ./...
```

## Current Status

LedgerLite is under active development.

The current implementation establishes the Go module and executable application entry point. Transaction modelling, ledger behaviour, validation, and command-line operations will be introduced incrementally.
