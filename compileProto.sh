#!/bin/sh
SRC_DIR="./proto/*"
MAIN_GO="main.go"
PACKAGE_PREFIX="_ "

rm -f $MAIN_GO
exec 3<>$MAIN_GO
echo "package main" >&3
echo "import (" >&3

for f in $SRC_DIR; do
    if [ -d "$f" ]; then
        protoc -I=$f --go_out=plugins=grpc:$f $f/*.proto
        echo $PACKAGE_PREFIX\"$f\" >&3
    fi
done

echo ")" >&3
echo "func main() {}" >&3
# Close fd 3
exec 3>&-

# format go file
go fmt $MAIN_GO
