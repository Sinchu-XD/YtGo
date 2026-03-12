FROM golang:1.21

WORKDIR /app

# Install dependencies
RUN apt-get update && apt-get install -y \
    yt-dlp \
    ffmpeg \
    ca-certificates \
    curl \
    && rm -rf /var/lib/apt/lists/*

# Copy project files
COPY . .

# Run test
CMD ["go", "run", "test.go"]
