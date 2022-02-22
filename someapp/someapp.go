package someapp

import "fmt"

type someapp struct {
	usename string
	use     int
	isopen  bool
}

func (app *someapp) open() {
	if app.isopen {
		fmt.Printf("%sさんのアプリは既に開いています¥n¥n", app.usename)
	} else {
		app.use++
		if app.use > 2 {
			fmt.Printf("%sさんの使用期間は満了です。ご購入をご検討ください¥n¥n", app.usename)
		} else {
			fmt.Printf("こんにちは%sさん、%d回目のご使用になります¥n¥n", app.usename, app.use)
		}
	}
}

func (app *someapp) close() {
	if app.isopen {
		fmt.Printf("さよなら%sさん、アプリを終了します。¥n¥n", app.usename)
	}
}

func newapp(username string) someapp {
	fmt.Printf("ようこそ%sさん、初めてのご利用になります¥n", username)
	return someapp{username, 1, true}
}

func Someapp() {
	ringo := newapp("りんご")
	mikan := newapp("みかん")

	ringo.open()
	ringo.open()
	mikan.close()
}
