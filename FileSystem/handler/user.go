package handler

import (
	dblayer "FileSystem/db"
	utils "FileSystem/utils"
	"fmt"
	"net/http"
	"time"
)

const (
	pwd_salt = "*#890"
)

// SignupHandler : 处理用户注册请求
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	//get方法处理，将页面返回
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signup.html", http.StatusFound)
		return
	}
	//解析参数
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	//合法性校验
	if len(username) < 3 || len(password) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}
	enc_passwd := utils.Sha1([]byte(password + pwd_salt))
	suc := dblayer.UserSignup(username, enc_passwd)
	if suc {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("FAILED"))
	}
}

// SigninHandler : 登陆接口
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/static/view/signin.html", http.StatusFound)
		return
	}
	//解析参数
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password") //此处获得的是加密后的密码
	encPasswd := utils.Sha1([]byte(password + pwd_salt))
	//校验用户名和密码
	pwdChecked := dblayer.UserSignin(username, encPasswd)
	if !pwdChecked {
		w.Write([]byte(`{"code": -1, "msg": "用户名或密码错误"}`))
		return
	}
	//生成访问凭证
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte(`{"code": -1, "msg": "Token更新失败"}`))
		return
	}
	//登陆成功后重定向到首页
	resp := utils.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Token    string
			Username string
			Location string
		}{
			Token:    token,
			Username: username,
			Location: "http://" + r.Host + "/static/view/home.html",
		},
	}
	w.Write(resp.JSONBytes())
}

// UserInfoHandler : 查询用户信息
func UserInfoHandler(w http.ResponseWriter, r *http.Request) {
	//解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	//token := r.Form.Get("token")
	//验证token是否有效
	//isValidToken := isTokenValid(token)
	//if !isValidToken {
	//	w.WriteHeader(http.StatusForbidden)
	//	return
	//}
	//查询用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//组装并响应用户数据
	resp := utils.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

// GenToken : 生成令牌
func GenToken(username string) string {
	//40位字符md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := utils.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// isTokenValid : 验证token是否有效
func isTokenValid(token string) bool {
	if len(token) < 40 {
		return false
	}
	// TODO:判断token的时效性，是否过期

	// TODO:从数据库表tbl_user_token查询username对应的token信息

	// TODO:对比两个token是否一致
	return true

}
