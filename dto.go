package pica

import "time"

// Response 返回体格式
type Response struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
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
}

// LoginRequest 登录的请求体 (PS:Email字段为账号, 并不一定是邮箱格式)
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse 登录的返回体
type LoginResponse struct {
	Response
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

// UserProfileResponse 获取个人信息接口返回体
type UserProfileResponse struct {
	Data struct {
		User UserProfile `json:"user"`
	} `json:"data"`
}

// UserProfile 获取个人信息接口返回内容 | 个人信息
type UserProfile struct {
	UserBasic
	Birthday  string    `json:"birthday"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsPunched bool      `json:"isPunched"` // 是否打了哔咔
}

// PunchResponse 打哔咔接口返回体
type PunchResponse struct {
	Response
	Data struct {
		Res PunchStatus `json:"res"`
	} `json:"data"`
}

// PunchStatus 打哔咔接口返回内容
type PunchStatus struct {
	Status         string `json:"status"`
	PunchInLastDay string `json:"punchInLastDay"`
}

// CategoriesResponse 获取分类接口返回体
type CategoriesResponse struct {
	Response
	Data struct {
		Categories []Category `json:"categories"`
	} `json:"data"`
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

// ComicsPageResponse 漫画列表接口返回体
type ComicsPageResponse struct {
	Response
	Data struct {
		Comics ComicsPage `json:"comics"`
	} `json:"data"`
}

// ComicsPage 漫画的分页
type ComicsPage struct {
	PageData
	Docs []ComicSimple `json:"docs"`
}

// ComicsResponse 漫画列表返回体 用于随机漫画, 排行榜等不分页的接口
type ComicsResponse struct {
	Response
	Data struct {
		Comics []ComicSimple `json:"comics"`
	} `json:"data"`
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

// ComicInfoResponse 获取漫画详情接口返回体
type ComicInfoResponse struct {
	Response
	Data struct {
		Comic ComicInfo `json:"comic"`
	} `json:"data" `
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
	Slogan    string `json:"slogan"`
	Role      string `json:"role"`
	Character string `json:"character"`
}

// EpPageResponse 获取漫画章节列表接口返回体
type EpPageResponse struct {
	Response
	Data struct {
		Eps EpPage `json:"eps"`
	} `json:"data"`
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

// ComicPicturePageResponse 获取章节图片接口返回体
type ComicPicturePageResponse struct {
	Response
	Data struct {
		Pages ComicPicturePage `json:"pages"`
		Ep    Ep               `json:"ep"`
	} `json:"data"`
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

// ActionResponse 点赞,收藏 等接口返回体
type ActionResponse struct {
	Data struct {
		Action string `json:"action"`
	} `json:"data"`
}

// CommentsResponse 获取漫画评论接口返回体
type CommentsResponse struct {
	Response
	Data struct {
		Comments    CommentsPage `json:"comments"`
		TopComments []Comment    `json:"topComments"`
	} `json:"data"`
}

// CommentsPage 漫画评论的分页
type CommentsPage struct {
	PageData
	Docs []Comment `json:"docs"`
}

// Comment 漫画的评论
type Comment struct {
	Id            string      `json:"_id"`
	Content       string      `json:"content"`
	User          CommentUser `json:"_user"`
	Comic         string      `json:"_comic"`
	IsTop         bool        `json:"isTop"`
	Hide          bool        `json:"hide"`
	CreatedAt     time.Time   `json:"created_at"`
	LikesCount    int         `json:"likesCount"`
	CommentsCount int         `json:"commentsCount"`
	IsLiked       bool        `json:"isLiked"`
}

// CommentUser 发出此评论的人
type CommentUser struct {
	UserBasic
	Role string `json:"role"`
}

// CommentChildrenResponse 获取子评论接口返回体
type CommentChildrenResponse struct {
	Response
	Data struct {
		Comments CommentChildrenPage `json:"comments"`
	} `json:"data"`
}

// CommentChildrenPage 子评论分页
type CommentChildrenPage struct {
	PageData
	Docs []CommentChild `json:"docs"`
}

// CommentChild 子评论
type CommentChild struct {
	Comment
	Parent string `json:"_parent"`
}

// MyCommentsPageResponse 我的评论接口返回体
type MyCommentsPageResponse struct {
	Response
	Data struct {
		Comments MyCommentsPage `json:"comments"`
	} `json:"data"`
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

// HotKeywordsResponse 大家搜在搜接口返回体
type HotKeywordsResponse struct {
	Response
	Data struct {
		Keywords []string `json:"keywords"`
	} `json:"data"`
}

// GamePageResponse 游戏列表接口返回体
type GamePageResponse struct {
	Response
	Data struct {
		Games GamePage `json:"games"`
	} `json:"data"`
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

// GameResponse 游戏详情接口返回体
type GameResponse struct {
	Response
	Data struct {
		Game GameInfo `json:"game"`
	} `json:"data"`
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

// LeaderboardOfKnightResponse 骑士榜接口返回体
type LeaderboardOfKnightResponse struct {
	Response
	Data struct {
		Users []Knight `json:"users"`
	} `json:"data"`
}

// Knight 用户(骑士榜)
type Knight struct {
	UserBasic
	Role           string `json:"role"`
	Character      string `json:"character"`
	ComicsUploaded int    `json:"comicsUploaded"`
}
