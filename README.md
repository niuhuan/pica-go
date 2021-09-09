PIKA API
===========
一个Golang版的哔卡客户端

## 实现功能

- [x] 用户
  - [x] 登录 / 获取用户信息 / 注册 / 打哔卡 / 我的评论
- [x] 漫画
  - [x] 分类 / 搜索 / 大家都在搜 / 随机本子 / 排行榜
  - [x] 收藏 / 喜欢 / 获取EP / 获取图片 /
  - [x] 看这个本子的也在看 / 评论 / 获取评论 / 子评论
- [x] 游戏
- [x] 社交
  - [x] 大家都在搜
- [x] 网络
  - [x] 分流 / 代理 (examples)

## 使用方法

- 将pica-golang添加到您git仓库的子模块
  ```shell
  git module add -b master https://github.com/niuhuan/pica-golang.git
  ```
- 在go.mod中加入以下内容
  ```
  require pica v0.0.0
  replace pica v0.0.0 => ./pica-golang
  ```
- 调用客户端
- ```text
  var client = pica.Client{}
  ```

## 其他语言

- 请参考 [picacomic](https://github.com/AnkiKong/picacomic) 搭建自己的客户端
