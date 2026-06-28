GitHub CI/CD & Automation Infrastructure
================================================================

&nbsp;

> This directory contains the immutable, data-driven automation, Continuous Integration (CI), Continuous Deployment (CD), and release management infrastructure for this repository.


&nbsp;
&nbsp;


________________________________________________________________________________

## Directory Structure

```text
.github/
  ├── scripts/
  │   └── trigger-release.sh  # Local utility script to trigger releases via terminal
  │
  ├── workflows/
  │     └── ci-cd.yml         # Static unified GitHub Actions pipeline (Test, Security & Deploy)
  │
  ├── ACTIONS.md              # This technical documentation guide
  ├── entrypoints.txt         # Data file containing paths to main.go targets (project-specific)
  └── goreleaser.yaml         # Static cross-platform compilation configuration (GoReleaser)
```


&nbsp;
&nbsp;


________________________________________________________________________________

## Pipeline Architecture & Lifecycle

The `workflows/ci-cd.yml` file orchestrates the complete software development lifecycle automatically through two sequential stages (Jobs):


&nbsp;


### 1. Test Suite & Security Job (`test`)

*   **Trigger:** Executes on any `push` or `pull_request` targeting the `main` branch.
*   **Core Responsibilities:**
    *   Provisions the Go environment utilizing the configured stable version.
    *   **Vulnerability Scanning (`govulncheck`):** Statically analyzes the codebase to detect known security vulnerabilities within upstream third-party dependencies.
    *   **Multi-Platform Matrix Testing:** Runs the full unit test suite (`go test -v -cover`) concurrently across three distinct operating systems: **Linux (Ubuntu), Windows, and macOS**.


&nbsp;


### 2. Release & Artifact Generation Job (`release`)

*   **Trigger:** Executes exclusively on `push` events to the `main` branch, **if and only if** all checks in the preceding `test` job complete with a 100% success rate.
*   **Core Responsibilities:**
    *   Calculates the next Semantic Versioning (SemVer) target dynamically based on commit history.
    *   Performs an integrity verification compilation (`go build`) across all local modules.
    *   Generates and publishes the new immutable version tag (`vX.Y.Z`) to the upstream repository.
    *   **Conditional Artifact Compilation:** Inspects `.github/entrypoints.txt`. If valid active paths are found (ignoring comments and empty lines), it triggers GoReleaser using the static configuration file (`-f .github/goreleaser.yaml`) to cross-compile binary executables and publish a formal **GitHub Release** populated with `.tar.gz` and `.zip` distribution artifacts. If the file is empty or contains only comments, it safely skips binary compilation, treating the project purely as a library.


&nbsp;
&nbsp;


________________________________________________________________________________

## Semantic Versioning Strategy

The automated versioning system parses the message payload of the **latest commit** to compute the next release iteration:

*   **Repository Bootstrapping:** If no prior Git tags are registered in the repository, the pipeline initializes the version baseline automatically at `v0.0.1`.
*   **Feature Increment (`Minor`):** Commits prefixed with `feat:` (e.g., `feat: add logging system`) increment the minor version component (`v0.1.0`).
*   **Patch / Maintenance Increment (`Patch`):** Standard commits or those prefixed with `fix:`, `chore:`, or `docs:` increment the lower patch component (`v0.0.2`).
*   **Breaking Changes (`Major`):** Any commit containing the phrase `BREAKING CHANGE` within its body or footer increments the major version component (`v1.0.0`).
*   **Manual Override Configuration:** Commits prefixed explicitly with `release: vX.Y.Z` bypass the semantic heuristic engine and force the application of the designated version string.


&nbsp;
&nbsp;


________________________________________________________________________________

## Execution: How to Trigger a Release

To prevent tag pollution and overhead during continuous development, you can push multiple iterative commits to `main` without creating a version. Once you are satisfied with the state of the codebase, trigger a formal release using one of the following methods:


&nbsp;


### Option A: Local Terminal Script (Automated Empty Commit)

Execute the bundled shell script to generate an infrastructure-only empty commit that signals the deployment engine:

1.  Grant executable permissions (first-time setup only):
    ```bash
    chmod +x .github/scripts/trigger-release.sh
    ```
2.  To trigger a default automated semantic patch increment (`+1 patch`):
    ```bash
    .github/scripts/trigger-release.sh
    ```
3.  To enforce a strict manual target version:
    ```bash
    .github/scripts/trigger-release.sh v1.0.0
    ```


&nbsp;


### Option B: GitHub UI Manual Dispatch (`workflow_dispatch`)

1.  Navigate to the **Actions** tab of your repository on GitHub.
2.  Under the left-hand workflows sidebar, select **CI/CD Pipeline**.
3.  Click the grey **Run workflow** dropdown component located on the right side of the interface.
4.  Target the `main` branch and click the green **Run workflow** button to initiate deployment without generating additional commit logs.
