# This script build application code(main.go) implemented by Go and generate executable file.

if !(go version); then
    echo "this script require Go1.17 programming language for building application."
    exit 1
fi


buildCmd="go build"
buildArgs="-ldflags='-s -w' -tags osusergo -o bin/run"

# generate executable file for Linux
GOOS="linux"
GOARCH="amd64"
build="$buildCmd $buildArgs-$GOOS-$GOARCH ."
eval $build

# generate executable file for Mac
GOOS="darwin"
build="$buildCmd $buildArgs-$GOOS-$GOARCH ."
eval $build

# generate executable file for Windows
GOOS="windows"
build="$buildCmd $buildArgs-$GOOS-$GOARCH ."
eval $build

