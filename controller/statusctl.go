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
	"net/http"
	"strings"

	"github.com/b3log/gulu"
	"github.com/b3log/pipe/model"
	"github.com/b3log/pipe/service"
	"github.com/b3log/pipe/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Status represents platform status and blog info.
type Status struct {
	*service.PlatformStatus

	Name      string              `json:"name"`
	Nickname  string              `json:"nickname"`
	AvatarURL string              `json:"avatarURL"`
	BlogTitle string              `json:"blogTitle"`
	BlogURL   string              `json:"blogURL"`
	Role      int                 `json:"role"`
	Blogs     []*service.UserBlog `json:"blogs"`
}

func getStatusAction(c *gin.Context) {
	result := gulu.Ret.NewResult()
	defer c.JSON(http.StatusOK, result)

	platformStatus, err := service.Init.Status()
	if nil != err {
		result.Code = util.CodeErr
		result.Msg = err.Error()

		return
	}

	data := &Status{
		PlatformStatus: platformStatus,
	}

	session := util.GetSession(c)
	if 0 != session.UID {
		user := service.User.GetUser(session.UID)
		if nil == user {
			session := sessions.Default(c)
			session.Options(sessions.Options{
				Path:   "/",
				MaxAge: -1,
			})
			session.Clear()
			if err := session.Save(); nil != err {
				logger.Errorf("saves session failed: " + err.Error())
			}

			return
		}

		data.Name = user.Name
		data.Nickname = user.Nickname
		data.AvatarURL = user.AvatarURL
		data.Role = model.UserRoleBlogAdmin

		if model.UserRoleNoLogin != session.URole && platformStatus.Inited {
			ownBlog := service.User.GetOwnBlog(user.ID)
			if nil != ownBlog {
				data.BlogTitle = ownBlog.Title
				data.BlogURL = ownBlog.URL
				// fixup SSL
				blogUrlTemp := ownBlog.URL
				if model.Conf.SSL && strings.HasPrefix(blogUrlTemp, "http://") {
					data.BlogURL = strings.Replace(blogUrlTemp, "http://", "https://", 1)
				} else if !model.Conf.SSL && strings.HasPrefix(blogUrlTemp, "https://") {
					data.BlogURL = strings.Replace(blogUrlTemp, "https://", "http://", 1)
				}
			}
			data.Blogs = service.User.GetUserBlogs(user.ID)
		}
	}

	result.Data = data
}
