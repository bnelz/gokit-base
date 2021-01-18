#!/usr/bin/env bash

# The name of your application
APP_NAME="gokit-base"

# The tag for your container
TAG_NAME="catpipegrep/gokit-base"

# Your project's module root
MODULE_ROOT="github.com/bnelz"

# The base directory for this project
BASE_DIR=$(cd "$(dirname "$0")/../" && pwd)

# The built binary location
PROJECT_BINARY="docker/${APP_NAME}/bin/${APP_NAME}"

# A customizable buildpath
BUILD_PATH="/usr/local/go/src/${MODULE_ROOT}/${APP_NAME}"

# The official golang container link and version for our build container
BASE_GOLANG_CONTAINER="golang:1.15"

function HELP {
  echo -e "Options"
  echo -e "   -c  --compiler-flags   Pass in compiler flags"
  echo -e "   -v  --version          Specify project version"
  echo -e "   -h, --help             Show this help (-h works with no other options)"\\n
  echo -e "Example:"
  echo -e "   `basename ${BASH_SOURCE[0]}` --install-deps"\\n
  exit 1
}

for arg in "$@"; do
  shift
  case "$arg" in
    "--compiler-flags") set -- "$@" "-c" ;;
    "--version")        set -- "$@" "-v" ;;
    "--help")           set -- "$@" "-h" ;;
    *)                  set -- "$@" "$arg"
  esac
done

# Get arguments
while getopts "ic:v:h" opt; do
	case $opt in
	i)
		INSTALL_DEPS=TRUE
		;;
	c)
	    COMPILER_FLAGS=$OPTARG
	    ;;
	v)
	    VERSION=$OPTARG
	    ;;
	h)
		HELP
		;;
	\?)
		echo "Invalid options: -$OPTARG"
		HELP
		;;
	esac
done

# Compile the application
echo "Compiling ${APP_NAME}..."
docker run -it --rm \
    -v $BASE_DIR:$BUILD_PATH \
    -v $HOME/.ssh:/root/.ssh \
    -w $BUILD_PATH \
    $BASE_GOLANG_CONTAINER \
    sh -c "go build -v -o $PROJECT_BINARY $COMPILER_FLAGS"

# Build and tag our container
echo "Building ${APP_NAME} container"
docker build --no-cache --pull=true $BASE_DIR/docker/${APP_NAME} -t "${TAG_NAME}:${VERSION}"
