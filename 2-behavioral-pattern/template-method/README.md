# 模板方法

## 定义

在超类中定义了一个算法框架，允许子类在不修改结构的情况下，重写算法的特定步骤

## 结构

![模板方法结构](https://img-blog.csdnimg.cn/f8178f432f1c4edf93ca23e39724cd35.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAX2thaWZlaQ==,size_20,color_FFFFFF,t_70,g_se,x_16)

## 应用场景

- 当你只希望客户端扩展某个特定算法步骤， 而不是整个算法或其结构时
- 当多个类的算法除一些细微不同之外几乎完全一样时