#! /bin/bash

echo "### Creating test folders ###"
mkdir testFolder
cd testFolder
mkdir src
echo "this is test" > src/index.js
echo "FROM ubuntu:20.04" > Dockerfile
echo "this is eslint" > .eslintrc`