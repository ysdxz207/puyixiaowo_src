+++
date = "2017-06-29T14:53:42+08:00"
title = "centos下安装redis"
tags = []
categories = ["linux"]
+++

1.到[http://download.redis.io/releases/](http://download.redis.io/releases/)下找到需要的版本
2.
```
wget http://download.redis.io/releases/redis-3.2.9.tar.gz
tar -zxf http://download.redis.io/releases/redis-3.2.9.tar.gz
cd redis-3.2.9.tar.gz
make
make install
cp redis.conf /etc
```


3.用户及日志
```
sudo useradd redis  
sudo mkdir -p /var/lib/redis  
sudo mkdir -p /var/log/redis  
sudo chown redis.redis /var/lib/redis ＃db文件放在这里，要修改redis.conf  
sudo chown redis.redis /var/log/redis
```

>需要密码认证则修改redis.conf中requirepass密码
>需要修改redis.conf,将 daemonize no 改为 daemonize yes才能后台运行redis
4.启动脚本
```
vim /etc/init.d/redis
```
代码如下，粘贴进去
```
###########################  
PATH=/usr/local/bin:/sbin:/usr/bin:/bin  
     
REDISPORT=6379  
EXEC=/usr/local/bin/redis-server  
REDIS_CLI=/usr/local/bin/redis-cli  
     
PIDFILE=/var/run/redis.pid  
CONF="/etc/redis.conf"  
     
case "$1" in  
    start)  
        if [ -f $PIDFILE ]  
        then  
                echo "$PIDFILE exists, process is already running or crashed"  
        else  
                echo "Starting Redis server..."  
                $EXEC $CONF  
        fi  
        if [ "$?"="0" ]   
        then  
              echo "Redis is running..."  
        fi  
        ;;  
    stop)  
        if [ ! -f $PIDFILE ]  
        then  
                echo "$PIDFILE does not exist, process is not running"  
        else  
                PID=$(cat $PIDFILE)  
                echo "Stopping ..."  
                $REDIS_CLI -p $REDISPORT SHUTDOWN  
                while [ -x ${PIDFILE} ]  
               do  
                    echo "Waiting for Redis to shutdown ..."  
                    sleep 1  
                done  
                echo "Redis stopped"  
        fi  
        ;;  
   restart|force-reload)  
        ${0} stop  
        ${0} start  
        ;;  
  *)  
    echo "Usage: /etc/init.d/redis {start|stop|restart|force-reload}" >&2  
        exit 1  
esac  
##############################  
```

执行权限`chmod +x /etc/init.d/redis`
开机启动`sudo chkconfig redis on`

5.redis启动和停止
```
service redis start   #或者 /etc/init.d/redis start  
service redis stop   #或者 /etc/init.d/redis stop  
```

6.测试
```
redis-cli
auth 123456
........
```