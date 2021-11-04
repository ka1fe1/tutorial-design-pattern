# 策略模式

## 定义

策略模式是一种行为模式，它能让你定义了一系列算法，并将每种算法分别定义在独立的类中，
以使算法对象能够相互替换。

## 结构

![策略模式结构图](https://img-blog.csdnimg.cn/98e4c11d95f24fc18ceda9473d82c0de.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAX2thaWZlaQ==,size_20,color_FFFFFF,t_70,g_se,x_16)

## 应用场景

- 当你想使用对象中各种不同的算法变体，并希望在运行中时切换算法时
- 当类中使用了复杂的条件运算符以在同一算法的不同变体中切换时

## 应用举例

- 可用来解决多条件判断执行的问题
