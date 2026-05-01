# hash-cracker

Fast hash cracking tool supporting multiple algorithms.

## Features
- MD5, SHA1, SHA256 support
- Dictionary attack
- Brute force mode
- Rainbow table support
- Multi-threading

## Usage
```bash
go run cracker.go -h <hash> -m <mode>
```

## Examples
```bash
# Dictionary attack
go run cracker.go -h 5f4dcc3b5aa765d61d8327deb882cf99 -m dict -w wordlist.txt

# Brute force
go run cracker.go -h 5f4dcc3b5aa765d61d8327deb882cf99 -m brute -l 8
```

## Requirements
- Go 1.19+
- wordlist (for dictionary mode)

## Disclaimer
For authorized and educational use only.
