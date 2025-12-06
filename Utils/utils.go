package Utils

import (
	"AOC2025/Config"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const year = 2025
const adventDomain = "https://adventofcode.com"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func ReadInput(day int) []string {
	sessionToken := Config.GetConfiguration().AdventSession
	// Construct the URL
	url := fmt.Sprintf("%s/%d/day/%d/input", adventDomain, year, day)
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	// Add the session cookie to the request
	req.Header.Set("cookie", fmt.Sprintf("%s", sessionToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Read each line from the response body
	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text() // Read one line at a time
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: Received status code %d\n", resp.StatusCode)
		return nil
	}

	fmt.Println("Request successful!")
	err = WriteToFile(fmt.Sprintf("Input/Day%v.txt", day), lines)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Successfully wrote to file.")
	}
	return lines
}

func WriteToFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func SubmitString(day, level int, ans string) {
	sessionToken := Config.GetConfiguration().AdventSession
	urlAdvent := fmt.Sprintf("%s/%d/day/%d/answer", adventDomain, year, day)

	formData := url.Values{}
	formData.Set("level", strconv.Itoa(level))
	formData.Set("answer", ans)

	req, err := http.NewRequest("POST", urlAdvent, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Cookie", sessionToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		scanner := bufio.NewScanner(resp.Body)
		messageRegex := regexp.MustCompile(`<article><p>(.*?)</p></article>`)
		linkRegex := regexp.MustCompile(`<a [^>]*>(.*?)</a>`)

		for scanner.Scan() {
			line := scanner.Text()
			matches := messageRegex.FindStringSubmatch(line)
			if len(matches) > 1 {
				message := matches[1]
				cleanMessage := linkRegex.ReplaceAllString(message, "$1")
				if strings.Contains(cleanMessage, "not the right answer") ||
					strings.Contains(cleanMessage, "trying again") ||
					strings.Contains(cleanMessage, "don't seem") {
					fmt.Println(Cyan + cleanMessage + Reset)
				} else {
					fmt.Println(Green + cleanMessage + Reset)
				}
				break
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading response body:", err)
		}
	} else {
		fmt.Printf("Request failed with status: %s\n", resp.Status)
	}
}

func Submit(day, level, ans int) {
	SubmitString(day, level, strconv.Itoa(ans))
}

func Abs(i, j int) int {
	if i-j < 0 {
		return j - i
	}
	return i - j
}

func NumDigits(num int) int {
	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}

func Pow10(exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= 10
	}
	return result
}
