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

package model

import (
	"github.com/b3log/pipe/util"
	"golang.org/x/crypto/bcrypt"
)

// User model.
type User struct {
	Model

	Name              string `gorm:"size:32" json:"name"`
	PasswdHash		  string `gorm:"size:64" json:"-"`
	Nickname          string `gorm:"size:32" json:"nickname"`
	AvatarURL         string `gorm:"size:255" json:"avatarURL"`
	B3Key             string `gorm:"size:32" json:"b3Key"`
	Locale            string `gorm:"size:32" json:"locale"`
	TotalArticleCount int    `json:"totalArticleCount"`
	GithubId          string `gorm:"255" json:"githubId"`
}

// User roles.
const (
	UserRoleNoLogin = iota
	UserRolePlatformAdmin
	UserRoleBlogAdmin
	UserRoleBlogUser
)

// AvatarURLWithSize returns avatar URL with the specified size.
func (u *User) AvatarURLWithSize(size int) string {
	return util.ImageSize(u.AvatarURL, size, size)
}

func (u *User) UpdatePasswd(newPwd string) error {
	if h, e := PasswdHash([]byte(newPwd)); e == nil {
		u.PasswdHash = h
		return e
	} else {
		return e
	}
}

func (u *User) VerifyPasswd(plainPwd string) error {
	return ComparePasswords(u.PasswdHash, []byte(plainPwd))
}

func PasswdHash(pwd []byte) (string, error) {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	} else {
		// GenerateFromPassword returns a byte slice so we need to
		// convert the bytes to a string and return it
		return string(hash), nil
	}
}

func ComparePasswords(hashedPwd string, plainPwd []byte) error {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return err
	}

	return nil
}