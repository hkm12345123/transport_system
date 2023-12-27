FROM golang:1.19.1-alpine3.16 AS builder

# Set necessary environment variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    RUNENV=docker

# Install gcc and related build tools
RUN apk update && apk add --no-cache gcc libc-dev

# Create necessary directories
WORKDIR /storage/private/zeebe
COPY ./storage/private/zeebe .

WORKDIR /storage/private/images
WORKDIR /storage/public/upload/images
COPY ./storage/public/upload/images .
WORKDIR /db
COPY ./db .
# Copy only go.mod and go.sum to download dependencies when they change
WORKDIR /build
COPY go.mod .
COPY go.sum .

# Download dependencies only if go.mod or go.sum changed
RUN go mod download

# Copy the code into the container
COPY . .

# Move to the specific directory containing the main code
WORKDIR /build/cmd/scem

# Build the Go executable
RUN go build -o golang_scem_binary .

# Move to / (root) directory as the place for resulting binary
WORKDIR /

# Copy binary from build to root folder
RUN cp /build/cmd/scem/golang_scem_binary .

# Expose necessary ports
EXPOSE 5000
EXPOSE 5001

# Define the command to run the application
CMD ["/golang_scem_binary"]

# # Build a small image
# FROM scratch

# COPY --from=builder /dist/main /

# # Command to run
# ENTRYPOINT ["/main"]