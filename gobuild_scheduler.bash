#!bin/bash

echo Starting Go Build Scheduler...
echo Beginning to clone repository from origin...

repo_name=poc-github-actions
repo_link=https://github.com/andre-ols/poc-github-actions-golang

git clone $repo_link

cd $repo_name

GOOS=linux

go build 

./$repo_name