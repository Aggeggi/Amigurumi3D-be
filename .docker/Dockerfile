FROM golang:1.22-alpine AS builder

WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./amigurumi-be .

FROM golang:1.22-alpine AS release

WORKDIR /app

# Copy our static executable.
COPY --from=builder /app/amigurumi-be .

ENV PORT 8000
ENV GIN_MODE release
EXPOSE 8000

# Run the executable
CMD ["./amigurumi-be"]