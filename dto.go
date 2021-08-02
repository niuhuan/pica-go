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

type ComicsResponse struct {
	Response
	Data struct {
		Comics ComicsPage `json:"comics"`
	} `json:"data"`
}

type ComicsPage struct {
	Page
	Docs []ComicSimple `json:"docs"`
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

type RecommendationResponse struct {
	Response
	Data struct {
		Comics []ComicSimple `json:"comics"`
	} `json:"data"`
}

type HotKeywordsResponse struct {
	Response
	Data struct {
		Keywords []string `json:"keywords"`
	} `json:"data"`
}
