package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"bufio"
	"strings"
	"time"
)

func md5Hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func sha1Hash(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func sha256Hash(text string) string {
	h := sha256.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func matchHash(input, target string) bool {
	if md5Hash(input) == target {
		return true
	}
	if sha1Hash(input) == target {
		return true
	}
	if sha256Hash(input) == target {
		return true
	}
	return false
}

func dictAttack(hash, wordlist string) {
	file, err := os.Open(wordlist)
	if err != nil {
		fmt.Println("[!] Wordlist not found:", wordlist)
		return
	}
	defer file.Close()

	fmt.Println("[*] Starting dictionary attack...")
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		count++
		if matchHash(word, hash) {
			fmt.Printf("[+] Hash cracked: %s\n", word)
			return
		}
	}
	fmt.Printf("[-] Hash not found after %d attempts\n", count)
}

func bruteForce(hash string, maxLen int) {
	fmt.Println("[*] Starting brute force attack...")
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	count := 0
	for length := 1; length <= maxLen; length++ {
		var generate func(current string)
		generate = func(current string) {
			if len(current) == length {
				count++
				if matchHash(current, hash) {
					fmt.Printf("[+] Hash cracked: %s\n", current)
					os.Exit(0)
				}
				return
			}
			for _, c := range chars {
				generate(current + string(c))
			}
		}
		generate("")
	}
	fmt.Printf("[-] Hash not found after %d attempts\n", count)
}

func main() {
	hash := flag.String("h", "", "Target hash")
	mode := flag.String("m", "dict", "Attack mode: dict or brute")
	wordlist := flag.String("w", "wordlist.txt", "Wordlist file")
	maxLen := flag.Int("l", 6, "Max length for brute force")
	flag.Parse()

	if *hash == "" {
		fmt.Println("[!] Hash required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("[*] Hash Cracker started: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("[*] Target hash: %s\n", *hash)
	fmt.Println("----------------------------------------")

	switch *mode {
	case "dict":
		dictAttack(*hash, *wordlist)
	case "brute":
		bruteForce(*hash, *maxLen)
	default:
		fmt.Println("[!] Unknown mode:", *mode)
	}

	fmt.Println("----------------------------------------")
	fmt.Println("[*] Done.")
}
