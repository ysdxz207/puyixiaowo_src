+++
date = "2017-05-19T17:22:37+08:00"
title = "mysql主从复制配置"
tags = ["主从"]
categories = ["mysql"]
+++

#### 1.两台机配置：

	`192.168.137.46`//主库地址
	`192.168.137.250`//从库地址
		
	主库中有4张表：
	```
	db_tnews_dictionary
	db_tnews_member
	db_tnews_news
	db_tnews_manager
	```
	从库与主库相同
#### 2.主库配置：

- 修改配置文件

	`/etc/mysql/my.cnf`
	
	```
	[mysqld]
	server-id		= 1 #主库的服务器ID
	log_bin			= /var/log/mysql/mysql-bin.log #这里需要开启bin log
	binlog_do_db		= db_tnews_news  #需要同步的库
	binlog_do_db		= db_tnews_dictionary #需要同步的库
	binlog_do_db		= db_tnews_member #需要同步的库
	binlog_do_db		= db_tnews_manager #需要同步的库
	#binlog_ignore_db	= include_database_name #忽略的库
	```
	
- 重启mysql服务
	
	`service mysql restart` 
	> 这里因为我用的是mysql-server-5.5的所以是服务名是mysql
	
- 配置同步用户
	
	```
	mysql -u root -p #登录root用户
	mysql>GRANT REPLICATION SLAVE ON *.* to 'baker'@'192.168.137.250' identified by '123456';
	mysql>flush privileges;
	```
- 然后查看一下主服务器当前二进制日志名和偏移量，也就是同步的起始点
	```
	mysql> show master status;
	+------------------+----------+--------------------------------------------------------------------+------------------+
	| File             | Position | Binlog_Do_DB                                                       | Binlog_Ignore_DB |
	+------------------+----------+--------------------------------------------------------------------+------------------+
	| mysql-bin.000004 |      107 | db_tnews_news,db_tnews_dictionary,db_tnews_member,db_tnews_manager |                  |
	+------------------+----------+--------------------------------------------------------------------+------------------+
	1 row in set (0.00 sec)

	```
	
#### 3.从库配置

- 修改配置文件

	`/etc/mysql/my.cnf`
	```
	server-id		= 2
	
	```
- 重启mysql服务

	`service mysql restart`
- 配置从库中的主库地址
	```
	mysql -u root -p #root登录从库
	mysql>STOP SLAVE;
	mysql> CHANGE MASTER TO MASTER_HOST='192.168.137.46',
    MASTER_PORT=3306,
    MASTER_USER='baker',
    MASTER_PASSWORD='123456',
    MASTER_LOG_FILE='mysql-bin.000004',
    MASTER_LOG_POS=107;
	mysql>START SLAVE;
	```
	
> MASTER_LOG_FILE是之前show master status;查询到的File值
> MASTER_LOG_POS是之前show master status;查询到的Position值

- 查看从库进程

    `mysql> show slave status\G;`

    Slave_IO_Running: Yes
    Slave_SQL_Running: Yes
    两项是否都在运行
    以及Slave_IO_State的值(运行状态):Waiting for master to send event
    证明已配置完成

#### 4.测试
    主库：
    ```
    CREATE TABLE `test` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      PRIMARY KEY (`id`)
    );
    ```
    从库：
    `show slave status\G;`
    
> Last_IO_Errno: 0

> Last_IO_Error: 
   
> Last_SQL_Errno: 0
   
> Last_SQL_Error:

> 没有报错，查看从库test表被创建