#!/usr/bin/env sh

check_err () {
    err=$?
    if [ $err -eq 1 ]; then
        echo "!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!"
        echo "Error occurred status: $err"
        exit
    fi
}
#get root directory of git repo
GIT_ROOT=$(git rev-parse --show-toplevel)
cd $GIT_ROOT/example
echo $GIT_ROOT
for example_dir in *; do
    if [ -d $GIT_ROOT/example/$example_dir ]; then
    cd $GIT_ROOT/example/$example_dir
    case "$example_dir" in
        "bufferedReadFile")
            echo "Running example/$example_dir option 1"
            go run main.go 1
            check_err
            echo "*************************************************"
            echo "Running example/$example_dir option 2"
            go run main.go 2
            check_err
            ;;
        *) 
            echo "Running example/$example_dir"
            go run main.go
            check_err
            ;;
    esac
    echo "================================================="
    fi
done
