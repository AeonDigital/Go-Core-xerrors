Go-Core-xerrors
================================

![Go Test Coverage](https://raw.githubusercontent.com/AeonDigital/Go-Core-xerrors/badges/.badges/main/coverage.svg)

> [Aeon Digital](http://aeondigital.com.br)  
> rianna@aeondigital.com.br

&nbsp;

> Structured, typed errors and terminal-friendly diagnostics for Go applications.

This package provides a lightweight error framework for separating validation failures,
operational failures, and CLI-oriented human-facing diagnostics. It is designed to
be predictable in logs, traces, and terminal output while keeping the dependency
surface limited to the Go standard library.




&nbsp;
________________________________________________________________________________

## INSTALLATION

Install the package from its current repository location:

```shell
go get github.com/AeonDigital/Go-Core-xerrors/pkg/xerrors@latest
```




&nbsp;
________________________________________________________________________________

## PURPOSE

`xerrors` standardizes error tracking and debugging ergonomics by offering three
focused error families:

- **Error 400 family (`IError400`)** — lightweight validation and client-side failures.
- **Error 500 family (`IError500`)** — operational/runtime failures with wrapping,
  component tracing, and debug metadata.
- **CLI error family (`IErrorCLI`)** — terminal-friendly errors with separate developer
  and user-facing messages.

Both structured error families also support immutable fluent helpers such as `WithArgs()`
and `WithCallerSkip()` for context enrichment.




&nbsp;
________________________________________________________________________________

## GLOBAL CONFIGURATION

The package includes atomic, thread-safe configuration helpers for toggling debug
rendering at runtime:

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
________________________________________________________________________________

## BASIC USAGE

### 1. User Validation Failures (Error 400)

Use `NewError400` for client-side or validation failures. It supports framework tokens,
domain-specific tokens, or plain formatted text.

```go
err := xerrors.NewError400(xerrors.XERR_FIELD_REQUIRED).WithArgs("email")

errPlain := xerrors.NewError400("invalid temporary session token: %s", tokenID)
```



&nbsp;
________________________________________________________________________________


### 2. Unexpected System Failures (Error 500)

Use `NewError500` for operational failures that should preserve an underlying cause
and runtime context.

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


When wrapping the error inside helper layers, you can shift the caller frame with
`WithCallerSkip`:

```go
return err.WithCallerSkip(1)
```



&nbsp;
________________________________________________________________________________


### 3. CLI and Terminal-Friendly Errors (Error CLI)

Use `NewErrorCLI` when you need separate messages for developers and end users, especially
in CLI flows or console tooling.

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
________________________________________________________________________________


### 4. Quick Debug Printing

For fast troubleshooting, `Print` sends the error text to `stderr` without requiring
a logger setup.

```go
xerrors.Print(err)
```




&nbsp;
________________________________________________________________________________

## EXTERNAL DEPENDENCIES

`xerrors` depends only on the Go standard library:

- `fmt`
- `os`
- `runtime`
- `strings`
- `sync`
- `sync/atomic`

No third-party modules or external frameworks are required.




&nbsp;
________________________________________________________________________________

## 4. ADDITIONAL INFORMATION

This project uses the [Semantic Versioning](https://semver.org/) system proposed
by **Tom Preston-Werner**.




&nbsp;
________________________________________________________________________________

## 5. LICENSE

This project is offered under the [MIT license](LICENSE.md).