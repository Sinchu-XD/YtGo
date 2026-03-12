package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Video struct {
	Title     string `json:"title"`
	Duration  int    `json:"duration"`
	Channel   string `json:"channel"`
	ViewCount int    `json:"view_count"`
	URL       string `json:"url"`
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

	out, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("❌ yt-dlp failed")
		fmt.Println(string(out))
		return
	}

	var video Video
	err = json.Unmarshal(out, &video)
	if err != nil {
		fmt.Println("JSON parse error:")
		fmt.Println(string(out))
		return
	}

	fmt.Println("=========== YOUTUBE TEST ===========")
	fmt.Println("Title:", video.Title)
	fmt.Println("Channel:", video.Channel)
	fmt.Println("Duration:", video.Duration)
	fmt.Println("Views:", video.ViewCount)
	fmt.Println("Stream URL:", video.URL)
}
