// Code generated by taoctl. DO NOT EDIT.
package types

type Post struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Meta struct {
	Total    int64 `json:"total"`
	Page     int64 `json:"page"`
	PageSize int64 `json:"page_size"`
}

type PostListReq struct {
	Page     int64 `form:"page,optional"`
	PageSize int64 `form:"pageSize,optional"`
}

type PostListResp struct {
	Post []*Post `json:"post"`
	Meta *Meta   `json:"meta"`
}

type PostDetailReq struct {
	Id int64 `path:"id"`
}

type PostDetailResp struct {
	Id        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	ReplyList []*Reply `json:"replies"`
}

type CreatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreatePostResp struct {
	Success bool `json:"success"`
}

type CreateReplyReq struct {
	PostID  int64  `path:"id"`
	Content string `form:"content"`
}

type CreateReplyResp struct {
	Success bool `json:"success"`
}

type GetReplyReq struct {
	PostID int64 `path:"id"`
}

type Reply struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
}

type GetReplyResp struct {
	List []*Reply `json:"list"`
}
