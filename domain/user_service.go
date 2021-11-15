package domain

type UserService interface {
	Login(req LoginReq) (LoginResp, error)
}
