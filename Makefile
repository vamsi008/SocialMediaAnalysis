
# This is how we want to name the binary output
BINARY=fullmetalgo

MKDIR_P = mkdir -p
COPY_R=cp -R
OUT_DIR="/usr/share/"
FOLD="fullmetal"

# These are the values we want to pass for Version and BuildTime
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`

compose:
		${COPY_R} ${FOLD}  ${OUT_DIR}
		
#builds the binar in the current location.
build:
		
		go build

#installs the binary in the bin location.
install:
		go get
		godep go install 
	  

# Cleans the porject and deletes the binaries.
clean:
		go clean

