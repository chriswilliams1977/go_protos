# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

#Turn on GO modules
ENV GO111MODULE=on

# Copy local code to the container image.
WORKDIR /app

#Copy over module files
COPY go.mod .
#Get generated when using specific package versions
#COPY go.sum .

#Download dependencies in mod file
RUN go mod download

#Copy app files to WORKDIR
COPY ./cmd/ .
ADD ./internal/. ./internal/
ADD ./web/. ./web/

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o snippets-example

# Use a Docker multi-stage build to create a lean production image.
#NOTE THIS WILL REMOVE APP DIR AND ONLY COPY GO COMPILED BINARY OVER
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine

WORKDIR /app
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/snippets-example ./snippets-example
#RUN mkdir ./ui/
#Create web dir and copy static files
COPY --from=builder /app/web ./web/

#set global env var to build var
ENV APP_ENV "production"

#set global env var to build var
ENV UI_PATH "./web/"

# Run the web service on container startup.
CMD ["/app/snippets-example"]