package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Fungsi untuk membaca konten file JavaScript dan mencari pola-pola tertentu
func analyzeFile(filename string, output *os.File) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(output, "Gagal membuka file: %s\n", filename)
		return
	}
	defer file.Close()

	// Regex untuk mencari pola API key, secret, dan path
	patterns := map[string]*regexp.Regexp{
		"API Key":     regexp.MustCompile(`(?i)(api_key|apikey|key|token|access_token|auth_token)\s*[:=]\s*['"]?([a-zA-Z0-9-_]+)['"]?`),
		"Secret Key":  regexp.MustCompile(`(?i)(secret|secret_key|client_secret)\s*[:=]\s*['"]?([a-zA-Z0-9-_]+)['"]?`),
		"URL Path":    regexp.MustCompile(`(?i)(https?://[^\s'"]+|/[^\s'"]+)`),
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		for name, pattern := range patterns {
			matches := pattern.FindAllString(line, -1)
			if matches != nil {
				fmt.Fprintf(output, "[%s] ditemukan di file %s pada baris %d: %v\n", name, filename, lineNumber, matches)
			}
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(output, "Gagal membaca file: %s\n", filename)
	}
}

// Fungsi untuk menangani list file atau single file
func analyzeFiles(filenames []string, output *os.File) {
	for _, filename := range filenames {
		fmt.Fprintf(output, "Analisis file: %s\n", filename)
		analyzeFile(filename, output)
	}
}

func main() {
	// Flag untuk file list
	fileList := flag.String("l", "", "Path ke file yang berisi daftar file JavaScript")
	flag.Parse()

	if *fileList == "" {
		fmt.Println("Masukkan path ke file daftar menggunakan parameter -l")
		return
	}

	// Membuka file output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Gagal membuat file output.txt")
		return
	}
	defer outputFile.Close()

	// Membaca file list
	listFile, err := os.Open(*fileList)
	if err != nil {
		fmt.Printf("Gagal membuka file daftar: %s\n", *fileList)
		return
	}
	defer listFile.Close()

	// Membaca setiap baris dari file list
	var filenames []string
	scanner := bufio.NewScanner(listFile)
	for scanner.Scan() {
		filename := strings.TrimSpace(scanner.Text())
		if filename != "" {
			filenames = append(filenames, filename)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Gagal membaca file daftar")
		return
	}

	// Menjalankan analisis file
	analyzeFiles(filenames, outputFile)

	fmt.Println("Analisis selesai. Hasil disimpan di output.txt")
}
