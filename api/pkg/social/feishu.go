package social

import (
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"io/ioutil"
	"mm-ndj/config"
	"net/http"
)

var FeiToken = ""

/**
1.https://open.feishu.cn/app?lang=zh-CN
2.点击"创建企业自建应用"
3.弹窗中输入：名称"oa";应用描述：“oa”
4.创建后进入应用详情：https://open.feishu.cn/app/cli_xxx/baseinfo
5.复制App ID,App Secret
6.点击菜单栏：“添加应用能力”，页面中，在机器人一栏中点击“添加能力”
7.点击菜单栏：“权限管理”，添加以下权限：contact:contact:readonly_as_app，contact:user.employee_id:readonly，contact:user.id:readonly，im:message:send_as_bot
8.添加后，点击左上角“创建版本”；创建后进行发布，若需要审核，进入管理后台：https://cqwummbb4yt.feishu.cn/admin/index【右上角菜单栏，点击“管理后台”】
9.管理后台中，点击菜单栏：工作台->应用审核，进行审核；
10.审核通过后，使用App ID,App Secret填入代码中todo的地方，调用SendFeiShuMsg发送消息
*/

func SendFeiShuMsg(mobile, msgStr string) (err error) {
	//获取用户编号
	userId, err := GetFeiShuUserId(mobile)
	if err != nil {
		return
	}

	rawurl := "https://open.feishu.cn/open-apis/message/v4/send/"

	msg := `"post": {"zh_cn": { "title": "有新的任务等待处理","content": [[{"tag": "text","un_escape": true,"text": "` + msgStr + `"}]]}}`
	data := make(map[string]interface{})
	data["open_id"] = userId
	data["msg_type"] = "post"
	data["content"] = msg
	reqBody, _ := json.Marshal(data)

	token, err := GetFeiShuToken()
	if err != nil {
		return
	}

	c := &http.Client{}
	req, _ := http.NewRequest("post", rawurl, bytes.NewReader(reqBody))

	token = "Bearer " + token
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	response, err := c.Do(req)
	if err != nil {
		config.Logger.Error("GetFeiShuUserId err ", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	ret := SendRet{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		config.Logger.Error("SendFeiShuMsg err ", zap.Error(err))
		return
	}
	if ret.Code != 0 {
		config.Logger.Error("SendFeiShuMsg err ", zap.Error(err))
		return
	}
	return
}

func GetFeiShuToken() (token string, err error) {

	if FeiToken != "" {
		token = FeiToken
		return
	}

	rawurl := "https://open.feishu.cn/open-apis/auth/v3/app_access_token/internal"

	data := make(map[string]interface{})
	data["app_id"] = ""     // todo
	data["app_secret"] = "" // todo
	reqBody, _ := json.Marshal(data)

	c := &http.Client{}
	req, _ := http.NewRequest("post", rawurl, bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	response, err := c.Do(req)
	if err != nil {
		config.Logger.Error("GetFeishuToken err ", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		config.Logger.Error("GetFeishuToken err ", zap.Error(err))
		return
	}

	ret := TokenRet{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		config.Logger.Error("GetFeishuToken err ", zap.Error(err))
		return
	}
	if ret.Code != 0 {
		config.Logger.Error("GetFeishuToken err ", zap.Any("msg", ret.Msg))
		return
	}
	token = ret.AppAccessToken
	FeiToken = token
	return
}

func GetFeiShuUserId(mobile string) (userId string, err error) {
	rawurl := "https://open.feishu.cn/open-apis/contact/v3/users/batch_get_id"

	data := make(map[string]interface{})
	data["mobiles"] = []string{mobile}
	reqBody, _ := json.Marshal(data)

	c := &http.Client{}
	req, _ := http.NewRequest("post", rawurl, bytes.NewReader(reqBody))
	token, err := GetFeiShuToken()
	if err != nil {
		return
	}
	token = "Bearer " + token
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	response, err := c.Do(req)
	if err != nil {
		config.Logger.Error("GetFeishuToken err ", zap.Error(err))
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	ret := UserRet{}
	err = json.Unmarshal(body, &ret)
	if err != nil {
		config.Logger.Error("GetFeiShuUserId err ", zap.Error(err))
		return
	}
	if ret.Code != 0 {
		config.Logger.Error("GetFeiShuUserId err ", zap.Error(err))
		return
	}
	l := ret.Data.UserList
	if len(l) > 0 {
		userId = l[0].UserId
		return
	}
	err = errors.New("empty user id")
	return
}
