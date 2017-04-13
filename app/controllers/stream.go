package controllers

import (
	"fmt"
	"net/http"

	"github.com/allie/revel-stream/app"
	"github.com/revel/revel"
)

type Stream struct {
	*revel.Controller
}

func (c Stream) Auth() revel.Result {
	key := c.Params.Get("name")
	sql := fmt.Sprintf("SELECT user FROM users WHERE key='%s'", key)
	rows, err := app.DB.Query(sql)

	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderText("Auth: Malformed request")
	}

	result := rows.Next()
	if !result {
		c.Response.Status = http.StatusUnauthorized
		return c.RenderText("Auth: Incorrect credentials")
	}

	var user string
	_ = rows.Scan(&user)

	return c.Redirect(user)
}

func (c Stream) Watch() revel.Result {
	user := c.Params.Get("user")
	sql := fmt.Sprintf("SELECT * FROM users WHERE user='%s'", user)
	rows, err := app.DB.Query(sql)

	if err != nil {
		c.Response.Status = http.StatusBadRequest
		return c.RenderText(err.Error())
	}

	result := rows.Next()
	if !result {
		c.Response.Status = http.StatusNotFound
		return c.NotFound("User does not exist")
	}

	base := revel.Config.StringDefault("stream.baseurl", "")
	return c.Render(base, user)
}

func (c Stream) Index() revel.Result {
	return c.Render()
}
