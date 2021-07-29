package models

//CreatePostModel ...
type CreatePostModel struct {
	Title	string	`json:"title"`
	Body	string	`json:"body"`
	Author	string	`json:"author"`
}

//UpdatePostModel ...
type UpdatePostModel struct {
	Title	string `json:"title"`
	Body	string `json:"body"`
	Author	string `json:"author"`
}

//GetPostModel ...
type GetPostModel struct {
	Id			uint32	`json:"id"`
	Title		string	`json:"title"`
	Body		string	`json:"body"`
	Author		string	`json:"author"`
	CreatedAt	string	`json:"created_at"`
}

//GetAllPostsModel ...
type GetAllPostsModel struct {
	Count		int				`json:"count"`
	Customers	[]GetPostModel	`json:"posts"`
}


