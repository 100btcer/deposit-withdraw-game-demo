package social

type RetweetRes struct {
	Data   []RetweetResData `json:"data"`
	Meta   RetweetResMeta   `json:"meta"`
	Status int              `json:"status"`
}

type RetweetResData struct {
	Username string `json:"username"`
}

type RetweetResMeta struct {
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}

type TokenRet struct {
	AppAccessToken string `json:"app_access_token"`
	Code           int    `json:"code"`
	Expire         int    `json:"expire"`
	Msg            string `json:"msg"`
}

type UserIdRet struct {
	UserId string `json:"user_id"`
}

type UserListRet struct {
	UserList []UserIdRet `json:"user_list"`
}

type UserRet struct {
	Data UserListRet `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

type SendRet struct {
	Error SendErr `json:"error"`
	Code  int     `json:"code"`
	Msg   string  `json:"msg"`
}

type SendErr struct {
	Message string `json:"message"`
	LogId   string `json:"log_id"`
}
