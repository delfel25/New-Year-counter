<p align="center">
  <img src="img/mainphoto.png" alt="New Year Counter Preview" width="600">
</p>

<h1 align="center">New Year Counter</h1>

<p align="center">
  <strong>An adaptive terminal countdown timer with dynamic scaling and midnight fireworks!</strong>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Language-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go">
  <img src="https://img.shields.io/badge/Platform-Terminal-4D4D4D?style=for-the-badge&logo=gnumetiterminal&logoColor=white" alt="Terminal">
  <img src="https://img.shields.io/github/license/delfel25/New-Year-counter?style=for-the-badge&color=important" alt="License">
</p>

---

### 📖 Description
This tool syncs with your system clock to provide a real-time countdown to the New Year. It features a custom-built matrix rendering engine that adjusts the size of the digits based on your terminal window dimensions.

---

### Requirements

* **Go Compiler**: Version `1.20` or higher.
* **Environment**: 
    * **Linux/macOS**: Any standard terminal (Bash, Zsh, Fish).
    * **Windows**: Highly recommended to use **Windows Terminal** or **PowerShell**.
* **System**: Access to `stty` (Unix) for auto-scaling support.
* **Font**: A **monospaced font** (e.g., JetBrains Mono, Courier) is essential for correct alignment.

---

### Installation

Follow these steps to get the counter running:

```bash
# 1. Clone the repository
git clone [https://github.com/delfel25/New-Year-counter.git](https://github.com/delfel25/New-Year-counter.git)

# 2. Enter the directory
cd New-Year-counter

# 3. Build the executable
go build main.go

# 4. Run the program
./main
