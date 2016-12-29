package linbook

import (
	"linbook/email"
	"net/smtp"

	"github.com/go-martini/martini"
)

func main() {

	m := martini.Classic()
	m.Get("/sendmail", func() {
		sendMail()
	})
	m.RunOnAddr("80")

}

func sendMail() {
	//test
	e := email.NewEmail()
	e.From = "Jordan Wright <alexlin@yeah.net>"
	e.To = []string{"linhuifeng@minday.com"}
	//e.Bcc = []string{"test_bcc@example.com"}
	//e.Cc = []string{"test_cc@example.com"}
	e.Subject = "Awesome Subject"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
	e.Send("smtp.yeah.net:25", smtp.PlainAuth("", "alexlin@yeah.net", "taisoga@1", "smtp.yeah.net"))
}
