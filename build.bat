@echo off

REM Set the custom binary name
SET BINARY_NAME=pp.exe

REM Build the Go source code with optimizations
go build -o "%BINARY_NAME%" -ldflags "-s -w" -gcflags "all=-N -l" .\src\main.go

REM Check if the build was successful
IF %ERRORLEVEL% EQU 0 (
    echo Build successful! The binary '%BINARY_NAME%' is ready.
) ELSE (
    echo Build failed.
)