+++
date = "2017-04-25T14:33:35+08:00"
title = "mysql 随机查询n条"
tags = ["随机查询"]
categories = ["mysql"]
+++

```mysql
DROP TABLE IF EXISTS test_random;
CREATE TABLE `test_random` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP FUNCTION IF EXISTS test_random_func;
CREATE FUNCTION test_random_func(num INT) RETURNS INT
BEGIN
	
	DECLARE n INT UNSIGNED DEFAULT 0;
	
	WHILE n < num DO
	
	INSERT INTO test_random(`create_time`) VALUES(NOW());

	SET n=n+1;
	END WHILE;
	RETURN n;
END;

SELECT test_random_func(200000) AS '数据条数';

EXPLAIN SELECT id AS 'RAND()查询结果' FROM test_random ORDER BY RAND() limit 0, 10;

SELECT 
    t.id '基于随机数查询结果'
FROM
    (SELECT 
        ROUND(RAND() * (SELECT 
                    MAX(id)
                FROM
                    test_random)) random_num,
            @num:=@num + 1
    FROM
        (SELECT @num:=0) AS a, test_random
    LIMIT 10) AS b,
    test_random AS t
WHERE
    b.random_num = t.id;
```