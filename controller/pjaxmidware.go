// Pipe - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017-2018, b3log.org
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
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func pjax(c *gin.Context) {
	isPJAX := isPJAX(c)
	dataModelVal, _ := c.Get("dataModel")
	dataModel := dataModelVal.(*DataModel)
	(*dataModel)["pjax"] = isPJAX
	c.Set("dataModel", dataModel)

	if !isPJAX {
		c.Next()

		return
	}

	c.Writer = &pjaxHTMLWriter{c.Writer, &strings.Builder{}, c}
	c.Next()
}

type pjaxHTMLWriter struct {
	gin.ResponseWriter
	bodyBuilder *strings.Builder
	c           *gin.Context
}

func (p *pjaxHTMLWriter) Write(data []byte) (int, error) {
	p.bodyBuilder.Write(data)
	if !strings.HasSuffix(string(data), "</html>\r\n") {
		return 0, nil
	}

	pjaxContainer := p.c.Request.Header.Get("X-PJAX-Container")
	body := p.bodyBuilder.String()
	r := regexp.MustCompile(`<!---- pjax \{` + pjaxContainer + `\} start ---->([\s\S]*)<!---- pjax \{` + pjaxContainer + `\} end ---->`)
	containers := r.FindAllStringSubmatch(body, -1)
	if 0 == len(containers) {
		return p.ResponseWriter.WriteString(body)
	}

	builder := &strings.Builder{}
	for _, v := range containers {
		builder.WriteString(v[1])
	}

	return p.ResponseWriter.WriteString(builder.String())
}

func isPJAX(c *gin.Context) bool {
	return "true" == c.Request.Header.Get("X-PJAX") && "" != c.Request.Header.Get("X-PJAX-Container")
}