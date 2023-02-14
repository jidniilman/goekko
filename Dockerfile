# Stage 1 Create the Builder image
FROM golang:1.20.0-buster AS builder

# Set Go Module ON
ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Jidni Ilman <jidni007@gmail.com>" developer="Jidni Ilman <jidni007@gmail.com>"

# Set Working Directory
WORKDIR /build

# Copy Go Module Files
COPY go.mod .

# Download Dependencies
RUN go mod download

# Copy Go Files To Docker From Repository
COPY . .

# Remove unused file
# RUN rm xyz.go

# Download Neccesary Tools
#RUN GO111MODULE=off go get -u github.com/jteeuwen/go-bindata/...
#RUN GO111MODULE=off go get -u github.com/mgechev/revive
#RUN GO111MODULE=off go get -u github.com/onsi/ginkgo/ginkgo

# Tidy up the code
RUN go mod tidy

# Generate Schema
#RUN cd pkg/app; go generate

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goekko -v cmd/goekko/main.go

# STAGE 2 Create the minimal runtime image
FROM alpine:3.16.1

COPY --from=builder /build/goekko /build/.en[v] ./

# Expose port 1330 to the outside world
EXPOSE 1330

# Command to run the executable
CMD ["./goekko"]