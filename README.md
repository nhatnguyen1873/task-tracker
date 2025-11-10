# Task Tracker

A simple command-line interface (CLI) application to track and manage your tasks. Built with Go to help you practice working with the filesystem, handling user inputs, and building CLI applications.

## Features

- ✅ Add, update, and delete tasks
- ✅ Mark tasks as in progress or done
- ✅ List all tasks or filter by status
- ✅ Store tasks in JSON format
- ✅ Automatic file creation and management
- ✅ Cross-platform support (Linux, macOS, Windows)

## Installation

### Prerequisites

- Go 1.16 or higher installed on your system
- Basic knowledge of command-line operations

### Steps

1. Clone or download this repository:

```bash
git clone <your-repository-url>
cd task-tracker
```

2. Install the application:

```bash
go install
```

This will build the application and install it to your `$GOPATH/bin` directory (usually `~/go/bin`).

3. Make sure `$GOPATH/bin` is in your system PATH:

**Linux/macOS:**

```bash
# Add to ~/.bashrc or ~/.zshrc
export PATH="$HOME/go/bin:$PATH"

# Reload your shell
source ~/.bashrc
```

**Windows:**

- Add `%USERPROFILE%\go\bin` to your PATH environment variable
- Restart your terminal

4. Verify installation:

```bash
task-cli
```

## Usage

### Basic Commands

#### Add a new task

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

#### Update a task

```bash
task-cli update 1 "Buy groceries and cook dinner"
# Output: Task 1 updated successfully
```

#### Delete a task

```bash
task-cli delete 1
# Output: Task 1 deleted successfully
```

#### Mark a task as in progress

```bash
task-cli mark-in-progress 1
# Output: Task 1 marked as in-progress
```

#### Mark a task as done

```bash
task-cli mark-done 1
# Output: Task 1 marked as done
```

#### List all tasks

```bash
task-cli list
```

#### List tasks by status

```bash
# List all tasks that are not done
task-cli list todo

# List all tasks that are in progress
task-cli list in-progress

# List all tasks that are done
task-cli list done
```

## Task Properties

Each task has the following properties:

- **id**: A unique identifier for the task
- **description**: A short description of the task
- **status**: The status of the task (`todo`, `in-progress`, `done`)
- **created_at**: The date and time when the task was created
- **updated_at**: The date and time when the task was last updated

## Data Storage

Tasks are stored in a JSON file with the following behavior:

- **Default location**: `~/.task-tracker/tasks.json` (Linux/macOS) or `C:\Users\YourUsername\.task-tracker\tasks.json` (Windows)
- The file is automatically created if it doesn't exist
- All tasks are persisted between sessions

### Custom Storage Location

You can customize where tasks are stored by setting the `TASK_CLI_DATA_DIR` environment variable:

**Linux/macOS:**

```bash
export TASK_CLI_DATA_DIR="/path/to/your/directory"
```

**Windows (PowerShell):**

```powershell
$env:TASK_CLI_DATA_DIR = "C:\path\to\your\directory"
```

**Windows (CMD):**

```cmd
setx TASK_CLI_DATA_DIR "C:\path\to\your\directory"
```

## Example Workflow

```bash
# Add some tasks
task-cli add "Buy groceries"
task-cli add "Write documentation"
task-cli add "Review pull requests"

# View all tasks
task-cli list

# Start working on a task
task-cli mark-in-progress 1

# Complete a task
task-cli mark-done 1

# Update a task description
task-cli update 2 "Write comprehensive documentation"

# View only pending tasks
task-cli list todo

# View completed tasks
task-cli list done

# Delete a task
task-cli delete 3
```

## Project Structure

```
task-tracker/
├── main.go       # CLI interface and command handlers
├── task.go       # Task struct and business logic
├── storage.go    # File I/O operations
├── utils.go      # Utility functions
├── go.mod        # Go module definition
└── README.md     # This file
```

## Error Handling

The application handles various error scenarios:

- Invalid command usage
- Non-existent task IDs
- File read/write errors
- Invalid status values
- Missing arguments

All errors are displayed with clear messages to help you understand what went wrong.

## Building from Source

If you want to build the application without installing it:

```bash
# Build for your current platform
go build -o task-cli

# Run the built executable
./task-cli list
```

### Cross-compilation

Build for different platforms:

```bash
# For Linux
GOOS=linux GOARCH=amd64 go build -o task-cli-linux

# For macOS
GOOS=darwin GOARCH=amd64 go build -o task-cli-macos

# For Windows
GOOS=windows GOARCH=amd64 go build -o task-cli.exe
```

## Development

To modify and test the application:

1. Make your changes to the source files
2. Rebuild and reinstall:

```bash
go install
```

3. Test your changes:

```bash
task-cli <your-command>
```

## Contributing

Contributions are welcome! Feel free to:

- Report bugs
- Suggest new features
- Submit pull requests
- Improve documentation

## License

This project is open source and available for educational purposes.

## Credits

This project is inspired by the [Task Tracker](https://roadmap.sh/projects/task-tracker) project from [roadmap.sh](https://roadmap.sh) - a community-driven platform for learning paths and project ideas for developers.

## Roadmap

Potential future enhancements:

- [ ] Task priority levels
- [ ] Due dates and reminders
- [ ] Task categories/tags
- [ ] Search functionality
- [ ] Export to different formats (CSV, Markdown)
- [ ] Task statistics and reports
- [ ] Colored output for better readability
- [ ] Task history and undo functionality

## Troubleshooting

### Command not found

If you get a "command not found" error:

1. Verify Go is installed: `go version`
2. Check if `$GOPATH/bin` is in your PATH: `echo $PATH` (Linux/macOS) or `echo %PATH%` (Windows)
3. Verify the binary exists: `ls ~/go/bin/task-cli` (Linux/macOS) or `dir %USERPROFILE%\go\bin\task-cli.exe` (Windows)

### Permission denied (Linux/macOS)

If you encounter permission issues:

```bash
chmod +x ~/go/bin/task-cli
```

### Cannot find tasks.json

The file will be created automatically when you add your first task. If you're looking for it manually, check:

- Default: `~/.task-tracker/tasks.json`
- Custom: Location specified in `TASK_CLI_DATA_DIR`

## Support

If you encounter any issues or have questions, please open an issue in the repository.

---

**Happy task tracking! 🚀**
