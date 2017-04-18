set port=1313
start hugo server --port=%port%
start grunt default
start http://localhost:%port%/puyixiaowo