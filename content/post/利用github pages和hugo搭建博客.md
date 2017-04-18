+++
date = "2017-04-15T11:55:48+08:00"
title = "利用github pages和hugo搭建个人博客"
tags = [ "hugo", "博客" ]
categories = [ "前端" ]
+++

:smile:
**hugo**生成静态网站，**sass+bulma**布局，**grunt**监控压缩css以及生成搜索索引文件，
放弃使用gulp是因为node-sass总是安装不上。全文搜索引擎使用**lunr.js**,因为lunr
不支持中文分词，故此使用**segment**分词配合使用[中文修改版lunr](https://github.com/codepiano/lunr.js)

-----------------------

全局安装grunt `npm install -g grunt-cli` 

grunt需要<a href="http://rubyinstaller.org/downloads/" target="_blank">ruby</a>
然后安装sass
```
gem install sass
gem install compass
```
