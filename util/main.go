package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

const urlTemplate string = "https://adventofcode.com/2023/day/%d/input"
const targetDir string = "data"
const sessionVar string = "AC_SESSION_ID"

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func GetDayInput(day int) (*os.File, error) {

	targetPath, err := resolveTargetPath(day)
	if err != nil {
		return nil, err
	}
	if checkFileExists(targetPath) {
		return os.Open(targetPath)
	}

	url, err := url.Parse(fmt.Sprintf(urlTemplate, day))
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}

	sessionKey := os.Getenv(sessionVar)
	if sessionKey == "" {
		return nil, fmt.Errorf("%s is unset", sessionVar)
	}

	response, err := downloadInputData(url, sessionKey)
	if err != nil {
		return nil, fmt.Errorf("error downloading data: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("expected 200 OK, got: %s", response.Status)
	}

	fp, err := os.Create(targetPath)
	if err != nil {
		return nil, fmt.Errorf("error opening %s for writing: %w", targetPath, err)
	}

	_, err = io.Copy(bufio.NewWriter(fp), response.Body)
	fp.Close()
	if err != nil {
		return nil, fmt.Errorf("error writing response data to %s: %w", targetPath, err)
	}

	return os.Open(targetPath)
}

func resolveTargetPath(day int) (string, error) {
	targetDirAbs, err := filepath.Abs(targetDir)
	if err != nil {
		return "", fmt.Errorf("error resolving target dir: %w", err)
	}
	targetPath := filepath.Join(targetDirAbs, fmt.Sprintf("day%d.txt", day))
	return targetPath, nil
}

func downloadInputData(url *url.URL, sessionKey string) (*http.Response, error) {
	var err error

	jar, err := makeSessionIdCookieJar(url, sessionKey)
	if err != nil {
		return nil, fmt.Errorf("error making cookie jar: %w", err)
	}

	client := &http.Client{Jar: jar}
	response, err := client.Get(url.String())
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}
	return response, nil
}

func makeSessionIdCookieJar(url *url.URL, sessionKey string) (*cookiejar.Jar, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	jar.SetCookies(url, []*http.Cookie{{Name: "session", Value: sessionKey}})
	return jar, nil
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}
