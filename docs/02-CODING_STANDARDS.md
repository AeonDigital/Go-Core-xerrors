CODING STANDARDS
================================

## 1. DOCUMENTATION STANDARD (HUMAN READABILITY FIRST)

### 1.1 General Constraints

* Language: 100% technical English for all code documentation, comments, and structure naming.
* Format: Avoid dense, blocky, inline comments. Use generous visual spacing, line breaks, and clear indentation so developers can scan the file effortlessly.


&nbsp;


### 1.2 Code Coverage Requirements

* Functions & Methods: 100% Mandatory documentation coverage. The only exception is trivial, obvious Getters/Setters with zero logic.
* Structs & Fields: Highly Preferred. Omit comments only if the field is standard and completely self-explanatory (e.g., ID string, CreatedAt time.Time). If any business rule applies, document it.


&nbsp;


### 1.3 Function Comment Anatomy

Every function documentation must follow a strict tiered format:

   1. Line 1 (Summary): A brief, single-line summary of what the function accomplishes.
   2. Line 2: A mandatory empty line (visual breathing space).
   3. Detailed View: Technical explanation detailing arguments, return values, special side-effects, and operational constraints.


&nbsp;
&nbsp;


________________________________________________________________________________

## Error & Panic Documentation Criteria:

* Panics: All potential application-halting execution paths (panic) must be explicitly declared.
* Return Errors (Contextual): If a failure is obvious (e.g., simple validation), no further explanation is needed. If a function is complex and can fail due to multiple distinct natures (e.g., network timeout vs. data corruption), those failure natures must be explicitly itemized.


&nbsp;
&nbsp;


________________________________________________________________________________

## 2. CODE EXAMPLE (THE GOLD STANDARD)

``` example.go
package fn

// FormatBytes converts an integer byte count into a string.
//
// Arguments:
//   - bytes: The raw payload size in bytes. Must be a positive integer.
//
// Returns:
//   - string: The formatted value (e.g., "1.5 GB", "42 MB").
//
// Error & Panic Natures:
//   - Panics: Will trigger a panic if the input bytes argument is a negative value.
//   - Complex Errors: Returns an empty string if the internal system math precision 
//     overflows during calculation.
func FormatBytes(bytes int64) string {
    if bytes < 0 {
        panic("fn: bytes argument cannot be negative")
    }
    // Implementation goes here
    return ""
}
```


&nbsp;
&nbsp;


________________________________________________________________________________

## 3. VERSIONING POLICY

This project and all libraries inside the monorepo strictly enforce Semantic Versioning (SemVer) via Git tags (vX.Y.Z) to ensure predictability and prevent breaking downstream applications.
