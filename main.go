package main

import "fmt"

type Author struct {
	Focus int
	Name string
	VIP bool
	Signature string
	Lv int
}

type Content struct {
	DanMu int
	Goodlooking bool
	Instruction string
	Zan bool
	Cang bool
	Bi bool
}

type Re struct {
	up string
	name string
	hot bool
}

type Video struct {
	Who Author
	What Content
	Tuijian Recommend
}

type Recommend []Re//定义结构体切片用type

func (s *Video)Dianzan() {
	s.What.Zan = true
}

func (s *Video)Toubi() {
	s.What.Bi = true
}

func (s *Video)Shoucang() {
	s.What.Cang = true
}

type Sanlian interface {
	Toubi()
	Shoucang()
	Dianzan()
}

func Yijiansanlian(Video Sanlian) {
	Video.Toubi()
	Video.Shoucang()
	Video.Dianzan()
}

func main() {
	var LuoHan = Video{
		Who: Author{
			Focus: 2659000,
			Name: "LuoHan",
			VIP: true,
			Signature: "神~~经病",
			Lv: 6,
		},
		What: Content{
			DanMu: 6000,
			Goodlooking: false,
			Instruction: "常言说的好 事出反常必有妖 拖更太久必定憋大招!如果说一年前的万毒归宗是利用经验快速发育打出的效果 呢么今天的毒奶组合可以说是完善后的大成版!喜欢本期视频的朋友请长按点赞一条龙哦 谢谢大家啦~",
			Zan: false,
			Bi: false,
			Cang: false,
		},
		Tuijian: Recommend{
			Re{
				up: "罗汉",
				hot: true,
				name: "什么鬼套路",
			},
			Re{
				up: "胡桃",
				hot: true,
				name: "胡桃养成",
			},
		},
	}
	Yijiansanlian(&LuoHan)//“&”别掉了
	fmt.Println(LuoHan)
}
