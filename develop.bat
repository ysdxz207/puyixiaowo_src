echo off
SETLOCAL ENABLEDELAYEDEXPANSION
::::::::::::::::::::::::::::::
set SITE_NAME=puyixiaowo
set GRUNT_TASK=lunr-search
set SITE_DIR=..\puyixiaowo
set SITE_IGNORE=(.git README.md)




::::::::::::::::::::::::::::::
set TEMP_DIR= %temp%\%SITE_NAME%

::delete public dir
rd public /s /q
::compass scss and regenerate search.json
echo grunt %GRUNT_TASK%
::generate static html
echo hugo
::clean site dir
rd %TEMP_DIR% /s /q

for %%a in %SITE_IGNORE% do (
	xcopy %SITE_DIR%\%%a\*.* %TEMP_DIR%\%%a\*.* /E /K
)
ENDLOCAL

pause