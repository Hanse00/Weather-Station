FROM golang:1.20

WORKDIR /app

# Download any required modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY *.go .

# Run the tests
RUN go test -v

# Build the executable
RUN go build -o weather-station

# Run the executable
ENTRYPOINT [ "/app/weather-station" ]
