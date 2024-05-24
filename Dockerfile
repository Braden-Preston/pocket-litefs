FROM golang:1.22

# Copy source code (w/.dockerignore)
COPY ./ /app

# Install system dependencies
RUN apt-get update -y && \
    apt-get install -y ca-certificates fuse3 sqlite3 curl

# Set working directory
WORKDIR /app

# Build golang packages and utitlies
RUN go mod download
RUN go build -o ./bin/pocketbase ./main.go
RUN chmod -R ug+x ./bin

# Open up app ports to host
EXPOSE 8080

# Install LiteFS
COPY --from=flyio/litefs:0.5 /usr/local/bin/litefs /usr/local/bin/litefs

# Use LifeFS as supervisor and entrypoint
ENTRYPOINT [ "litefs", "mount" ]
