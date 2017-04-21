echo off
SETLOCAL ENABLEDELAYEDEXPANSION
::::::::::::::::::::::::::::::
set SITE_NAME=puyixiaowo
set GRUNT_TASK=lunr-search
set SITE_DIR=D:\workspace\hugo\puyixiaowo
set SRC_DIR=D:\workspace\hugo\puyixiaowo_src
set SITE_IGNORE=.git README.md CNAME robots.txt




::::::::::::::::::::::::::::::
set TEMP_DIR= %temp%\%SITE_NAME%

::delete public dir
cd %SRC_DIR%
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
::commit empty
echo commit empty...
call:commit
::copy public/ to site dir
cd %SRC_DIR%
echo copy public/ to site dir...
xcopy public\*.* %SITE_DIR%\*.* /E /Y

::commit files
call:commit
call:push
::echo.&pause&goto:eof
exit


::commit function
:commit
cd %SITE_DIR%
call git pull
call git add .
call git commit -m "auto commit"
goto:eof

::push function
:push
cd %SITE_DIR%
call git push
goto:eof

ENDLOCAL
