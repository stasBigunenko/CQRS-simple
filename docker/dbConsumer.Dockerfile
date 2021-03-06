FROM golang:1.17

WORKDIR /dbConsumer

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

ENTRYPOINT ["go", "run", "./cmd/dbConsumer/dbConsumerMain.go"]