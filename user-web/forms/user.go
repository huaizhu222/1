package forms

type PassWordLoginForm struct {
	User_Id  int32  `form:"user_id" json:"user_id" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

type RegisterForm struct {
	User_Id   int32  `form:"user_id" json:"user_id" binding:"required"`
	PassWord  string `form:"password" json:"password" binding:"required"`
	Nick_name string `form:"Nick_name" json:"Nick_name" binding:"required"`
	Like      string `form:"like" json:"like" binding:"required"`
}
