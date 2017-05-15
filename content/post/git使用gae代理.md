+++
date = "2017-05-08T15:38:22+08:00"
title = "git使用gae代理"
tags = ["git","gae"]
categories = ["版本管理"]
+++

用户目录下**.gitconfig**配置如下:

```
[user]
	email = xxx@qq.com
	name = xxx
[http]
	proxy = http://127.0.0.1:8087
	sslVerify = false
[https]
	proxy = http://127.0.0.1:8087
	sslVerify = false

```