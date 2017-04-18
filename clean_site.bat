@echo off
setlocal
set "folder=..\puyixiaowo"
set exclude=.git README.md
for /f "eol=: delims=" %%F in ('dir /b /a "%folder%" ^| findstr /vib "%exclude%"') do (
	del "%folder%\%%F" /q
	if exist "%folder%\%%F" (
		rd "%folder%\%%F" /s /q
	)
)

pause