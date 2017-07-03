+++
date = "2017-06-29T21:46:29+08:00"
title = "centos下安装mysql"
tags = [""]
categories = ["linux"]
+++

```
wget http://dev.mysql.com/get/mysql-community-release-el7-5.noarch.rpm
rpm -ivh mysql-community-release-el7-5.noarch.rpm

```

默认安装的是　mysql5.6的，如果要安装　mysql 5.5 或mysql 5.7 需要修改　/etc/yum.repos.d/mysql-community.repo 里把　5.5或5.7开启　把　5.6禁用
然后安装
`yum install mysql-community-server`
成功安装之后重启mysql服务
`service mysqld restart`

设置密码
`set password for 'root'@'localhost' = password('mypasswd');`
远程访问
```
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'mypasswd' WITH GRANT OPTION;
flush privileges;
```
