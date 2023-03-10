# go项目学习记录

go语言文档地址：https://books.studygolang.com/gopl-zh/ch6/ch6-01.html

## 1 go语言学习

### 1.1 接口

**1、接口是合约**

- **接口类型是对其他类型行为的抽象和概括；**
  - 接口类型是一种抽象的类型。
  - 它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；
  - 它们只会表现出它们自己的方法

**2、接口类型**

- **接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。**
  - 接口也可以内嵌，类似于结构体

**3、实现接口的条件**

- **一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口**
- **空接口类型**
  - ![image-20230224204442528](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224204442528.png)

### 1.2 结构体

- **1、结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员**
  - ![image-20230224210130552](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224210130552.png)
- **2、空结构体**
  - ![image-20230224205015026](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224205015026.png)
- **3、结构体嵌入和匿名对象**
  - ![image-20230224205150497](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224205150497.png)



### 1.3 方法

- **1、方法声明**
  - ![image-20230224205522927](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224205522927.png)
  - ![image-20230224205504528](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224205504528.png)
- **2、基于指针对象的方法**
  - ![image-20230224205712387](.assets/image-20230224205712387.png)
- **3、注意**
  - ![image-20230224210006047](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224210006047.png)

### 1.4 json

- **1、结构体内成员大写字母，才默认导出，只有导出的结构体成员才会被编码；**

- **2、omitempty选项：**
  - 表示Go语言中的成员为空或零值时不生产该JSON对象（这里false是零值，相对于Bool类型）





## 2 Viper

### 2.1 yaml文件类型

* **1、什么是YAML**
  * ![image-20230224215532055](https://raw.githubusercontent.com/caterpillar-0/picture/main/image-20230224215532055.png)





## 3 go利用multipart/from_data实现文件的上传与下载

### 3.1 form简介

> - 1、form（表单），是HTML标记语法的重要语法元素；
> - 2、HTTP通常作为HTML传输的载体；
>   - ![image-20230301212910069](.assets/image-20230301212910069.png)





### 3.2 publish/action

> - multipart/form-data
>   - ![image-20230301214927478](.assets/image-20230301214927478.png)
>   - 而且是多表单，分别是title,token和data
> - post，表单数据在body里
>   - ![image-20230301214915911](.assets/image-20230301214915911.png)





## 4 gorm

中文文档地址：https://gorm.cn/zh_CN/docs/preload.html

### 4.1 关联 

#### 4.1.1 belongs to

- video视频belongs to 用户user;
  - ![image-20230302211812528](.assets/image-20230302211812528.png)
  - ![image-20230302211850512](.assets/image-20230302211850512.png)
- **gorm通过preload、Joins预加载belongs to关联的记录，**
- **preload预加载**
  - ![image-20230302212036836](.assets/image-20230302212036836.png)

- **注意：gorm中一对多关系，belongs to && have many**
  - ![image-20230303192002177](.assets/image-20230303192002177.png)





















