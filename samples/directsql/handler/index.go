/**
 * desc : 登录后起始页
 * author:畅雨
 * date:  2016.05.16
 * history:
 *
 */
package handler

import (
	"github.com/henrylee2cn/thinkgo"
)

func Index() thinkgo.HandlerFunc {
	return func(c *thinkgo.Context) error {
		return c.Render(200, "view/index.html", nil)
	}
}