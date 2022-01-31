package main

import (
	"io"
	"net/http"
	"os"
)

func downloadPuzzleInput(filename string) (err error) {
	if _, err = os.Stat(filename); !os.IsNotExist(err) {
		// already downloaded
		return
	}
	url := "https://codingmystery.com/assets/puzzle-input/the-beginning/" + filename
	// https://mholt.github.io/curl-to-go
	resp, err := http.Get(url)
	if err != nil {
		return (err)
	}
	defer resp.Body.Close()

	os.Mkdir("inputs", os.ModeDir)
	out, err := os.Create("inputs/" + filename)
	if err != nil {
		return (err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return (err)
	}
	return
}

func readPuzzleInput(filename string) (data string, err error) {
	os.Mkdir("inputs", os.ModeDir)
	var buf []byte
	file, err := os.Open("inputs/" + filename)
	if err != nil {
		return "", err
	}
	buf, err = io.ReadAll(file)
	data = string(buf)
	return
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	println("https://codingmystery.com/the-beginning")
	// shredRecovery()
	// navigateTunnels()
}
