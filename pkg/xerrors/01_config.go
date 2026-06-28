package xerrors

import (
	"sync/atomic"
)

/*
  ARCHITECTURE & SCOPE LIMITATION:
  configs.go acts as the single point of entry for parsing, validating, and
  loading application configuration settings (flags, environment variables, or files).

  Design Constraints:
  - This file must only expose the final configuration structures or reading mechanisms.
  - No domain logic or business orchestration is allowed within this package.
  - If a specific configuration subsystem (e.g., Database, CLI UI) grows complex,
    isolate its mapping into a dedicated struct file within this folder.
*/

// Insert configuration constants, structs and config functions below.

//
//
//

var (
	// debugMode uses an atomic int32 (0 for false, 1 for true)
	// to prevent concurrent data races during runtime switches.
	debugMode int32 = 0
)

// GetDebugMode returns the current runtime state of the debug flag.
//
// It exposes a thread-safe read of the package-wide debug toggle used by the
// error rendering engine to decide whether to show technical details.
func GetDebugMode() bool {
	return atomic.LoadInt32(&debugMode) == 1
}

// EnableDebugMode enables technical error details for the whole package.
//
// It stores a persistent runtime flag that causes formatted errors to include
// component traces and wrapped cause information during rendering.
func EnableDebugMode() {
	atomic.StoreInt32(&debugMode, 1)
}

// DisableDebugMode disables technical error details for the whole package.
//
// It resets the package-wide debug flag so formatted errors remain concise and
// user-oriented when the application is running in standard mode.
func DisableDebugMode() {
	atomic.StoreInt32(&debugMode, 0)
}

// ToggleDebugMode flips the global debug flag to its opposite state.
//
// It uses an atomic compare-and-swap loop to preserve thread safety while
// switching between technical and sanitized error output modes.
func ToggleDebugMode() {
	for {
		current := atomic.LoadInt32(&debugMode)
		var next int32 = 0
		if current == 0 {
			next = 1
		}

		// Performs a Compare-And-Swap loop to guarantee safe state alteration
		if atomic.CompareAndSwapInt32(&debugMode, current, next) {
			break
		}
	}
}
