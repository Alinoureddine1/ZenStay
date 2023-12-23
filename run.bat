@echo off
setlocal

set SOURCE_DIR=D:\Projects\ZenStay

cd %SOURCE_DIR%

go build -o zenstay.exe ./cmd/web

zenstay.exe

endlocal

