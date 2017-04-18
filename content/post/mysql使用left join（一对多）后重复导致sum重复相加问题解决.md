+++
date = "2015-10-23T11:43:30+08:00"
title = "mysql使用left join（一对多）后重复导致sum重复相加问题解决"
tags = [ "mysql", "一对多", "sum", "重复" ]
categories = [
  "mysql"
]
+++

http://stackoverflow.com/questions/2436284/mysql-sum-for-distinct-rows

 

示例：
```
sum后总和 = 每一组的和*重复次数

=>每一组的sum和 = sum后总和 / 重复次数

 

重复次数= 查询总条数 / 每一组的条数

而每一组的sum和是我们所要的结果

所以 每一组的sum和 = sum后总和 * 每一组的条数 / 查询总条数
```