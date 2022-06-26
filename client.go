// Package pica 哔咔漫画Golang客户端
package pica

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var json jsoniter.API

func init() {
	extra.RegisterFuzzyDecoders()
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

// 哔咔的网址
const server = "https://picaapi.picacomic.com/"

// Client
// 客户端 struct
type Client struct {
	http.Client
	Token string
}

// 设置签名header
func (client *Client) header(req *http.Request) {
	// define
	var apiKey = "C69BAF41DA5ABD1FFEDC6D2FEA56B"
	var nonce = "b1ab87b4800d4d4590a11701b8551afa"
	// header
	var header = map[string]string{
		"api-key":           apiKey,
		"accept":            "application/vnd.picacomic.com.v1+json",
		"app-channel":       "2",
		"time":              strconv.FormatInt(time.Now().Unix(), 10),
		"nonce":             nonce,
		"signature":         "",
		"app-version":       "2.2.1.2.3.3",
		"app-uuid":          "defaultUuid",
		"app-platform":      "android",
		"app-build-version": "44",
		"Content-Type":      "application/json; charset=UTF-8",
		"User-Agent":        "okhttp/3.8.1",
		"authorization":     client.Token,
		"image-quality":     "original",
	}
	// sign
	var raw = strings.TrimPrefix(req.URL.RequestURI(), "/") + header["time"] + nonce + req.Method + apiKey
	raw = strings.ToLower(raw)
	h := hmac.New(sha256.New, []byte("~d}$Q7$eIni=V)9\\RK/P.RM4;9[7|@/CA}b~OW!3?EV`:<>M7pddUBL5n|0/*Cn"))
	h.Write([]byte(raw))
	header["signature"] = hex.EncodeToString(h.Sum(nil))
	// put in req
	if req != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}
}

// postToPica 向哔咔发送请求
func (client *Client) bodyRequestToPica(method string, path string, body interface{}) (*http.Request, error) {
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest(method, server+path, nil)
	} else {
		bodyBuff, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyStream := bytes.NewBuffer(bodyBuff)
		req, err = http.NewRequest(method, server+path, bodyStream)
	}
	if err != nil {
		return nil, err
	}
	client.header(req)
	return req, nil
}

// postToPica 向哔咔发送POST请求
func (client *Client) postToPica(path string, body interface{}) (*http.Request, error) {
	return client.bodyRequestToPica("POST", path, body)
}

// putToPica 向哔咔发送PUT请求
func (client *Client) putToPica(path string, body interface{}) (*http.Request, error) {
	return client.bodyRequestToPica("PUT", path, body)
}

// getToPica 向哔咔发送GET请求
func (client *Client) getToPica(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", server+path, nil)
	if err != nil {
		return nil, err
	}
	client.header(req)
	return req, nil
}

// getToPicaWithQuality 向哔咔发送GET请求, 并修改 "image-quality" 请求头
func (client *Client) getToPicaWithQuality(path string, quality ImageQuality) (*http.Request, error) {
	req, err := http.NewRequest("GET", server+path, nil)
	if err != nil {
		return nil, err
	}
	client.header(req)
	req.Header.Set("image-quality", string(quality))
	return req, nil
}

// responseFromPica 从哔咔接口返回体, 并解析异常信息
func responseFromPica[T any](client *Client, req *http.Request) (*Response[T], error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response Response[T]
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	if response.Code != 200 {
		return nil, errors.New(response.Message)
	}
	return &response, nil
}

// Register 注册新用户
func (client *Client) Register(dto RegisterDto) error {
	_, err := client.postToPica("auth/register", &dto)
	return err
}

// Login 登录, 登录无异常则注入TOKEN
func (client *Client) Login(username string, password string) error {
	req, err := client.postToPica("auth/sign-in", &LoginRequest{
		Email:    username,
		Password: password,
	})
	if err != nil {
		return err
	}
	response, err := responseFromPica[LoginResult](client, req)
	if err != nil {
		return err
	}
	client.Token = response.Data.Token
	return nil
}

// UserProfile 用户信息
func (client *Client) UserProfile() (*UserProfile, error) {
	req, err := client.getToPica("users/profile")
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[UserProfileResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.User, nil
}

// PunchIn 打哔卡
func (client *Client) PunchIn() (*PunchStatus, error) {
	req, err := client.postToPica("users/punch-in", nil)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[PunchResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Res, nil
}

// Categories 获取分类
func (client *Client) Categories() ([]Category, error) {
	req, err := client.getToPica("categories")
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[CategoriesResult](client, req)
	if err != nil {
		return nil, err
	}
	return response.Data.Categories, nil
}

// Comics 分类下的漫画
// category 为空字符串则为所有分类
func (client *Client) Comics(category string, tag string, creatorId string, chineseTeam string, sort Sort, page int) (*ComicsPage, error) {
	mUrl := "comics?"
	if len(category) > 0 {
		mUrl = mUrl + fmt.Sprintf("c=%s&", url.QueryEscape(category))
	}
	if len(tag) > 0 {
		mUrl = mUrl + fmt.Sprintf("t=%s&", url.QueryEscape(tag))
	}
	if len(creatorId) > 0 {
		mUrl = mUrl + fmt.Sprintf("ca=%s&", creatorId)
	}
	if len(chineseTeam) > 0 {
		mUrl = mUrl + fmt.Sprintf("ct=%s&", url.QueryEscape(chineseTeam))
	}
	req, err := client.getToPica(mUrl + "s=" + string(sort) + "&page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicsPageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comics, nil
}

// SearchComics 搜索漫画
// PS : 此接口并没有返回 PagesCount EpsCount
func (client *Client) SearchComics(categories []string, keyword string, sort Sort, page int) (*ComicsPage, error) {
	params := map[string]interface{}{
		"keyword": keyword,
		"sort":    sort,
	}
	if categories != nil && len(categories) > 0 {
		params["categories"] = categories
	}
	req, err := client.postToPica("comics/advanced-search?page="+strconv.Itoa(page), params)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicsPageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comics, nil
}

// SearchComicsInCategories 搜索漫画
func (client *Client) SearchComicsInCategories(keyword string, sort Sort, page int, categories []string) (*ComicsPage, error) {
	params := map[string]interface{}{}
	params["categories"] = categories
	params["keyword"] = keyword
	params["sort"] = sort
	req, err := client.postToPica("comics/advanced-search?page="+strconv.Itoa(page), params)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicsPageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comics, nil
}

// RandomComics 随机漫画
func (client *Client) RandomComics() ([]ComicSimple, error) {
	req, err := client.getToPica("comics/random")
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicsResult](client, req)
	if err != nil {
		return nil, err
	}
	return response.Data.Comics, nil
}

// Leaderboard 排行榜
func (client *Client) Leaderboard(leaderboardType string) ([]ComicSimple, error) {
	req, err := client.getToPica(fmt.Sprintf("comics/leaderboard?tt=%s&ct=VC", leaderboardType))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicsResult](client, req)
	if err != nil {
		return nil, err
	}
	return response.Data.Comics, nil
}

// ComicInfo 漫画详情
func (client *Client) ComicInfo(comicId string) (*ComicInfo, error) {
	req, err := client.getToPica("comics/" + comicId)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicInfoResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comic, nil
}

// ComicEpPage 漫画EP信息
func (client *Client) ComicEpPage(comicId string, page int) (*EpPage, error) {
	req, err := client.getToPica("comics/" + comicId + "/eps?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[EpPageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Eps, nil
}

// ComicPicturePage 漫画图片
func (client *Client) ComicPicturePage(comicId string, epOrder int, page int) (*ComicPicturePage, error) {
	req, err := client.getToPica("comics/" + comicId + "/order/" + strconv.Itoa(epOrder) + "/pages?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicPicturePageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Pages, nil
}

// ComicPicturePageWithQuality 漫画图片
func (client *Client) ComicPicturePageWithQuality(comicId string, epOrder int, page int, quality ImageQuality) (*ComicPicturePage, error) {
	req, err := client.getToPicaWithQuality("comics/"+comicId+"/order/"+strconv.Itoa(epOrder)+"/pages?page="+strconv.Itoa(page), quality)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicPicturePageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Pages, nil
}

// SwitchLike (取消)喜欢漫画
// 第一次喜欢，第二次是取消喜欢 action是最终结果
func (client *Client) SwitchLike(comicId string) (*string, error) {
	req, err := client.postToPica("comics/"+comicId+"/like", nil)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ActionResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Action, nil
}

// SwitchFavourite (取消)收藏漫画
// 第一次收藏，第二次是取消收藏 action是最终结果
func (client *Client) SwitchFavourite(comicId string) (*string, error) {
	req, err := client.postToPica("comics/"+comicId+"/favourite", nil)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ActionResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Action, nil
}

// FavouriteComics 收藏的漫画
func (client *Client) FavouriteComics(sort Sort, page int) (*ComicsPage, error) {
	req, err := client.getToPica("users/favourite?s=" + string(sort) + "&page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicsPageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comics, nil
}

// ComicCommentsPage 漫画的评论
func (client *Client) ComicCommentsPage(comicId string, page int) (*CommentsPage, error) {
	req, err := client.getToPica("comics/" + comicId + "/comments?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[CommentsResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comments, nil
}

func (client *Client) MyComments(page int) (*MyCommentsPage, error) {
	req, err := client.getToPica(fmt.Sprintf("users/my-comments?page=%d", page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[MyCommentsPageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comments, nil
}

// PostComment 对漫画进行评论, 但是评论后无法删除
func (client *Client) PostComment(comicId string, content string) error {
	req, err := client.postToPica(fmt.Sprintf("comics/%s/comments", comicId), map[string]string{
		"content": content,
	})
	if err != nil {
		return err
	}
	_, err = responseFromPica[interface{}](client, req)
	return err
}

// HideComment 哔咔API里的接口, 不知道做什么用的, 推测是管理员用接口
func (client *Client) HideComment(commentId string) error {
	req, err := client.postToPica(fmt.Sprintf("comments/%s/delete", commentId), nil)
	if err != nil {
		return err
	}
	_, err = responseFromPica[interface{}](client, req)
	return err
}

// CommentChildren 获取子评论
func (client *Client) CommentChildren(commentId string, page int) (*CommentChildrenPage, error) {
	req, err := client.getToPica(fmt.Sprintf("comments/%s/childrens?page=%d", commentId, page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[CommentChildrenResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comments, nil
}

// PostChildComment 对漫画/游戏的评论进行回复(子评论), 但是评论后无法删除
func (client *Client) PostChildComment(commentId string, content string) error {
	req, err := client.postToPica(fmt.Sprintf("comments/%s", commentId), map[string]string{
		"content": content,
	})
	if err != nil {
		return err
	}
	_, err = responseFromPica[interface{}](client, req)
	return err
}

// SwitchLikeComment (取消)喜欢评论/子评论
// 第一次喜欢，第二次是取消喜欢 action是最终结果 ( ActionLike or ActionUnlike )
func (client *Client) SwitchLikeComment(commentId string) (*string, error) {
	req, err := client.postToPica("comments/"+commentId+"/like", nil)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ActionResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Action, nil
}

// ComicRecommendation 看了这个本子的也在看
func (client *Client) ComicRecommendation(comicId string) ([]ComicSimple, error) {
	req, err := client.getToPica("comics/" + comicId + "/recommendation")
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[ComicsResult](client, req)
	if err != nil {
		return nil, err
	}
	return response.Data.Comics, nil
}

// HotKeywords 大家都在搜
func (client *Client) HotKeywords() ([]string, error) {
	req, err := client.getToPica("keywords")
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[HotKeywordsResutl](client, req)
	if err != nil {
		return nil, err
	}
	return response.Data.Keywords, nil
}

// LeaderboardOfKnight 骑士榜
func (client *Client) LeaderboardOfKnight() ([]Knight, error) {
	req, err := client.getToPica("comics/knight-leaderboard")
	if err != nil {
		panic(err)
	}
	response, err := responseFromPica[LeaderboardOfKnightResult](client, req)
	if err != nil {
		return nil, err
	}
	return response.Data.Users, nil
}

// GamePage 游戏列表
func (client *Client) GamePage(page int) (*GamePage, error) {
	req, err := client.getToPica("games?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[GamePageResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Games, nil
}

// GameInfo 游戏详情
func (client *Client) GameInfo(gameId string) (*GameInfo, error) {
	req, err := client.getToPica("games/" + gameId)
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[GameResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Game, nil
}

// GameCommentsPage 游戏评论分页
func (client *Client) GameCommentsPage(gameId string, page int) (*GameCommentsPage, error) {
	req, err := client.getToPica("games/" + gameId + "/comments?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[GameCommentsResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comments, nil
}

// PostGameComment 对游戏进行评论, 但是评论后无法删除
func (client *Client) PostGameComment(gameId string, content string) error {
	req, err := client.postToPica(fmt.Sprintf("games/%s/comments", gameId), map[string]string{
		"content": content,
	})
	if err != nil {
		return err
	}
	_, err = responseFromPica[interface{}](client, req)
	return err
}

// GameCommentChildren 游戏评论的回复分页 (和漫画接口是同一个, 只有"_comic/_game"字段不一样)
func (client *Client) GameCommentChildren(commentId string, page int) (*GameCommentChildrenPage, error) {
	req, err := client.getToPica(fmt.Sprintf("comments/%s/childrens?page=%d", commentId, page))
	if err != nil {
		return nil, err
	}
	response, err := responseFromPica[GameCommentChildrenResult](client, req)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comments, nil
}

// UpdatePassword 修改密码
func (client *Client) UpdatePassword(oldPassword string, newPassword string) error {
	body := map[string]string{
		"old_password": oldPassword,
		"new_password": newPassword,
	}
	req, err := client.putToPica("users/password", body)
	if err != nil {
		return err
	}
	_, err = responseFromPica[interface{}](client, req)
	return err
}

// UpdateSlogan 修改签名
func (client *Client) UpdateSlogan(slogan string) error {
	body := map[string]string{
		"slogan": slogan,
	}
	req, err := client.putToPica("users/profile", body)
	if err != nil {
		return err
	}
	_, err = responseFromPica[interface{}](client, req)
	return err
}

// UpdateAvatar 修改头像
// 请压缩头像成正方形, 200x200,并尽量减少图片体积, 编码必须为JPEG
func (client *Client) UpdateAvatar(jpegBytes []byte) error {
	body := map[string]string{
		"avatar": "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(jpegBytes),
	}
	req, err := client.putToPica("users/avatar", body)
	if err != nil {
		return err
	}
	_, err = responseFromPica[interface{}](client, req)
	return err
}

func (client *Client) Collections() ([]Collection, error) {
	buff, err := client.getToPica("collections")
	if err != nil {
		panic(err)
	}
	var response CollectionsResponse
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.Collections, nil
}
