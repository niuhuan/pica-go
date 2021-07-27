PIKA API
===========
一个Golang版的哔卡客户端

## 实现功能


- [x] 用户
  - [x] 登录
  - [ ] 其他
- [x] 漫画
  - [x] 分类
  - [x] 搜索
  - [x] EP
  - [x] 获取图片
- [x] 网络
  - [x] 分流 / 代理 (examples)

## 使用方法

下载zip或克隆仓库, 然后在go.mod中加入以下内容, 并调用 client := pica.Client{}
```
require pica v0.0.0
replace pica v0.0.0 => ./pica
```
