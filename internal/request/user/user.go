package user

type IndexReq struct {
	Name string `form:"name" query:"name" json:"name" xml:"name" validate:"required"`
	Id   int    `form:"id" query:"id" xml:"id" validate:"required"`
}

type LoginReq struct {
	Account  string `form:"account" json:"account" validate:"required"`
	Password string `form:"password" json:"password" validate:"required"`
}
