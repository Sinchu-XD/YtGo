#!/bin/bash

echo "Installing yt-dlp..."
pip install -U yt-dlp

echo "Running Go test..."
go run test.go
