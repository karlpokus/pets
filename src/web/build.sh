#!/bin/bash

npm run test --silent
TEST_RESULT=`echo $?`
if test $TEST_RESULT -ne 0; then
  echo "tests failed. Exiting"
  exit $TEST_RESULT
fi

VERSION="v`cat package.json | jq -r .version`"

docker build -t pokus2000/pets-web:$VERSION .
