package wechat

import (
	"fmt"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/menu"
)

type MenuController struct {
	BaseController
}

func (c *MenuController) CreateMenu()  {

	c.EnableRender = false

	wc := wechat.NewWechat(config)
	mu := wc.GetMenu()

	buttons := make([]*menu.Button, 1)
	btn := new(menu.Button)

	//创建click类型菜单
	btn.SetClickButton("name", "key123")
	buttons[0] = btn

	//设置btn为二级菜单
	btn2 := new(menu.Button)
	btn2.SetSubButton("subButton", buttons)

	buttons2 := make([]*menu.Button, 1)
	buttons2[0] = btn2

	log.Infoln(&buttons2)
	//发送请求
	err := mu.SetMenu(buttons2)
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}

}
