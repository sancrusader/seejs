package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// Fungsi untuk mengambil dan menganalisis konten dari URL
func analyzeURL(url string, output *os.File) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(output, "Gagal mengunduh URL: %s\n", url)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(output, "Gagal membuka URL: %s, Status Code: %d\n", url, resp.StatusCode)
		return
	}

	// Regex untuk mendeteksi API keys yang lebih kompleks
	apiKeyRegex := regexp.MustCompile(`(?i)(?:(?:AIza[0-9A-Za-z-_]{35})|(?:[0-9A-Za-z]{32}|[0-9A-Za-z]{39})|(?:AKIA[0-9A-Z]{16})|(?:aws_access_key_id\s*[:=]\s*['"]?([A-Z0-9]{20})['"]?)|(?:[A-Za-z0-9/+=]{40})|(?:aws_secret_access_key\s*[:=]\s*['"]?([A-Za-z0-9/+=]{40})['"]?)|(?:[0-9a-f]{32})|(?:[A-Za-z0-9]{24}-[0-9a-zA-Z]{6}-[0-9a-f]{36})|(?:sk_(test|live)_[0-9a-zA-Z]{24})|(?:whsec_[0-9a-zA-Z]{24})|(?:AC[0-9a-fA-F]{32})|(?:SK[0-9a-zA-Z]{32})|(?:ghp_[0-9A-Za-z]{36})|(?:gho_[0-9A-Za-z]{36})|(?:github_pat_[0-9]{12}\.[0-9a-zA-Z_-]{7}\.[0-9a-zA-Z_-]{27})|(?:AIza[0-9A-Za-z-_]{35})|(?:[0-9]{1,}:AA[a-zA-Z0-9_-]{33})|(?:xoxb-[0-9]{12}-[0-9]{12}-[0-9A-Za-z]{24})|(?:xoxp-[0-9]{12}-[0-9]{12}-[0-9]{12}-[0-9A-Za-z]{24})|(?:xoxs-[0-9]{12}-[0-9]{12}-[0-9A-Za-z]{24})|(?:[0-9a-zA-Z]{8,}:[0-9a-zA-Z]{40})|(?:EAAB[0-9A-Za-z]+)|(?:[A-Za-z0-9]{16}\|[A-Za-z0-9]{32})|(?:key-[0-9a-zA-Z]{32})|(?:SG\.[A-Za-z0-9_-]{22}\.[A-Za-z0-9_-]{43})|(?:[0-9a-zA-Z]{20}-us[0-9]{1,3})|(?:[A-Za-z0-9]{32})|(?:[0-9a-f]{32})|(?:[0-9a-zA-Z]{40}))`)

	// Regex untuk mendeteksi URL
	urlRegex := regexp.MustCompile(`(?:"|')(?:(?:(?:[a-zA-Z]{1,10}://|//)[^"'/]{1,}\.[a-zA-Z]{2,}[^"']{0,})|(?:(?:/|\.\./|\./)[^"'><,;| *()%(\\$^/\\\[\]]+[^"'><,;|()]{1,})|(?:[a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/.]{1,}\.(?:[a-zA-Z]{1,4}|action)(?:[\?|#][^"|']{0,}|))|(?:[a-zA-Z0-9_\-/]{1,}/[a-zA-Z0-9_\-/]{3,}(?:[\?|#][^"|']{0,}|))|(?:[a-zA-Z0-9_\-]{1,}\.(?:php|asp|aspx|jsp|json|action|html|js|txt|xml)(?:[\?|#][^"|']{0,}|)))|(?:"|')`)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()

		// Temukan semua API keys
		apiKeys := apiKeyRegex.FindAllString(line, -1)
		for _, key := range apiKeys {
			key = strings.TrimSpace(key)
			if key != "" {
				fmt.Fprintln(output, key)
			}
		}

		// Temukan semua URL
		urls := urlRegex.FindAllString(line, -1)
		for _, u := range urls {
			u = strings.Trim(u, "\"'") // Menghapus delimiter
			if u != "" {
				analyzeURL(u, output) // Menganalisis URL lain yang ditemukan
			}
		}
	}
}

// Fungsi untuk menganalisis URL dari file
func analyzeURLsFromFile(filePath string, output *os.File) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Gagal membuka file: %s\n", filePath)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		analyzeURL(url, output)
	}
}

func main() {
	fileList := flag.String("l", "", "Path ke file yang berisi daftar URL")
	flag.Parse()

	if *fileList == "" {
		fmt.Println("Tolong masukkan path ke file URL dengan flag -l")
		return
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Gagal membuat file output.txt")
		return
	}
	defer outputFile.Close()

	analyzeURLsFromFile(*fileList, outputFile)
}
