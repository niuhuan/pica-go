PIKA API
===========
一个Golang版的哔卡客户端

## 实现功能

- [x] 用户
  - [x] 登录 / 获取用户信息
  - [x] 打哔卡
- [x] 漫画
  - [x] 分类 / 搜索
  - [x] 获取EP / 获取图片
  - [x] 收藏 / 喜欢
  - [x] 获取评论
  - [ ] 评论
  - [x] 看这个本子的也在看
- [x] 游戏
- [x] 社交
  - [x] 大家都在搜
- [x] 网络
  - [x] 分流 / 代理 (examples)

## 使用方法

下载zip或克隆仓库, 然后在go.mod中加入以下内容, 并调用 client := pica.Client{}
```
require pica v0.0.0
replace pica v0.0.0 => ./pica
```
