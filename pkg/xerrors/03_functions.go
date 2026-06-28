package xerrors

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

/*
  ARCHITECTURE & SCOPE LIMITATION:
  functions.go groups stateless, decoupled utility behaviors and pure computational
  routines required to back up the central proposal of the package.

  Design Constraints:
  - Every function placed here should ideally be deterministic (same input produces same output).
  - No global state mutation or complex side-effects are allowed within these routines.
  - If routines grow complex or introduce stateful context, split them into dedicated files.
*/

// Insert standalone functions or mathematical algorithms below.

//
//
//

var callerFunc = runtime.Caller
var funcForPC = runtime.FuncForPC
var lastIndexFunc = strings.LastIndex

// Print writes the supplied error to the standard error stream.
//
// It acts as a lightweight helper for terminal diagnostics and emits the error
// text to os.Stderr when the provided value is non-nil.
func Print(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err.Error())
}

// TraceCallerLocation resolves a package-qualified function name from the runtime stack.
//
// It inspects the call stack using the supplied skip depth and returns a stable
// package::function identifier, falling back to safe "unknown" values when the
// runtime metadata is unavailable.
func TraceCallerLocation(skip int) string {
	// skip+1 balances out this internal utility function frame level allocation
	pc, _, _, ok := callerFunc(skip + 1)
	if !ok {
		return "unknown::unknown"
	}

	details := funcForPC(pc)
	if details == nil {
		return "unknown::unknown"
	}

	fullName := details.Name()
	lastDot := lastIndexFunc(fullName, ".")
	if lastDot == -1 {
		return "unknown::" + fullName
	}

	funcName := fullName[lastDot+1:]
	remainingPath := fullName[:lastDot]

	lastSlash := lastIndexFunc(remainingPath, "/")
	pkgName := remainingPath
	if lastSlash != -1 {
		pkgName = remainingPath[lastSlash+1:]
	}

	return pkgName + "::" + funcName
}

