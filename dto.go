package pica

import "time"

type Response struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

type Page struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

type Image struct {
	OriginalName string `json:"originalName"`
	Path         string `json:"path"`
	FileServer   string `json:"fileServer"`
}

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

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Response
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type UserProfileResponse struct {
	Data struct {
		User UserProfile `json:"user"`
	} `json:"data"`
}

type UserProfile struct {
	UserBasic
	Birthday  string    `json:"birthday"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	IsPunched bool      `json:"isPunched"`
}

type PunchResponse struct {
	Response
	Data struct {
		Res PunchStatus `json:"res"`
	} `json:"data"`
}

type PunchStatus struct {
	Status         string `json:"status"`
	PunchInLastDay string `json:"punchInLastDay"`
}

type CategoriesResponse struct {
	Response
	Data struct {
		Categories []Category `json:"categories"`
	} `json:"data"`
}

type Category struct {
	Id          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumb       Image  `json:"thumb"`
	IsWeb       bool   `json:"isWeb"`
	Active      bool   `json:"active"`
	Link        string `json:"link"`
}

type ComicsPageResponse struct {
	Response
	Data struct {
		Comics ComicsPage `json:"comics"`
	} `json:"data"`
}

type ComicsPage struct {
	Page
	Docs []ComicSimple `json:"docs"`
}

type ComicsResponse struct {
	Response
	Data struct {
		Comics []ComicSimple `json:"comics"`
	} `json:"data"`
}

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

type ComicInfoResponse struct {
	Response
	Data struct {
		Comic ComicInfo `json:"comic"`
	} `json:"data" `
}

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

type Creator struct {
	UserBasic
	Slogan    string `json:"slogan"`
	Role      string `json:"role"`
	Character string `json:"character"`
}

type EpPageResponse struct {
	Response
	Data struct {
		Eps EpPage `json:"eps"`
	} `json:"data"`
}

type EpPage struct {
	Page
	Docs []Ep `json:"docs"`
}

type Ep struct {
	Id        string    `json:"_id"`
	Title     string    `json:"title"`
	Order     int       `json:"order"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ComicPicturePageResponse struct {
	Response
	Data struct {
		Pages ComicPicturePage `json:"pages"`
		Ep    Ep               `json:"ep"`
	} `json:"data"`
}

type ComicPicturePage struct {
	Page
	Docs []ComicPicture `json:"docs"`
}

type ComicPicture struct {
	Media Image  `json:"media"`
	Id    string `json:"_id"`
}

type ActionResponse struct {
	Data struct {
		Action string `json:"action"`
	} `json:"data"`
}

type CommentsResponse struct {
	Response
	Data struct {
		Comments    CommentsPage `json:"comments"`
		TopComments []Comment    `json:"topComments"`
	} `json:"data"`
}

type CommentsPage struct {
	Page
	Docs []Comment `json:"docs"`
}

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

type CommentUser struct {
	UserBasic
	Role string `json:"role"`
}

type CommentChildrenResponse struct {
	Response
	Data struct {
		Comments    CommentChildrenPage `json:"comments"`
	} `json:"data"`
}

type CommentChildrenPage struct {
	Page
	Docs []CommentChild `json:"docs"`
}

type CommentChild struct {
	Comment
	Parent string `json:"_parent"`
}

type MyCommentsPageResponse struct {
	Response
	Data struct{
		Comments MyCommentsPage `json:"comments"`
	} `json:"data"`
}

type MyCommentsPage struct {
	Page
	Docs []MyComment `json:"docs"`
}

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

type HotKeywordsResponse struct {
	Response
	Data struct {
		Keywords []string `json:"keywords"`
	} `json:"data"`
}

type GamePageResponse struct {
	Response
	Data struct {
		Games GamePage `json:"games"`
	} `json:"data"`
}

type GamePage struct {
	Page
	Docs []GameSimple `json:"docs"`
}

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

type GameResponse struct {
	Response
	Data struct {
		Game GameInfo `json:"game"`
	} `json:"data"`
}

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
