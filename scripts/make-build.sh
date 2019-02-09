DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
ROOT=$DIR/..

go build -o $ROOT/dist/rc $ROOT/src; chmod +x $ROOT/dist