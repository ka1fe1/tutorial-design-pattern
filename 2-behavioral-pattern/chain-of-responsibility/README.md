# 责任链模式

## 定义

责任链模式是一种行为设计模式，它允许你将请求沿着处理者链进行发送。每个处理者收到请求后均可对请求进行处理，或将其向下传递。

## 结构

![责任链模式结构](https://img-blog.csdnimg.cn/a0e7b43a9c01423d9e38f84bcaa986c7.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAX2thaWZlaQ==,size_20,color_FFFFFF,t_70,g_se,x_16)

## 应用场景

- 当程序需要使用不同方式处理请求，且请求类型和顺序无法预测时
- 当必须按顺序执行多个处理者时
- 如果所需处理者及其顺序需要运行时改变时

## 应用举例

- 可用来解决参数多重校验的问题