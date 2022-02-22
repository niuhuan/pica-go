package main

import (
	"context"
	"encoding/json"
	"github.com/niuhuan/pica-go"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

// 异常的说明
// token失效则会抛出异常 unauthorized
// 其他异常则是异常信息
// 网络异常则直接返回
func main() {
}

// ExampleNewSocks5ProxyClient
// 使用代理创建客户端
// 使用http的代理URL 例如 "socks5://127.0.0.1:1080", "http://localhost:1087"
func ExampleNewSocks5ProxyClient() *pica.Client {
	proxyUrl, proxyErr := url.Parse("socks5://127.0.0.1:1080")
	proxy := func(_ *http.Request) (*url.URL, error) {
		return proxyUrl, proxyErr
	}
	client := &pica.Client{}
	client.Transport = &http.Transport{
		Proxy:                 proxy,
		TLSHandshakeTimeout:   time.Second * 10,
		ExpectContinueTimeout: time.Second * 10,
		ResponseHeaderTimeout: time.Second * 10,
		IdleConnTimeout:       time.Second * 10,
	}
	return client
}

// ExampleNewClientAllocateIp
// 使用分流创建客户端
// 例如 "172.67.7.24:443" "104.20.180.50:443" "172.67.208.169:443"
func ExampleNewClientAllocateIp() *pica.Client {
	var pattern, _ = regexp.Compile("^.+picacomic\\.com:\\d+$")
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	client := &pica.Client{}
	client.Transport = &http.Transport{
		TLSHandshakeTimeout:   time.Second * 10,
		ExpectContinueTimeout: time.Second * 10,
		ResponseHeaderTimeout: time.Second * 10,
		IdleConnTimeout:       time.Second * 10,
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			if pattern.MatchString(addr) {
				addr = "104.20.180.50:443"
			}
			return dialer.DialContext(ctx, network, addr)
		},
	}
	return client
}

// ExampleLogin
// 登录, 登录后页可以保留TOKEN下次使用
func ExampleLogin() string {
	client := ExampleNewClientAllocateIp()
	err := client.Login("username", "password")
	if err != nil {
		panic(err)
	}
	return client.Token
}

// ExampleSetToken
// 设置Token
func ExampleSetToken(client *pica.Client) {
	client.Token = "eyJhbGciOiJIUz..."
}

// ExampleCategories (需要登录)
// 获取所有分类
func ExampleCategories(client *pica.Client) {
	categories, err := client.Categories()
	if err != nil {
		panic(err)
	}
	buff, _ := json.Marshal(&categories)
	println(string(buff))
}

// ExampleCategoryComics (需要登录)
// 获取分类下的漫画
func ExampleCategoryComics(client *pica.Client) {
	comicPage, err := client.Comics("", "", "", "", pica.SortDefault, 1)
	if err != nil {
		panic(err)
	}
	buff, _ := json.Marshal(comicPage)
	println(string(buff))
}

// ExampleComicInfo (需要登录)
// 获取漫画的信息
func ExampleComicInfo(client *pica.Client) {
	c, err := client.ComicInfo("60e9bd1c9172eb531c491359")
	if err != nil {
		panic(err)
	}
	buff, _ := json.Marshal(c)
	println(string(buff))
}

// ExampleComicEpPage (需要登录)
// 获取漫画的ep
func ExampleComicEpPage(client *pica.Client) {
	c, err := client.ComicEpPage("60e9bd1c9172eb531c491359", 1)
	if err != nil {
		panic(err)
	}
	buff, _ := json.Marshal(c)
	println(string(buff))
}

// ExampleComicPicturePage (需要登录)
// 获取漫画的图片 (原图)
func ExampleComicPicturePage(client *pica.Client) {
	c, err := client.ComicPicturePage("60e9bd1c9172eb531c491359", 1, 1)
	if err != nil {
		panic(err)
	}
	buff, _ := json.Marshal(c)
	println(string(buff))
}

// ExampleComicPicturePageWithQuality (需要登录)
// 获取漫画的图片 (选择质量)
func ExampleComicPicturePageWithQuality(client *pica.Client) {
	c, err := client.ComicPicturePageWithQuality("60e9bd1c9172eb531c491359", 1, 1, pica.ImageQualityMedium)
	if err != nil {
		panic(err)
	}
	buff, _ := json.Marshal(c)
	println(string(buff))
}

// ExampleSearchComics (需要登录)
// 搜索漫画
func ExampleSearchComics(client *pica.Client) {
	comicPage, err := client.SearchComics(nil, "难言之隐", pica.SortDefault, 1)
	if err != nil {
		panic(err)
	}
	buff, _ := json.Marshal(comicPage)
	println(string(buff))
}

// ExampleUpdateAvatar (需要登录)
// 修改头像
func ExampleUpdateAvatar(client *pica.Client) {
	buff, err := ioutil.ReadFile("1.jpg")
	if err != nil {
		panic(err)
	}
	err = client.UpdateAvatar(buff)
	if err != nil {
		panic(err)
	}
	println("OK")
}