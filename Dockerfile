# Build $docker build -t {image}:{tag} .
# tag $docker tag {image}:{tag} {gcr full path}:{tag}
# push $gcloud docker -- push {gcr full path}:{tag}
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.7

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/baveltman/web_client

#build the application
RUN go get github.com/tools/godep
RUN cd /go/src/github.com/baveltman/web_client/main && godep go build

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/src/github.com/baveltman/web_client/main/main

# Document that the service listens on port 8080.
EXPOSE 8080
