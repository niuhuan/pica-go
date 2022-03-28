package pica

// 排序方式

type Sort string

const SortDefault Sort = "ua"
const SortTimeNewest Sort = "dd"
const SortTimeOldest Sort = "da"
const SortLikeMost Sort = "ld"
const SortViveMost Sort = "vd"

// 图片质量

type ImageQuality string

const ImageQualityOriginal ImageQuality = "original"
const ImageQualityLow ImageQuality = "low"
const ImageQualityMedium ImageQuality = "medium"
const ImageQualityHigh ImageQuality = "high"

// 一些请求结果

const ActionLike = "like"
const ActionUnlike = "unlike"

const ActionFavourite = "favourite"
const ActionUnFavourite = "un_favourite"

// 排行榜类型

const LeaderboardH24 = "H24"
const LeaderboardD7 = "D7"
const LeaderboardD30 = "D30"
