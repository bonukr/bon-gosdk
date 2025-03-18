@echo off
setlocal
cd /d "%~dp0"
cls

:: push tags
git push --tags

endlocal
timeout 3