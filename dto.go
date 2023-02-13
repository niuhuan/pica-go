package pica

import "time"

// Response 返回体格式
type Response[T any] struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
	Data    T      `json:"data"`
}

// PageData 分页格式
type PageData struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

// Image 图片资源
type Image struct {
	OriginalName string `json:"originalName"`
	Path         string `json:"path"`
	FileServer   string `json:"fileServer"`
}

// RegisterDto 注册接口请求体
type RegisterDto struct {
	Email     string `json:"email"`    // 邮箱
	Password  string `json:"password"` // 8字以上
	Name      string `json:"name"`     // 2 - 50 字
	Birthday  string `json:"birthday"` // 2012-01-01
	Gender    string `json:"gender"`   // m, f, bot
	Answer1   string `json:"answer1"`
	Answer2   string `json:"answer2"`
	Answer3   string `json:"answer3"`
	Question1 string `json:"question1"`
	Question2 string `json:"question2"`
	Question3 string `json:"question3"`
}

// UserBasic 用户的基本信息
type UserBasic struct {
	Id         string   `json:"_id"`
	Gender     string   `json:"gender"`
	Name       string   `json:"name"`
	Title      string   `json:"title"`
	Verified   bool     `json:"verified"`
	Exp        int      `json:"exp"`
	Level      int      `json:"level"`
	Characters []string `json:"characters"`
	Avatar     Image    `json:"avatar"`
	Slogan     string   `json:"slogan"` // 有可能是null, 从未设置过slogan的人
}

// LoginRequest 登录的请求体 (PS:Email字段为账号, 并不一定是邮箱格式)
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	Token string `json:"token"`
}

// UserProfileResult 获取个人信息接口返回内容
type UserProfileResult struct {
	User UserProfile `json:"user"`
}

// UserProfile 获取个人信息接口返回内容 | 个人信息
type UserProfile struct {
	UserBasic
	Birthday  string    `json:"birthday"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsPunched bool      `json:"isPunched"` // 是否打了哔咔
}

// PunchResult 打哔咔接口返回内容
type PunchResult struct {
	Res PunchStatus `json:"res"`
}

// PunchStatus 打哔咔接口返回内容
type PunchStatus struct {
	Status         string `json:"status"`
	PunchInLastDay string `json:"punchInLastDay"`
}

// CategoriesResult 获取分类接口返回内容
type CategoriesResult struct {
	Categories []Category `json:"categories"`
}

// Category 分类
type Category struct {
	Id          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumb       Image  `json:"thumb"`
	IsWeb       bool   `json:"isWeb"`
	Active      bool   `json:"active"`
	Link        string `json:"link"`
}

// ComicsPageResult 漫画列表接口返回内容
type ComicsPageResult struct {
	Comics ComicsPage `json:"comics"`
}

// ComicsPage 漫画的分页
type ComicsPage struct {
	PageData
	Docs []ComicSimple `json:"docs"`
}

// ComicsResult 漫画列表返回内容 用于随机漫画, 排行榜等不分页的接口
type ComicsResult struct {
	Comics []ComicSimple `json:"comics"`
}

// ComicSimple 漫画摘要内容, 列表页面使用
type ComicSimple struct {
	Id         string   `json:"_id"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	PagesCount int      `json:"pagesCount"`
	EpsCount   int      `json:"epsCount"`
	Finished   bool     `json:"finished"`
	Categories []string `json:"categories"`
	Thumb      Image    `json:"thumb"`
	LikesCount int      `json:"likesCount"`
}

// ComicInfoResult 获取漫画详情接口返回内容
type ComicInfoResult struct {
	Comic ComicInfo `json:"comic"`
}

// ComicInfo 漫画详情
type ComicInfo struct {
	ComicSimple
	Creator       Creator   `json:"_creator"`
	Description   string    `json:"description"`
	ChineseTeam   string    `json:"chineseTeam"`
	Tags          []string  `json:"tags"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	AllowDownload bool      `json:"allowDownload"`
	ViewsCount    int       `json:"viewsCount"`
	IsFavourite   bool      `json:"isFavourite"`
	IsLiked       bool      `json:"isLiked"`
	CommentsCount int       `json:"commentsCount"`
}

// Creator 漫画的创建人
type Creator struct {
	UserBasic
	Role      string `json:"role"`
	Character string `json:"character"`
}

// EpPageResult 获取漫画章节列表接口返回内容
type EpPageResult struct {
	Eps EpPage `json:"eps"`
}

// EpPage 漫画的章节的分页
type EpPage struct {
	PageData
	Docs []Ep `json:"docs"`
}

// Ep 漫画的章节
type Ep struct {
	Id        string    `json:"_id"`
	Title     string    `json:"title"`
	Order     int       `json:"order"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ComicPicturePageResult 获取章节图片接口返回内容
type ComicPicturePageResult struct {
	Pages ComicPicturePage `json:"pages"`
	Ep    Ep               `json:"ep"`
}

// ComicPicturePage 章节图片的分页
type ComicPicturePage struct {
	PageData
	Docs []ComicPicture `json:"docs"`
}

// ComicPicture 章节的图片
type ComicPicture struct {
	Media Image  `json:"media"`
	Id    string `json:"_id"`
}

// ActionResult 点赞,收藏 等接口返回内容
type ActionResult struct {
	Action string `json:"action"`
}

// CommentBase 评论
type CommentBase struct {
	Id            string      `json:"_id"`
	Content       string      `json:"content"`
	User          CommentUser `json:"_user"`
	IsTop         bool        `json:"isTop"`
	Hide          bool        `json:"hide"`
	CreatedAt     time.Time   `json:"created_at"`
	LikesCount    int         `json:"likesCount"`
	CommentsCount int         `json:"commentsCount"`
	IsLiked       bool        `json:"isLiked"`
}

// ChildOfComment 自评的字段
type ChildOfComment struct {
	Parent string `json:"_parent"`
}

// CommentUser 发出此评论的人
type CommentUser struct {
	UserBasic
	Role string `json:"role"`
}

// CommentsResult 获取漫画评论接口返回内容
type CommentsResult struct {
	Comments    CommentsPage `json:"comments"`
	TopComments []Comment    `json:"topComments"`
}

// CommentsPage 漫画评论的分页
type CommentsPage struct {
	PageData
	Docs []Comment `json:"docs"`
}

// Comment 漫画的评论
type Comment struct {
	CommentBase
	Comic string `json:"_comic"`
}

// CommentChildrenResult 获取子评论接口返回内容
type CommentChildrenResult struct {
	Comments CommentChildrenPage `json:"comments"`
}

// CommentChildrenPage 子评论分页
type CommentChildrenPage struct {
	PageData
	Docs []CommentChild `json:"docs"`
}

// CommentChild 子评论
type CommentChild struct {
	Comment
	ChildOfComment
}

// MyCommentsPageResult 我的评论接口返回内容
type MyCommentsPageResult struct {
	Comments MyCommentsPage `json:"comments"`
}

// MyCommentsPage 我的评论分页
type MyCommentsPage struct {
	PageData
	Docs []MyComment `json:"docs"`
}

// MyComment 我的评论
type MyComment struct {
	Id      string `json:"_id"`
	Content string `json:"content"`
	Comic   struct {
		Id    string `json:"_id"`
		Title string `json:"title"`
	} `json:"_comic"`
	Hide          bool      `json:"hide"`
	CreatedAt     time.Time `json:"created_at"`
	LikesCount    int       `json:"likesCount"`
	CommentsCount int       `json:"commentsCount"`
	IsLiked       bool      `json:"isLiked"`
}

// LeaderboardOfKnightResult 骑士榜接口返回内容
type LeaderboardOfKnightResult struct {
	Users []Knight `json:"users"`
}

// Knight 用户(骑士榜)
type Knight struct {
	UserBasic
	Role           string `json:"role"`
	Character      string `json:"character"`
	ComicsUploaded int    `json:"comicsUploaded"`
}

// HotKeywordsResutl 大家搜在搜接口返回内容
type HotKeywordsResutl struct {
	Keywords []string `json:"keywords"`
}

// GamePageResult  游戏列表接口返回内容
type GamePageResult struct {
	Games GamePage `json:"games"`
}

// GamePage 游戏列表
type GamePage struct {
	PageData
	Docs []GameSimple `json:"docs"`
}

// GameSimple 游戏摘要
type GameSimple struct {
	Id         string `json:"_id"`
	Title      string `json:"title"`
	Version    string `json:"version"`
	Icon       Image  `json:"icon"`
	Publisher  string `json:"publisher"`
	Adult      bool   `json:"adult"`
	Suggest    bool   `json:"suggest"`
	LikesCount int    `json:"likesCount"`
	Android    bool   `json:"android"`
	Ios        bool   `json:"ios"`
}

// GameResult 游戏详情接口返回内容
type GameResult struct {
	Game GameInfo `json:"game"`
}

// GameInfo 游戏详情
type GameInfo struct {
	GameSimple
	Description    string    `json:"description"`
	UpdateContent  string    `json:"updateContent"`
	VideoLink      string    `json:"videoLink"`
	Screenshots    []Image   `json:"screenshots"`
	CommentsCount  int       `json:"commentsCount"`
	DownloadsCount int       `json:"downloadsCount"`
	IsLiked        bool      `json:"isLiked"`
	AndroidLinks   []string  `json:"androidLinks"`
	AndroidSize    float32   `json:"androidSize"`
	IosLinks       []string  `json:"iosLinks"`
	IosSize        float32   `json:"iosSize"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// GameCommentsResult 获取漫画评论接口返回内容
type GameCommentsResult struct {
	Comments    GameCommentsPage `json:"comments"`
	TopComments []Comment        `json:"topComments"`
}

// GameCommentsPage 游戏评论的分页
type GameCommentsPage struct {
	PageData
	Docs []GameComment `json:"docs"`
}

// GameComment 游戏的评论
type GameComment struct {
	CommentBase
	Game string `json:"_game"`
}

// GameCommentChildrenResult 获取游戏子评论接口返回内容
type GameCommentChildrenResult struct {
	Comments GameCommentChildrenPage `json:"comments"`
}

// GameCommentChildrenPage 游戏子评论分页
type GameCommentChildrenPage struct {
	PageData
	Docs []GameCommentChild `json:"docs"`
}

// GameCommentChild 游戏子评论
type GameCommentChild struct {
	GameComment
	ChildOfComment
}

// CollectionsResult 合集返回体
type CollectionsResult struct {
	Collections []Collection `json:"collections"`
}

// Collection 合集
type Collection struct {
	Title  string        `json:"title"`
	Comics []ComicSimple `json:"comics"`
}

// ForgotPasswordResult 找回密码-获取问题-返回DATA
type ForgotPasswordResult struct {
	Question1 string `json:"question1"`
	Question2 string `json:"question2"`
	Question3 string `json:"question3"`
}

// ForgotPasswordResponse 找回密码-获取问题-返回
type ForgotPasswordResponse struct {
	Response
	Data ForgotPasswordResult `json:"data"`
}

// ResetPasswordResult 找回密码-根据答案重置密码
type ResetPasswordResult struct {
	Password string `json:"password"`
}

// ResetPasswordResponse 找回密码-获取问题-返回
type ResetPasswordResponse struct {
	Response
	Data ResetPasswordResult `json:"data"`
}

type InitInfo struct {
	Status    string   `json:"status"`
	Addresses []string `json:"addresses"`
	Waka      string   `json:"waka"`
	AdKeyword string   `json:"adKeyword"`
}
