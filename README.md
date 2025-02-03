# Obsidian Emoji Decorator

![Go](https://img.shields.io/badge/Go-1.23.1-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)
![Version](https://img.shields.io/badge/Version-v1.0.0-blue.svg)

---

## Other Languages

This documentation is also available in:
- [Russian](README.ru.md)

---

## Overview

**Obsidian Emoji Decorator** is a Go-based tool designed to automate the process of adding emojis to folder and file names within your Obsidian vault.  
When you have a large number of notes, it can be challenging to quickly identify which folder is needed at a glance. One effective method is to use emojis as visual markers—but doing this manually is tedious and error-prone. This tool solves that problem by automatically propagating a designated "folder" emoji (e.g. 📂) to all nested folders and files.  
If you wish for an internal folder to display a different emoji, simply add it manually, and the contained Markdown files will inherit that emoji. **Note:** The process is irreversible—if you decide to remove the emojis later, it must be done manually.

---

## Features

- **Automatic Emoji Inheritance:**  
  Adding an emoji (e.g., 📂) at the end of a folder name automatically applies that emoji to all nested folders and Markdown files.
- **Selective Emoji Application:**  
  If an internal folder or file already contains an emoji or a different one, the script will preserve it.
- **Markdown-Specific Processing:**  
  Only files with the `.md` extension are processed, ensuring that other file types remain unaffected.
- **Configurable Skip Patterns:**  
  You can define a list of prefixes (e.g., hidden folders starting with `.` or `_`) to be skipped during processing.
- **Concurrency Control:**  
  The script uses a configurable concurrency limit to process directories in parallel without overloading the system.
- **Error Logging and Reporting:**  
  Built-in logging (using Logrus) provides detailed insights into the renaming process and potential errors.

---

## Project Structure and Visual Diagrams

The project is organized into several packages:

- **`config`** – Loads and parses configuration from a YAML file.
- **`decorator`** – Contains the core logic to recursively process directories and rename files and folders.
- **`emoji`** – Provides utility functions to extract, add, and check for emojis in strings.
- **`main`** – Initializes the application, loads configuration, and starts the decorator process.

### Visual Diagrams

- **Class Diagram:**  
  View the [Class Diagram](docs/ClassDiagram.mmd) for an overview of the main structures and their methods.
  
- **Flowchart Diagram:**  
  For a detailed process flow, see the [Flowchart Diagram](docs/Flowchart.mmd).

*(You can open these Mermaid files in the [Mermaid Live Editor](https://mermaid.live) or include them in your documentation site.)*

---

## Installation

### Prerequisites

- **Go:** Ensure you have Go version 1.23.1 or higher installed.  
  Download and install Go from the [official website](https://golang.org/dl/).

### Getting Started

1. **Clone the Repository**

   ```bash
   git clone https://github.com/ANkulagin/golang_decorator_emoji_sb.git
   cd golang_decorator_emoji_sb
   ```

2. **Install Dependencies**

   Run the following command to download and tidy up dependencies:

   ```bash
   go mod tidy
   ```

3. **Build the Application**

   You can build the binary using:

   ```bash
   go build -o emoji_decorator .
   ```

4. **Run the Application**

   Use the provided Makefile commands or run directly. For example:

   ```bash
   ./emoji_decorator -config=configs/config.yaml
   ```

---

## Configuration

The application uses a YAML configuration file (default: `configs/config.yaml`). Below is an example configuration:

```yaml
src_dir: "/path/to/your/obsidian/vault"
log_level: "info"
concurrency_limit: 12
skip_patterns:
  - "."
  - "_"
  - "any"
```

### Configuration Parameters

- **`src_dir`**  
  The path to the Obsidian vault or the folder containing your Markdown notes.  
  *(Can be an absolute or relative path.)*

- **`log_level`**  
  Logging level for the application. Valid options are:  
  `trace`, `debug`, `info`, `warn`, `error`, `fatal`, `panic`.

- **`concurrency_limit`**  
  Maximum number of concurrent goroutines to use when processing directories.

- **`skip_patterns`**  
  A list of filename prefixes; directories whose names start with any of these strings will be skipped during processing.

---

## Usage

### How It Works

1. **Folder Emoji Propagation:**  
   Rename your folder to include the designated emoji at the end (e.g., rename `Project` to `Project 📂`).  
   All nested folders and Markdown files will inherit this emoji automatically.

2. **Custom Emoji for Subdirectories:**  
   If you wish for a subdirectory to display a different emoji, simply rename it accordingly. The tool will apply the new emoji to the files within that directory.

3. **File Naming Convention:**  
   To ensure correct processing, name Markdown files in the following format:  
   `"ordinal_number file_name.md"`  
   For example, a file named `"1 file.md"` will be processed and renamed to include the inherited emoji (e.g., `"1 🤩 file.md"`).

4. **Irreversible Process:**  
   The operation is irreversible—if you decide later that you do not want the emojis, they must be removed manually.

### Running the Tool

You can run the application via the command line:

```bash
./emoji_decorator -config=configs/config.yaml
```

Or use the provided **Makefile** commands:

- **Build:** `make build`
- **Run:** `make run`
- **Test:** `make test`
- **Coverage:** `make coverage`
- **Help:** `make help`

---

## Contribution and Support

Contributions to the project are welcome!  
If you have suggestions, find bugs, or want to add features, please:

- Fork the repository.
- Create a feature branch.
- Submit a pull request with your changes.

For any issues or questions, you can reach out via [Telegram](https://t.me/ANkulagin03) or open an issue on GitHub.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Summary

**Obsidian Emoji Decorator** simplifies the management of your Obsidian notes by automatically adding visual markers (emojis) to folders and Markdown files. With its flexible configuration and automatic recursive processing, the tool saves time and effort, enabling you to quickly identify the relevant folders and files at a glance.

*Note:* The emoji process is irreversible—ensure you have a backup if you need to revert the changes.

---
