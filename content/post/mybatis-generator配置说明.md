+++
date = "2017-04-26T14:06:59+08:00"
title = "mybatis-generator配置说明"
tags = ["mybatis"]
categories = ["mybatis"]
+++

	mybatis generator config
```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE generatorConfiguration
  PUBLIC "-//mybatis.org//DTD MyBatis Generator Configuration 1.0//EN"
  "http://mybatis.org/dtd/mybatis-generator-config_1_0.dtd">


<!-- 根节点<generatorConfiguration> :节点没有任何属性，直接写节点即可，如下： -->
<generatorConfiguration>

    <!-- <properties> 引入外部文件 使用${property}调用 -->
    <!-- <properties resource="jdbc.properties"/> -->
    <!-- resource：指定**classpath**下的属性文件，使用类似com/myproject/generatorConfig.properties这样的属性值。 -->
    <!-- url：可以指定文件系统上的特定位置，例如file:///C:/myfolder/generatorConfig.properties -->


    <!-- <context>元素用于指定生成一组对象的环境。例如指定要连接的数据库，要生成对象的类型和要处理的数据库中的表。运行MBG的时候还可以指定要运行的<context>。 -->
    <!-- id:必选属性，该id属性可以在运行时的使用。 defaultModelType:这个属性定义了MBG如何生成*实体类*。 conditional:*这是默认值*,这个模型和下面的hierarchical类似，除了如果那个单独的类将只包含一个字段，将不会生成一个单独的类。 
        因此,如果一个表的主键只有一个字段,那么不会为该字段生成单独的实体类,会将该字段合并到基本实体类中。 flat:该模型为每一张表只生成一个实体类。这个实体类包含表中的所有字段。*这种模型最简单，推荐使用。* 
        hierarchical:如果表有主键,那么该模型会产生一个单独的主键实体类,如果表还有BLOB字段， 则会为表生成一个包含所有BLOB字段的单独的实体类,然后为所有其他的字段生成一个单独的实体类。 
        MBG会在所有生成的实体类之间维护一个继承关系。 targetRuntime:此属性用于指定生成的代码的运行时环境。该属性支持以下可选值： MyBatis3:*这是默认值* 
        MyBatis3Simple Ibatis2Java2 Ibatis2Java5 一般情况下使用默认值即可，有关这些值的具体作用以及区别请查看中文文档的详细内容。 
        introspectedColumnImpl:该参数可以指定扩展org.mybatis.generator.api.IntrospectedColumn该类的实现类。该属性的作用可以查看扩展MyBatis 
        Generator。 -->
    <context id="testTables" targetRuntime="MyBatis3">

        <!-- <property name="" value=""/> autoDelimitKeywords :当表名或者字段名为SQL关键字的时候，可以设置该属性为true，MBG会自动给表名或字段名添加分隔符。 
            beginningDelimiter ：由于beginningDelimiter和endingDelimiter的默认值为双引号(")，在Mysql中不能这么写，所以还要将这两个默认值改为**反单引号(`)**， 
            endingDelimiter ：<property name="beginningDelimiter" value="`"/> <property 
            name="endingDelimiter" value="`"/> javaFileEncoding ：属性javaFileEncoding设置要使用的Java文件的编码，默认使用当前平台的编码，只有当生产的编码需要特殊指定时才需要使用，一般用不到。 
            javaFormatter ：最后两个javaFormatter和xmlFormatter属性**可能会**很有用，如果你想使用模板来定制生成的java文件和xml文件的样式，你可以通过指定这两个属性的值来实现。 
            xmlFormatter ：最后两个javaFormatter和xmlFormatter属性**可能会**很有用，如果你想使用模板来定制生成的java文件和xml文件的样式，你可以通过指定这两个属性的值来实现。 -->

        <!-- <plugin> 元素用来定义一个插件。插件用于扩展或修改通过MyBatis Generator (MBG)代码生成器生成的代码。插件将按在配置中配置的顺序执行。 -->

        <!-- 生成注释信息 -->
        <commentGenerator>
            <!-- suppressAllComments:阻止生成注释，默认为false -->
            <!-- suppressDate:阻止生成的注释包含时间戳，默认为false -->
            <!-- 是否去除自动生成的注释 true：是 ： false:否 -->
            <property name="suppressAllComments" value="true" />
            <property name="suppressDate" value="true" />
        </commentGenerator>

        <!--<<jdbcConnection> 用于指定数据库连接信息，该元素必选，并且只能有一个。 -->
        <!--数据库连接的信息：驱动类、连接地址、用户名、密码 -->
        <jdbcConnection driverClass="com.mysql.jdbc.Driver"
            connectionURL="jdbc:mysql://localhost:3306/ssi" userId="root"
            password="root">
        </jdbcConnection>


        <!-- 这个元素的配置用来指定JDBC类型和Java类型如何转换。 -->
        <!-- 默认false，把JDBC DECIMAL 和 NUMERIC 类型解析为 Integer，为 true时把JDBC DECIMAL 
            和 NUMERIC 类型解析为java.math.BigDecimal -->
        <javaTypeResolver>
            <property name="forceBigDecimals" value="false" />
        </javaTypeResolver>

        <!-- targetProject:生成POJO类的位置 targetPackage:生成实体类存放的包名，一般就是放在该包下。实际还会受到其他配置的影响(<table>中会提到)。 
            targetProject:指定目标项目路径，使用的是文件系统的绝对路径。 -->
        <javaModelGenerator targetPackage="com.mybatis.pojo"
            targetProject="src/main/java">
            <!-- constructorBased:该属性只对MyBatis3有效，如果true就会使用构造方法入参，如果false就会使用setter方式。默认为false 
                enableSubPackages:如果true，MBG会根据catalog和schema来生成子包。如果false就会直接用targetPackage属性。默认为false。 
                immutable:该属性用来配置实体类属性是否可变，如果设置为true，那么constructorBased不管设置成什么，都会使用构造方法入参，并且不会生成setter方法。如果为false，实体类属性就可以改变。默认为false。 
                rootClass:设置所有实体类的基类。如果设置，需要使用类的全限定名称。并且如果MBG能够加载rootClass，那么MBG不会覆盖和父类中完全匹配的属性。匹配规则： 
                属性名完全相同 属性类型相同 属性有getter方法 属性有setter方法 trimStrings:是否对数据库查询结果进行trim操作，如果设置为true就会生成类似这样public 
                void setUsername(String username) {this.username = username == null ? null 
                : username.trim();}的setter方法。默认值为false -->
            <!-- enableSubPackages:是否让schema作为包的后缀 -->
            <property name="enableSubPackages" value="false" />
            <!-- 从数据库返回的值被清理前后的空格 -->
            <property name="trimStrings" value="true" />
        </javaModelGenerator>

        <!-- targetProject:mapper映射文件生成的位置 -->
        <sqlMapGenerator targetPackage="com.mybatis.mapper" targetProject="src/main/java">
            <!-- enableSubPackages:是否让schema作为包的后缀 -->
            <property name="enableSubPackages" value="false" />
        </sqlMapGenerator>

        <!-- targetPackage：mapper接口生成的位置 -->
        <javaClientGenerator type="XMLMAPPER" targetPackage="com.mybatis.mapper" targetProject="src/main/java">
            <!-- enableSubPackages:是否让schema作为包的后缀 -->
            <property name="enableSubPackages" value="false" />
        </javaClientGenerator>
        <!-- 指定数据库表 -->

        <table domainObjectName="User"  tableName="tb_user"></table>
        <table domainObjectName="Order"   tableName="tb_order"></table>
        <table domainObjectName="OrderdDetail"   tableName="tb_orderdetail"></table>
        <table domainObjectName="Item"   tableName="tb_item"></table>
		
		<table schema="&dbName;" tableName="agent" domainObjectName="Agent" enableCountByExample="false" enableUpdateByExample="false" enableDeleteByExample="false"
        enableSelectByExample="false" selectByExampleQueryId="false">
		  <property name="useActualColumnNames" value="false"/>
		  <generatedKey column="id" sqlStatement="MySql" identity="true" />
		  <!-- 重写字段类型，jdbcType为xxxMapper.xml中类型，javaType为java类型-->
		  <columnOverride column="gender" jdbcType="INTEGER" javaType="Integer" />
		  <columnOverride column="card_type" jdbcType="INTEGER" javaType="Integer" />
		  <columnOverride column="check" jdbcType="INTEGER" javaType="Integer" />
		  <columnOverride column="sys_flag" jdbcType="INTEGER" javaType="Integer" />
		</table> 
    </context>

</generatorConfiguration>
```