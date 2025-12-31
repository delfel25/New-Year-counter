<p align="center">

**New Year counter - it counts the time on your computer along with the date and tells you how much time it is until New Year's**


![photo of what it looks like](img/mainphoto.png)

---
 ![Go](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go&logoColor=white)  ![Terminal](https://img.shields.io/badge/Platform-Terminal-4D4D4D?style=flat&logo=gnumetiterminal&logoColor=white)


### Requirements
* Go Compiler: Version 1.20 or higher.

* Linux/macOS: Any standard terminal (Bash, Zsh, Fish).
* Windows: It is highly recommended to use Windows Terminal or PowerShell. Older versions of cmd.exe may not support ANSI colors and dynamic resizing properly.

* Permissions: Access to execute stty (on Unix systems) to detect screen size automatically.

* Font: A monospaced font (like JetBrains Mono, Cascadia Code, or Courier) is required for the ASCII digits to line up correctly.

### Install

``` bash
# Clone the repositories
git clone https://github.com/delfel25/New-Year-counter.git

# Go to the project address
cd New-Year-counter

# Build the project
go create main.go

# Run the program
./main
