PICA-GO GOLANG
===========
GoLang哔卡漫画客户端

## 实现功能

- [x] 用户
    - [x] 登录 / 获取用户信息 / 注册 / 打哔卡 / 我的评论
- [x] 漫画
    - [x] 分类 / 搜索 / 大家都在搜 / 随机本子 / 排行榜
    - [x] 收藏 / 喜欢 / 获取EP / 获取图片
    - [x] 看这个本子的也在看 / 评论 / 获取评论 / 子评论
- [x] 游戏
- [x] 社交
    - [x] 大家都在搜
- [x] 网络
    - [x] 分流 / 代理 ([examples](https://github.com/niuhuan/pica-go/blob/master/examples/examples.go))

## 使用方法

```text
package main

import "github.com/niuhuan/pica-go"

func main(){
  client := pica.Client{}  
  err := client.Login(username, password)
  comicsPage, err := client.Comics("", "", "", "", pica.SortDefault, 1)
}
```

## 其他语言

- 请参考 [picacomic](https://github.com/AnkiKong/picacomic) 搭建自己的客户端
