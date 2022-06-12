FROM golang:1.18-alpine

WORKDIR /app

# Install dependencies
COPY go.* .
RUN go mod download

# Copy sources
COPY *.go ./

# Compile
RUN go build -o /terminal-ui

ENTRYPOINT [ "/terminal-ui" ]
