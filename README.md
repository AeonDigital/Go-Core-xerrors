Go-Core-xerrors
================================

> [Aeon Digital](http://www.aeondigital.com.br)  
> rianna@aeondigital.com.br

&nbsp;

> Structured, typed errors and terminal-friendly diagnostics for Go applications.

This package provides a lightweight error framework for separating validation failures, operational failures, and CLI-oriented human-facing diagnostics. It is designed to be predictable in logs, traces, and terminal output while keeping the dependency surface limited to the Go standard library.

The repository has been moved to its own standalone module and now exposes the public API through the package path under the module root.


&nbsp;
&nbsp;


________________________________________________________________________________

## Purpose

`xerrors` standardizes error tracking and debugging ergonomics by offering three focused error families:

*   **Error 400 family (`IError400`)** — lightweight validation and client-side failures.
*   **Error 500 family (`IError500`)** — operational/runtime failures with wrapping, component tracing, and debug metadata.
*   **CLI error family (`IErrorCLI`)** — terminal-friendly errors with separate developer and user-facing messages.

Both structured error families also support immutable fluent helpers such as `WithArgs()` and `WithCallerSkip()` for context enrichment.


&nbsp;
&nbsp;


________________________________________________________________________________

## Installation

Install the package from its current repository location:

```shell
go get github.com/AeonDigital/Go-Core-xerrors/pkg/xerrors@latest
```

Import it in your code:

```go
import (
    "errors"

    "github.com/AeonDigital/Go-Core-xerrors/pkg/xerrors"
)
```


&nbsp;
&nbsp;


________________________________________________________________________________

## Global Configuration

The package includes atomic, thread-safe configuration helpers for toggling debug rendering at runtime:

```go
// Enable technical layout extensions such as component tracking and root-cause dumps
xerrors.EnableDebugMode()

// Fallback to sanitized, user-friendly messages
xerrors.DisableDebugMode()

// Check or alternate the current state atomically
isEnabled := xerrors.GetDebugMode()
xerrors.ToggleDebugMode()
```


&nbsp;
&nbsp;


________________________________________________________________________________

## Basic Usage

### 1. User Validation Failures (Error 400)
Use `NewError400` for client-side or validation failures. It supports framework tokens, domain-specific tokens, or plain formatted text.

```go
err := xerrors.NewError400(xerrors.XERR_FIELD_REQUIRED).WithArgs("email")

errPlain := xerrors.NewError400("invalid temporary session token: %s", tokenID)
```


&nbsp;


### 2. Unexpected System Failures (Error 500)
Use `NewError500` for operational failures that should preserve an underlying cause and runtime context.

```go
dbErr := errors.New("connection timeout downstream")

richErr := xerrors.NewError500(
    xerrors.XERR_PKGCTX,
    xerrors.XERR_UNKNOWN,
    dbErr,
    "database repository failure",
    `{"retry_count": 3}`,
).WithArgs("user_id_123")
```

When wrapping the error inside helper layers, you can shift the caller frame with `WithCallerSkip`:

```go
return err.WithCallerSkip(1)
```


&nbsp;


### 3. CLI and Terminal-Friendly Errors (Error CLI)
Use `NewErrorCLI` when you need separate messages for developers and end users, especially in CLI flows or console tooling.

```go
cliErr := xerrors.NewErrorCLI().
    SetDevMessage("database connection failed")

cliErr = cliErr.
    SetUserMessage("please try again later")

fmt.Println(cliErr.Error())
fmt.Println(cliErr.GetUserMessage())
```

You can also append content progressively:

```go
cliErr = xerrors.NewErrorCLI().
    AppendLNDevMessage("step 1 failed")
    .AppendLNDevMessage("step 2 failed")
```


&nbsp;


### 4. Quick Debug Printing

For fast troubleshooting, `Print` sends the error text to `stderr` without requiring a logger setup.

```go
xerrors.Print(err)
```


&nbsp;
&nbsp;


________________________________________________________________________________

## External Dependencies

`xerrors` depends only on the Go standard library:

*   `fmt`
*   `os`
*   `runtime`
*   `strings`
*   `sync`
*   `sync/atomic`

No third-party modules or external frameworks are required.


&nbsp;
&nbsp;


________________________________________________________________________________

## Licence

This project is offered under the [MIT license](LICENCE.md).