// Solo.go - A small and beautiful golang blogging system, Solo's golang version.
// Copyright (C) 2017, b3log.org
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

package service

import (
	"math"

	"github.com/b3log/solo.go/model"
	"github.com/b3log/solo.go/util"
	"github.com/jinzhu/gorm"
)

var Article = &articleService{}

type articleService struct {
}

// Article pagination arguments of admin console.
const (
	adminConsoleArticleListPageSize    = 15
	adminConsoleArticleListWindowsSize = 20
)

func (srv *articleService) AddArticle(article *model.Article) error {
	tx := db.Begin()

	if err := tx.Create(article).Error; nil != err {
		tx.Rollback()

		return err
	}

	tx.Commit()

	return nil
}

func (srv *articleService) ConsoleGetArticles(page int) (ret []*model.Article, pagination *util.Pagination) {
	if 1 > page {
		page = 1
	}

	offset := (page - 1) * adminConsoleArticleListPageSize
	count := 0
	db.Model(model.Article{}).Select("id, title, tags, view_count, comment_count").Where(model.Article{Status: model.ArticleStatusPublished}).
		Order("ID desc").Count(&count).
		Offset(offset).Limit(adminConsoleArticleListPageSize).
		Find(&ret)

	pageCount := int(math.Ceil(float64(count) / adminConsoleArticleListPageSize))
	pagination = util.NewPagination(page, adminConsoleArticleListPageSize, pageCount, adminConsoleArticleListWindowsSize, count)

	return
}

func (srv *articleService) ConsoleGetArticle(id uint) *model.Article {
	ret := &model.Article{}
	if nil != db.First(ret, id).Error {
		return nil
	}

	return ret
}

func (srv *articleService) RemoveArticle(id uint) error {
	article := &model.Article{
		Model: gorm.Model{ID: id},
	}

	return db.Delete(article).Error
}

func (srv *articleService) UpdateArticle(article *model.Article) error {
	return db.Model(&model.Article{}).Updates(article).Error
}