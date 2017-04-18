+++
date = "2017-02-21T11:11:26+08:00"
title = "mybatis延迟加载导致dubbo值为null解决办法"
tags = [ "mybatis", "dubbo", "延迟加载" ]
categories = [
  "mybatis"
]
+++

场景条件：

 

1、bean里有两个属性：
```java
private String subjectCategoryName;//分类名
private String className;//班级分类名称
```

2、使用mybatis嵌套查询：
```xml
<association property="className" fetchType="eager" javaType="java.lang.String" column="{id=class_id}" select="selectClassName" />
<association property="subjectCategoryName" fetchType="eager" javaType="java.lang.String" column="{id=subject_category}" select="selectCategoryName" />

```

3、查询报错：
```java
Caused by: com.alibaba.com.caucho.hessian.io.HessianProtocolException: com.alibaba.com.caucho.hessian.io.ObjectDeserializer: unexpected object java.lang.String (id)
    at com.alibaba.com.caucho.hessian.io.AbstractDeserializer.error(AbstractDeserializer.java:108)
    at com.alibaba.com.caucho.hessian.io.AbstractDeserializer.readMap(AbstractDeserializer.java:95)
    at com.alibaba.com.caucho.hessian.io.Hessian2Input.readObject(Hessian2Input.java:1550)
    at com.alibaba.com.caucho.hessian.io.JavaDeserializer$ObjectFieldDeserializer.deserialize(JavaDeserializer.java:396)
    ... 47 more
```

4、debug发现service里有值，但无法进入action

5、google 搜索mybatis association dubbo找到https://github.com/alibaba/dubbo/issues/156



于是明白了原因，将mybatis设置中<setting name="lazyLoadingEnabled" value="true" />改为false，发现确实没报错了，也有数据了，于是把mybatis 设置改回true，然后更改如下代码
```xml
<association property="subjectCategoryName" fetchType="eager" javaType="java.lang.String" column="{id=subject_category}" select="selectCategoryName" />
<association property="className" fetchType="eager" javaType="java.lang.String" column="{id=class_id}" select="selectClassName" 

```
增加了fetchType="eager"问题解决