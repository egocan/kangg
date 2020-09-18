package wechat

import (
	"github.com/silenceper/wechat/message"
)

type MessageController struct {
	BaseController
}

func ResponseMsgType(msg message.MixMessage) *message.Reply {

	switch msg.MsgType {
		//文本消息
		case message.MsgTypeText:
			//回复消息：演示回复用户发送的消息
			text := message.NewText(msg.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			//图片消息
		case message.MsgTypeImage:
			//do something
			text := message.NewText("这是图片！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			//语音消息
		case message.MsgTypeVoice:
			//do something
			text := message.NewText("这是语音！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			//视频消息
		case message.MsgTypeVideo:
			//do something
			text := message.NewText("这是视频！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			//小视频消息
		case message.MsgTypeShortVideo:
			text := message.NewText("这是小视频！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
			//do something

			//地理位置消息
		case message.MsgTypeLocation:
			text := message.NewText("这是地理位置！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
			//do something

			//链接消息
		case message.MsgTypeLink:
			//do something
			text := message.NewText("这是衔接！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			//事件推送消息
		case message.MsgTypeEvent:
			return ResponseMsgTypeEvent(msg)

		default:
			text := message.NewText("逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}
}

func ResponseMsgTypeEvent(msg message.MixMessage) *message.Reply {
	//事件推送消息

	switch msg.Event {
		//EventSubscribe 订阅
		case message.EventSubscribe:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		//取消订阅
		case message.EventUnsubscribe:
			//do something
			text := message.NewText("你不好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		//用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
		case message.EventScan:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 上报地理位置事件
		case message.EventLocation:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 点击菜单拉取消息时的事件推送
		case message.EventClick:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 点击菜单跳转链接时的事件推送
		case message.EventView:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 扫码推事件的事件推送
		case message.EventScancodePush:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 扫码推事件且弹出“消息接收中”提示框的事件推送
		case message.EventScancodeWaitmsg:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 弹出系统拍照发图的事件推送
		case message.EventPicSysphoto:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 弹出拍照或者相册发图的事件推送
		case message.EventPicPhotoOrAlbum:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 弹出微信相册发图器的事件推送
		case message.EventPicWeixin:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		// 弹出地理位置选择器的事件推送
		case message.EventLocationSelect:
			//do something
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

		default:
			text := message.NewText("你好逗比！")
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	}

}

/*func (c *MessageController) Hello() {


	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	token := "b8cf671eaa1a270a9b53ddb894dd9029"
	echostr := c.GetString("echostr")

	var tempArray  = []string{token, timestamp, nonce}
	sort.Strings(tempArray)
	//将三个参数字符串拼接成一个字符串进行sha1加密
	var sha1String string = ""
	for _, v := range tempArray {
		sha1String += v
	}
	h := sha1.New()
	h.Write([]byte(sha1String))
	sha1String = hex.EncodeToString(h.Sum([]byte("")))
	//获得加密后的字符串可与signature对比
	if sha1String == signature {
		print(echostr)
	} else {
		fmt.Println("验证失败")
	}


	c.Data["json"] = echostr
	c.ServeJSON()
	c.StopRun()
	//配置微信参数
	c.EnableRender = false
	wc := wechat.NewWechat(config)

	// 传入request和responseWriter
	server := wc.GetServer(ctx.Request, ctx.ResponseWriter)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {


		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}
	//发送回复的消息
	server.Send()
}*/