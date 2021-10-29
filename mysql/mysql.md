
#### MySql 之一条查询sql的执行过程


##### mysql -h ${host} -p ${port} -u${name} -p 

##### 连接器 
去权限表中查权限 保持wait_time的长链接  超时重新连接

##### 查询缓存
成功连接 执行sql成功查询 会将sql作为key 查询结果作为value 存入缓存 下次再查询会直接查询缓存

但是update操作会清空该表的所有缓存  活跃更新的表不适合做查询缓存 

mysql 8.0 版本已经把缓存功能完全移除

##### 分析器

sql语句进行 “词法分析” 和 “语法分析”

“词法分析” : 识别操作类型 识别表名 识别字段名

“语法分析”: 检查sql语句 语法


##### 优化器

某表有多个索引的时候 决定用哪一个索引

多关联（join）查询的时候，决定关联的顺序

select * from user u join  score s using(ID)  where u.id=20 and s.scores=80;

既可以先从表user里面取出id=20的记录的ID值，再根据ID值关联到表socre，再判断score表里scores的值是否等于80。
也可以先从表score里面取出scores=80的记录的ID值，再根据ID值关联到user，再判断user表里面id的值是否等于20。

依据执行效率选择


##### 执行器

先检查权限 

默认引擎还是 指定的引擎 默认innodb

判断数据项 该行不符合 则继续下一行

直到取到满足所有查询条件的行 
执行器将查询结果返回给客户端

####
mysql -h$ip -P$端口 -u$登录名 -p 

mysql -h ${host} -p ${port} -u${name} -p 