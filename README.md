# GopherQuill
A Powerful Code Review Tool for Developers and Engineering Teams

GopherQuill is an open-source tool designed to streamline code reviews by analyzing staged files across various programming languages. It leverages GPT-powered insights to deliver high-quality, actionable feedback, helping developers maintain consistent code standards and engineering managers improve their team's productivity.

## Features
- Multi-Language Support: Analyzes files written in Python, Go, JavaScript, React, and more.
- Staged Files Focus: Only reviews files staged for commit, ensuring relevance.
- Customizable Rules: Supports dynamic configuration through .gopherquillrc.
- Pre-Commit Integration: Blocks commits with actionable feedback for unresolved issues.
- Developer and Manager-Friendly:
  - CLI: Lightweight and fast for developers.
   - UI: Integrates with GitHub for engineering managers and code reviewers.
## Table of Contents
1. Installation
2. Usage
    - CLI
    - Pre-Commit Hook
3. Configuration
4. Supported Languages
5. Contribution
6. License
   
## Installation
### Install via Homebrew

```
brew tap Princesso/gopherquill
brew install gopherquill
```

### Manual Installation
1. Clone the repository:

    ```
    git clone https://github.com/Princesso/gopherquill.git
    cd gopherquill 
    ```

2. Build the binary:
    ```
    go build -o gopherquill main.go
    ```
3. Add the binary to your PATH:
    ```
    mv gopherquill /usr/local/bin/
    ```

## Usage
### CLI
#### Analyze a Git Directory
```
    gopherquill -analyze -dir /path/to/your/git/project
```

#### Analyze Staged Files
```
gopherquill -staged
```

### Pre-Commit Hook
1. Add GopherQuill as a pre-commit hook:

    ```
    cp scripts/pre-commit .git/hooks/pre-commit
    chmod +x .git/hooks/pre-commit
    ```

2. Test by committing staged files:

    ```
    git add .
    git commit -m "Test pre-commit"
    ```

If there are issues, the commit will be blocked, and feedback will be provided.

## Configuration
Create a .gopherquillrc file in the root of your project to customize settings:

```
{
  "languages": ["python", "javascript", "go", "react"],
  "max_feedback": 10,
  "gpt_model": "gpt-4",
  "feedback_format": "detailed",
  "ignore_files": ["README.md", "docs/"]
}

```

## Supported Languages

GopherQuill can analyze codebases written in:

- Python
- Go
- JavaScript
- React
- Java
- Other Languages (Generic Code Reviews)
  
## Contribution
Contributions are welcome! Follow these steps to contribute:

1. Fork the repository and clone your fork.
2. Create a new branch:
    ```
    git checkout -b feature/new-feature
    ```
3. Make your changes and commit:
    ```
    git commit -m "Add a new feature"
    ```
4. Push to your branch:
    ```
    git push origin feature/new-feature
    ```
5. Open a pull request on GitHub.

## License
GopherQuill is licensed under the MIT License.

Roadmap
UI Integration: Develop a user-friendly web interface for GitHub and other VCS platforms.
Advanced Analysis: Add more programming languages and frameworks.
Team Analytics: Provide team-level insights and historical code review data.
For detailed documentation and updates, visit GopherQuill GitHub Repository.

