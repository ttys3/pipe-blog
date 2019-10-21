// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-present, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package controller

import (
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"fmt"
	"net/http"

	"github.com/b3log/gulu"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userName = "admin"
const DftAvatarURl = "https://avatars0.githubusercontent.com/u/41882455?s=40&v=4"

// logoutAction logout a user.
func logoutAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	session := sessions.Default(c)
	session.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	session.Clear()
	if err := session.Save(); nil != err {
		logger.Errorf("saves session failed: " + err.Error())
	}
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)
	result.Code = util.CodeErr

	// init the admin user
	if !service.Init.Inited() {
		user := &model.User{
			Name:      "admin",
			AvatarURL: DftAvatarURl,
			B3Key:     "",
			GithubId:  "",
		}
		if err := user.UpdatePasswd("admin"); err != nil {
			logger.Errorf("init admin user failed: " + err.Error())
			result.Msg = "init admin user failed: " + err.Error()
			return
		}

		if err := service.Init.InitPlatform(user); nil != err {
			logger.Errorf("init platform via github login failed: " + err.Error())
			result.Msg = "init platform via github login failed: " + err.Error()
			return
		}
	}

	var req loginReq
	if err :=c.BindJSON(&req); err != nil {
		logger.Errorf("invalid login request: " + err.Error())
		result.Msg = "invalid login request: " + err.Error()
		return
	}

	user := service.User.GetUserByName(req.Username)
	if user == nil {
		logger.Errorf("no user found: " + req.Username)
		result.Msg = "user or password validation failed"
		return
	}

	ownBlog := service.User.GetOwnBlog(user.ID)
	if nil == ownBlog {
		logger.Warnf("can not get blog for user [" + userName + "]")
		result.Msg = "can not get blog for user [" + userName + "]"
		return
	}

	if err := user.VerifyPasswd(req.Password); err != nil {
		logger.Warnf("user password verify errror: %s", err)
		result.Msg = "user or password validation failed"
		return
	}

	session := &util.SessionData{
		UID:     user.ID,
		UName:   user.Name,
		UNickname: user.Nickname,
		UB3Key:  user.B3Key,
		UAvatar: user.AvatarURL,
		URole:   ownBlog.UserRole,
		BID:     ownBlog.ID,
		BURL:    ownBlog.URL,
	}
	if err := session.Save(c); nil != err {
		logger.Errorf("saves session failed: " + err.Error())
		result.Msg = fmt.Sprintf("saves session failed: %s", err.Error())
	}
	result.Code = util.CodeOk
	result.Msg = "login success"
}
