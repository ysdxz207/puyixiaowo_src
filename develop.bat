echo off
SETLOCAL ENABLEDELAYEDEXPANSION
::::::::::::::::::::::::::::::
set SITE_NAME=puyixiaowo
set GRUNT_TASK=lunr-search
set SITE_DIR=..\puyixiaowo
set SITE_IGNORE=.git README.md




::::::::::::::::::::::::::::::
set TEMP_DIR= %temp%\%SITE_NAME%

::delete public dir

echo delete public dir...
if exist "public" (
	rd public /s /q
)

::install dependencies
echo installing dependencies...
call npm install
::compass scss and regenerate search.json
echo compass scss and regenerate search.json...
call grunt %GRUNT_TASK%
::generate static html
echo generate static html...
call hugo
::clean site dir
echo clean site dir...
for /f "eol=: delims=" %%F in ('dir /b /a "%SITE_DIR%" ^| findstr /vib "%SITE_IGNORE%"') do (
	set filename=%SITE_DIR%\%%F
	del "!filename!" /q
	if exist "!filename!" (
		rd "!filename!" /s /q
	)
)
::copy public/ to site dir
echo copy public/ to site dir...
xcopy public\*.* %SITE_DIR%\*.* /E

::commit files
cd %SITE_DIR%
call git pull
call git add .
call git commit -m "auto commit"
call git push

ENDLOCAL

pause