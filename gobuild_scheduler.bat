@echo off 

echo Starting Go Build Scheduler...
echo Beginning to clone repository from origin...

set repo_name=poc-github-actions
set repo_link=https://github.com/andre-ols/poc-github-actions-golang

git clone %repo_link%

cd %repo_name%

set GOOS=windows

go build 

%repo_name%.exe