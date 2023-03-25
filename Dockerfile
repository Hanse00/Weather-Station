FROM golang:1.20 AS base
WORKDIR /app

# Download any required modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code
COPY *.go .

FROM base as test
# When test is executed, run the tests
CMD [ "go", "test", "-v", "." ]

FROM base as build
# Build the executable
WORKDIR /app
RUN CGO_ENABLED=0 go build -o weather-station

FROM gcr.io/distroless/static-debian11 as prod
COPY --from=build /app/weather-station /weather-station

# Set the default port for the container to use
# This can be overridden at runtime
ENV HTTP_PORT=8080

# Run the executable
ENTRYPOINT [ "/weather-station" ]
