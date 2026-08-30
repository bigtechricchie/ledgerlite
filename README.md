# LedgerLite CLI

LedgerLite CLI is a lightweight command-line banking ledger built in Go for recording financial transactions, calculating balances, and enforcing basic ledger invariants.

The project models a small transaction-based ledger with an emphasis on clear code, explicit business rules, predictable behaviour, and maintainable application design.

Built by Edward Siwek.

## Features

LedgerLite CLI currently supports:

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

LedgerLite CLI currently has no third-party runtime dependencies.

## Project Structure

```text
ledgerlite-cli/
├── .gitignore
├── README.md
├── go.mod
├── main.go
└── main_test.go
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

## Code Quality

Format the source code:

```bash
go fmt ./...
```

Run static analysis:

```bash
go vet ./...
```

Run the full test suite:

```bash
go test ./...
```

These checks should be run before committing changes.

## Building the Application

Build the executable from the project root:

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

## Current Status

LedgerLite CLI is under active development.

The current implementation provides an interactive in-memory ledger with
deposit and withdrawal recording, transaction lookup, derived balance
calculation, account and transaction validation, overdraft protection,
automatic transaction identifiers, and automated tests for core ledger rules.

## Disclaimer

LedgerLite CLI is a software portfolio project. All organisations, accounts,
transactions, identifiers, and financial data used in this repository are
fictional and provided solely for demonstration purposes. Any resemblance
to real entities or financial activity is coincidental.
