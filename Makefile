# Build Path
BUILD_PATH=./cmd/apiserver

# This how we want to name the binary output
BINARY=./bin/searcher

# These are the values we want to pass for VERSION and BUILD
VERSION=`git describe --abbrev=6 --always --tag`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.version=${VERSION}"
GOFLAGS=-a -tags bg-partner-offer -installsuffix searcher -mod=vendor
bin-linux:
	echo "  >  Building linux-amd64 binary \"searcher\" $(VERSION)..."
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY)-linux $(BUILD_PATH)