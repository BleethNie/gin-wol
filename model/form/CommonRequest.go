package req

type KeywordQuery struct {
	Keyword string `form:"keyword"`
}

type PageQuery struct {
	PageSize int    `form:"page_size"`
	PageNum  int    `form:"page_num"`
	Keyword  string `form:"keyword"`
}

type SoftDelete struct {
	Ids      []int `json:"ids"`
	IsDelete *int8 `json:"is_delete" validate:"required,min=0,max=1"` // 软删除到回收站, 没有的字段不使用
}

type UpdateReview struct {
	Ids      []int `json:"ids"`
	IsReview *int8 `json:"is_review" validate:"required,min=0,max=1"`
}
