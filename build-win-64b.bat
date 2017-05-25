@echo off


rem  GOARCH: 386 amd64
set GOARCH=amd64

rem GOOS:windows,linux
set GOOS=windows

set TARGETEXT=.exe

rem 编译结果目标目录
set OUTPUTDIR=E:\grih\publish\servers





set GOHOSTARCH=amd64
set GOHOSTOS=windows


set TARGETDIR=%OUTPUTDIR%\%GOOS%-%GOARCH%

mkdir %TARGETDIR%

for %%i in ("%~dp0\.") do (
  set CurDir=%%~ni
)

set TARGET=bin/%CurDir%-%GOOS%-%GOARCH%%TARGETEXT%

set cnf=%CurDir%.conf.example

echo building %TARGET% ...

go build  -o %TARGET%

cp %TARGET% %TARGETDIR%
cp %cnf% %TARGETDIR%
cp %TARGET% ./

echo "OK"

pause