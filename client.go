// Package pica 哔咔漫画Golang客户端
package pica

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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

// 向哔咔发送POST请求
func (client *Client) postToPica(path string, body interface{}) ([]byte, error) {
	var req *http.Request
	var err error
	if body == nil {
		req, err = http.NewRequest("POST", server+path, nil)
	} else {
		bodyBuff, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyStream := bytes.NewBuffer(bodyBuff)
		req, err = http.NewRequest("POST", server+path, bodyStream)
	}
	if err != nil {
		return nil, err
	}
	client.header(req)
	return client.responseFromPica(req)
}

// 向哔咔发送GET请求
func (client *Client) getToPica(path string) ([]byte, error) {
	req, err := http.NewRequest("GET", server+path, nil)
	if err != nil {
		return nil, err
	}
	client.header(req)
	return client.responseFromPica(req)
}

// 向哔咔发送GET请求, 并修改 "image-quality" 请求头
func (client *Client) getToPicaWithQuality(path string, quality string) ([]byte, error) {
	req, err := http.NewRequest("GET", server+path, nil)
	if err != nil {
		return nil, err
	}
	client.header(req)
	req.Header.Set("image-quality", quality)
	return client.responseFromPica(req)
}

// 从哔咔接口返回体, 并解析异常信息
func (client *Client) responseFromPica(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	buff, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response Response
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	if response.Code != 200 {
		return nil, errors.New(response.Message)
	}
	return buff, nil
}

// Register 注册新用户
func (client *Client) Register(dto RegisterDto) error {
	_, err := client.postToPica("auth/register", &dto)
	return err
}

// Login 登录, 登录无异常则注入TOKEN
func (client *Client) Login(username string, password string) error {
	buff, err := client.postToPica("auth/sign-in", &LoginRequest{
		Email:    username,
		Password: password,
	})
	if err != nil {
		return err
	}
	var loginResponse LoginResponse
	err = json.Unmarshal(buff, &loginResponse)
	if err != nil {
		return err
	}
	client.Token = loginResponse.Data.Token
	return nil
}

// UserProfile 用户信息
func (client *Client) UserProfile() (*UserProfile, error) {
	buff, err := client.getToPica("users/profile")
	if err != nil {
		return nil, err
	}
	var userProfileResponse UserProfileResponse
	err = json.Unmarshal(buff, &userProfileResponse)
	if err != nil {
		return nil, err
	}
	return &userProfileResponse.Data.User, nil
}

// PunchIn 打哔卡
func (client *Client) PunchIn() (*PunchStatus, error) {
	buff, err := client.postToPica("users/punch-in", nil)
	if err != nil {
		return nil, err
	}
	var response PunchResponse
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data.Res, nil
}

func (client *Client) MyComments(page int) (*MyCommentsPage, error) {
	buff, err := client.getToPica(fmt.Sprintf("users/my-comments?page=%d", page))
	if err != nil {
		return nil, err
	}
	buff = replaceStringPage(buff)
	var response MyCommentsPageResponse
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comments, nil
}

// PostComment 对漫画进行评论, 但是评论后无法删除
func (client *Client) PostComment(comicId string, content string) error {
	_, err := client.postToPica(fmt.Sprintf("comics/%s/comments", comicId), map[string]string{
		"content": content,
	})
	return err
}

// HideComment 哔咔API里的接口, 不知道做什么用的, 推测是管理员用接口
func (client *Client) HideComment(commentId string) error {
	_, err := client.postToPica(fmt.Sprintf("comments/%s/delete", commentId), nil)
	return err
}

// CommentChildren 获取子评论
func (client *Client) CommentChildren(commentId string, page int) (*CommentChildrenPage, error) {
	buff, err := client.getToPica(fmt.Sprintf("comments/%s/childrens?page=%d", commentId, page))
	if err != nil {
		return nil, err
	}
	buff = replaceStringPage(buff)
	var response CommentChildrenResponse
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data.Comments, nil
}

// PostChildComment 对漫画评论进行回复(子评论), 但是评论后无法删除
func (client *Client) PostChildComment(commentId string, content string) error {
	_, err := client.postToPica(fmt.Sprintf("comments/%s", commentId), map[string]string{
		"content": content,
	})
	return err
}

// Categories 获取分类
func (client *Client) Categories() ([]Category, error) {
	buff, err := client.getToPica("categories")
	if err != nil {
		return nil, err
	}
	var response CategoriesResponse
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	return response.Data.Categories, nil
}

// Comics 分类下的漫画
// category 为空字符串则为所有分类
func (client *Client) Comics(category string, tag string, creatorId string, chineseTeam string, sort string, page int) (*ComicsPage, error) {
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
	buff, err := client.getToPica(mUrl + "s=" + sort + "&page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	var comicsResponse ComicsPageResponse
	err = json.Unmarshal(buff, &comicsResponse)
	if err != nil {
		return nil, err
	}
	return &comicsResponse.Data.Comics, nil
}

// SearchComics 搜索漫画
func (client *Client) SearchComics(categories []string, keyword string, sort string, page int) (*ComicsPage, error) {
	params := map[string]interface{}{
		"keyword": keyword,
		"sort":    sort,
	}
	if categories != nil && len(categories) > 0 {
		params["categories"] = categories
	}
	buff, err := client.postToPica("comics/advanced-search?page="+strconv.Itoa(page), params)
	if err != nil {
		return nil, err
	}
	var comicsResponse ComicsPageResponse
	err = json.Unmarshal(buff, &comicsResponse)
	if err != nil {
		return nil, err
	}
	return &comicsResponse.Data.Comics, nil
}

// SearchComicsInCategories 搜索漫画
func (client *Client) SearchComicsInCategories(keyword string, sort string, page int, categories []string) (*ComicsPage, error) {
	params := map[string]interface{}{}
	params["categories"] = categories
	params["keyword"] = keyword
	params["sort"] = sort
	buff, err := client.postToPica("comics/advanced-search?page="+strconv.Itoa(page), params)
	if err != nil {
		return nil, err
	}
	var comicsResponse ComicsPageResponse
	err = json.Unmarshal(buff, &comicsResponse)
	if err != nil {
		return nil, err
	}
	return &comicsResponse.Data.Comics, nil
}

// RandomComics 随机漫画
func (client *Client) RandomComics() ([]ComicSimple, error) {
	buff, err := client.getToPica("comics/random")
	if err != nil {
		return nil, err
	}
	var comicsResponse ComicsResponse
	err = json.Unmarshal(buff, &comicsResponse)
	if err != nil {
		return nil, err
	}
	return comicsResponse.Data.Comics, nil
}

// Leaderboard 排行榜
func (client *Client) Leaderboard(leaderboardType string) ([]ComicSimple, error) {
	buff, err := client.getToPica(fmt.Sprintf("comics/leaderboard?tt=%s&ct=VC", leaderboardType))
	if err != nil {
		return nil, err
	}
	var comicsResponse ComicsResponse
	err = json.Unmarshal(buff, &comicsResponse)
	if err != nil {
		return nil, err
	}
	return comicsResponse.Data.Comics, nil
}

// ComicInfo 漫画详情
func (client *Client) ComicInfo(comicId string) (*ComicInfo, error) {
	buff, err := client.getToPica("comics/" + comicId)
	if err != nil {
		return nil, err
	}
	var comicInfoResponse ComicInfoResponse
	err = json.Unmarshal(buff, &comicInfoResponse)
	if err != nil {
		return nil, err
	}
	return &comicInfoResponse.Data.Comic, nil
}

// ComicEpPage 漫画EP信息
func (client *Client) ComicEpPage(comicId string, page int) (*EpPage, error) {
	buff, err := client.getToPica("comics/" + comicId + "/eps?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	var epPageResponse EpPageResponse
	err = json.Unmarshal(buff, &epPageResponse)
	if err != nil {
		return nil, err
	}
	return &epPageResponse.Data.Eps, nil
}

// ComicPicturePage 漫画图片
func (client *Client) ComicPicturePage(comicId string, epOrder int, page int) (*ComicPicturePage, error) {
	buff, err := client.getToPica("comics/" + comicId + "/order/" + strconv.Itoa(epOrder) + "/pages?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	var epPageResponse ComicPicturePageResponse
	err = json.Unmarshal(buff, &epPageResponse)
	if err != nil {
		return nil, err
	}
	return &epPageResponse.Data.Pages, nil
}

// ComicPicturePageWithQuality 漫画图片
func (client *Client) ComicPicturePageWithQuality(comicId string, epOrder int, page int, quality string) (*ComicPicturePage, error) {
	buff, err := client.getToPicaWithQuality("comics/"+comicId+"/order/"+strconv.Itoa(epOrder)+"/pages?page="+strconv.Itoa(page), quality)
	if err != nil {
		return nil, err
	}
	var epPageResponse ComicPicturePageResponse
	err = json.Unmarshal(buff, &epPageResponse)
	if err != nil {
		return nil, err
	}
	return &epPageResponse.Data.Pages, nil
}

// SwitchLike (取消)收藏漫画
// 第一次收藏，第二次是取消收藏 action是最终结果
func (client *Client) SwitchLike(comicId string) (*string, error) {
	buff, err := client.postToPica("comics/"+comicId+"/like", nil)
	if err != nil {
		return nil, err
	}
	var actionResponse ActionResponse
	err = json.Unmarshal(buff, &actionResponse)
	if err != nil {
		return nil, err
	}
	return &actionResponse.Data.Action, nil
}

// SwitchFavourite (取消)收藏漫画
// 第一次收藏，第二次是取消收藏 action是最终结果
func (client *Client) SwitchFavourite(comicId string) (*string, error) {
	buff, err := client.postToPica("comics/"+comicId+"/favourite", nil)
	if err != nil {
		return nil, err
	}
	var actionResponse ActionResponse
	err = json.Unmarshal(buff, &actionResponse)
	if err != nil {
		return nil, err
	}
	return &actionResponse.Data.Action, nil
}

// FavouriteComics 收藏的漫画
func (client *Client) FavouriteComics(sort string, page int) (*ComicsPage, error) {
	buff, err := client.getToPica("users/favourite?s=" + sort + "&page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	var comicsResponse ComicsPageResponse
	err = json.Unmarshal(buff, &comicsResponse)
	if err != nil {
		return nil, err
	}
	return &comicsResponse.Data.Comics, nil
}

// ComicCommentsPage 漫画的评论
func (client *Client) ComicCommentsPage(comicId string, page int) (*CommentsPage, error) {
	buff, err := client.getToPica("comics/" + comicId + "/comments?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	buff = replaceStringPage(buff)
	var commentsResponse CommentsResponse
	err = json.Unmarshal(buff, &commentsResponse)
	if err != nil {
		return nil, err
	}
	return &commentsResponse.Data.Comments, nil
}

// ComicRecommendation 看了这个本子的也在看
func (client *Client) ComicRecommendation(comicId string) ([]ComicSimple, error) {
	buff, err := client.getToPica("comics/" + comicId + "/recommendation")
	if err != nil {
		return nil, err
	}
	var recommendationResponse ComicsResponse
	err = json.Unmarshal(buff, &recommendationResponse)
	if err != nil {
		return nil, err
	}
	return recommendationResponse.Data.Comics, nil
}

// HotKeywords 大家都在搜
func (client *Client) HotKeywords() ([]string, error) {
	buff, err := client.getToPica("keywords")
	if err != nil {
		return nil, err
	}
	var hotKeywordsResponse HotKeywordsResponse
	err = json.Unmarshal(buff, &hotKeywordsResponse)
	if err != nil {
		return nil, err
	}
	return hotKeywordsResponse.Data.Keywords, nil
}

// GamePage 游戏列表
func (client *Client) GamePage(page int) (*GamePage, error) {
	buff, err := client.getToPica("games?page=" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}
	var response GamePageResponse
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data.Games, nil
}

// GameInfo 游戏详情
func (client *Client) GameInfo(gameId string) (*GameInfo, error) {
	buff, err := client.getToPica("games/" + gameId)
	if err != nil {
		return nil, err
	}
	var response GameResponse
	err = json.Unmarshal(buff, &response)
	if err != nil {
		return nil, err
	}
	return &response.Data.Game, nil
}

//
var stringPageRegexp, _ = regexp.Compile("\"page\": \"(\\d+)\",")
var stringPageReplaceTo = []byte("\"page\": $1,")

func replaceStringPage(buff []byte) []byte {
	return stringPageRegexp.ReplaceAll(buff, stringPageReplaceTo)
}
