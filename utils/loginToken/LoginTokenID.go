package loginToken

import (
	"encoding/json"
	"time"

	"github.com/cn-joyconn/goutils/encrypt"
	"github.com/cn-joyconn/goutils/strtool"
)

type LoginTokenID struct {
	Uid       string `json:"uid"`
	Pwd       string `json:"pwd"`
	Timestamp int64  `json:"timestamp"`
	Sign      string `json:"sign"`
	AesKey    string `json:"-"`
}

/**
*  由一个令牌字符串转换成令牌对象
* @param tokenStr 令牌字符串
* Created by Eric.Zhang on 2016/12/29.
 */
func ParseLoginTokenID(tokenStr string, token_ekey string) *LoginTokenID {
	if strtool.IsBlank(tokenStr) {
		return nil
	}
	//解析令牌
	bytes := encrypt.AesDecryptECB([]byte(tokenStr), []byte(token_ekey))
	if len(bytes) == 0 {
		return nil
	}
	var loginTokenID *LoginTokenID
	err := json.Unmarshal(bytes, &loginTokenID)
	if err != nil {
		return nil
	}
	loginTokenID.AesKey = token_ekey
	return loginTokenID
}

/**
*   由账号、密码生成一个令牌
* @param userid 账号
* @param password 密码
* Created by Eric.Zhang on 2016/12/29.
 */
func CreateLoginTokenID(sign string,userid string, password string, token_ekey string) *LoginTokenID {
	
	loginTokenID := &LoginTokenID{
		Uid:       userid,
		Pwd:       password,
		Timestamp: time.Now().Unix(),
		Sign:      sign,
		AesKey:    token_ekey,
	}
	return loginTokenID
}

func (l *LoginTokenID) toString() string {
	bytes, err := json.Marshal(&l)
	if err != nil {
		return ""
	}

	bytes = encrypt.AesEncryptECB(bytes, []byte(l.AesKey))

	return string(bytes)

}
