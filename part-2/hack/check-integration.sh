#!/bin/sh
set -o errexit
set -o nounset

rm -rf ./tests/debug/*

check_all_tests_on_dir(){
    echo "Running tests on dir: "$dir
    for test_file in $dir/*.py; do
        echo "Running test: "$test_file
        python "$test_file";
    done
}

single_test(){
    cmd="python ./tests/integration/$@"
    echo "Running: $cmd"
    $cmd
}

if [ $# -ne 0 ]; then
    single_test $@
    exit
fi

dir=./tests/integration
check_all_tests_on_dir
