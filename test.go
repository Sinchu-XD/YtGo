package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Entry struct {
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
	Channel   string `json:"channel"`
	ViewCount int    `json:"view_count"`
	URL       string `json:"url"`
}

type Result struct {
	Entries []Entry `json:"entries"`
}

func main() {

	query := "ytsearch1:dil song"

	cmd := exec.Command(
		"yt-dlp",
		"-j",
		"--no-playlist",
		"-f",
		"bestaudio",
		query,
	)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("yt-dlp error:", err)
		return
	}

	var result Result
	err = json.Unmarshal(out, &result)
	if err != nil {
		fmt.Println("JSON parse error:", err)
		return
	}

	if len(result.Entries) == 0 {
		fmt.Println("No results found")
		return
	}

	video := result.Entries[0]

	fmt.Println("========== YOUTUBE TEST ==========")
	fmt.Println("Title:", video.Title)
	fmt.Println("Channel:", video.Channel)
	fmt.Println("Duration:", video.Duration, "seconds")
	fmt.Println("Views:", video.ViewCount)
	fmt.Println("Stream URL:", video.URL)
}
