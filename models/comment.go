package models

type Comment struct {
	ID            int     `json:"id"`
	ArticleID     int     `json:"article_id"`
	Article       Article `json:"article"`
	UserID        int     `json:"user_id"`
	User          User    `json:"user"`
	ReplyToUserID int     `json:"reply_to_user_id"`
	ReplyToUser   User    `json:"replyToUser"`
	Content       string  `json:"content"`
	Status        int     `json:"status"`
	ParentID      int     `json:"parentId"`
	CreatedAt     MyTime  `json:"createdAt"`
	UpdatedAt     MyTime  `json:"updatedAt"`
}

type CommentVo struct {
	ID            int       `json:"id"`
	UserID        int       `json:"user_id"`
	User          User      `json:"user"`
	ReplyToUserID int       `json:"reply_to_user_id"`
	ReplyToUser   User      `json:"replyToUser"`
	Content       string    `json:"content"`
	Status        int       `json:"status"`
	ParentID      int       `json:"parent_id"`
	SubComments   []Comment `json:"subComments" gorm:"ForeignKey:parent_id;"`
	CreatedAt     MyTime    `json:"createdAt"`
	UpdatedAt     MyTime    `json:"updatedAt"`
}

type CommentReq struct {
	UserId      int     `json:"user_id"`
	Article     Article `json:"article" `
	Content     string  `json:"content" `
	ReplyToUser User    `json:"replyToUser"`
	ParentId    int     `json:"parentId" `
}

// type SubComment struct {
// 	ID            int    `json:"id"`
// 	UserId        int    `json:"userid"`
// 	User          User   `json:"user" gorm:"-"`
// 	ReplyToUserId int    `json:"replytouserid" gorm:"-"`
// 	ReplyToUser   User   `json:"replyToUser" gorm:"-"`
// 	Content       string `json:"content"`
// 	Status        int    `json:"status"`
// 	CreatedAt     string `json:"createdAt"`
// 	ParentId      int    `json:"parentId"`
// }
