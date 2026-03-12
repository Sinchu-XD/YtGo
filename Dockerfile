FROM golang:1.21

WORKDIR /app

RUN apt update && apt install -y \
    python3 \
    python3-pip \
    ffmpeg \
    curl

COPY . .

RUN chmod +x start.sh

CMD ["bash", "start.sh"]
