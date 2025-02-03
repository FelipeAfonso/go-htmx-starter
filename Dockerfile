# STEP 1: Build frontend assets
FROM oven/bun:1 AS frontend-builder

WORKDIR /app

# Copy package files
COPY package.json bun.lockb ./

# Install dependencies
RUN bun install --frozen-lockfile
# RUN bun add lightningcss-linux-x64-gnu @tailwindcss/oxide-linux-x64-gnu
# Copy frontend source files
COPY assets/ ./assets/
COPY pages/ ./pages/
COPY static/ ./static/
COPY vite.config.js ./
COPY postcss.config.cjs ./

# Build frontend assets
RUN bun run build

# STEP 2: Build executable binary
FROM golang:alpine AS backend-builder

RUN adduser -D -g '' appuser
RUN mkdir /app

WORKDIR /app

# copy go mod and sum files
COPY go.mod .
COPY go.sum .

RUN go mod download

# copy the source code
COPY . .

# Copy built assets from frontend builder
COPY --from=frontend-builder /app/dist ./dist

# build the binary
RUN go run github.com/a-h/templ/cmd/templ@latest generate
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app .

RUN chmod 0555 /go/bin/app

# STEP 3: Build final image
FROM alpine:latest

WORKDIR /app

# Install ca-certificates
RUN apk --no-cache add ca-certificates

COPY --from=backend-builder /etc/passwd /etc/passwd

# Copy static assets
COPY --from=frontend-builder /app/dist ./dist

USER appuser

# Copy our static executable
COPY --from=backend-builder /go/bin/app ./app_binary

EXPOSE 3000

CMD ["./app_binary"]

