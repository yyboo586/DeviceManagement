package logics

import (
	"bytes"
	"context"
	"html/template"
	"sync"

	gomail "github.com/go-mail/mail/v2"
)

var (
	mailerOnce          sync.Once
	mailerLogicInstance *mailer
)

type mailer struct {
	dialer *gomail.Dialer
	from   string
}

// TemplateData 模板数据结构
type TemplateData struct {
	Content string
}

func NewMailer() *mailer {
	mailerOnce.Do(func() {
		mailerLogicInstance = &mailer{
			dialer: &gomail.Dialer{
				Host:     "sandbox.smtp.mailtrap.io",
				Port:     25,
				Username: "8df5de08b5f13f",
				Password: "fcb5034938135d",
				SSL:      false,
			},
			from: "noreply@example.com",
		}
	})
	return mailerLogicInstance
}

func (m *mailer) getTemplate() (tmpl *template.Template, err error) {
	tmpl, err = template.New("email").Parse(templateStr)
	if err != nil {
		return nil, err
	}

	return
}

// SendTemplateMail 使用模板发送邮件
func (m *mailer) SendTemplateMail(ctx context.Context, to string, data TemplateData) (err error) {
	tmpl, err := m.getTemplate()
	if err != nil {
		return err
	}

	// 执行模板生成纯文本内容
	var plainBody bytes.Buffer
	if err := tmpl.ExecuteTemplate(&plainBody, "plainBody", data); err != nil {
		return err
	}

	// 执行模板生成HTML内容
	var htmlBody bytes.Buffer
	if err := tmpl.ExecuteTemplate(&htmlBody, "htmlBody", data); err != nil {
		return err
	}

	// 获取主题
	var subject bytes.Buffer
	if err := tmpl.ExecuteTemplate(&subject, "subject", data); err != nil {
		return err
	}

	msg := gomail.NewMessage()

	msg.SetHeader("To", to)
	msg.SetHeader("From", m.from)
	msg.SetHeader("Subject", subject.String())
	msg.SetBody("text/plain", plainBody.String())
	msg.AddAlternative("text/html", htmlBody.String())

	return m.dialer.DialAndSend(msg)
}

var templateStr = `
{{define "subject"}}Device Alarm!{{end}}

{{define "plainBody"}}
Hi, Someone has triggered an alarm on your device.
{{.Content}}
{{end}}

{{define "htmlBody"}}
<!doctype html>
<html>

<head>
    <meta name="viewport" content="width=device-width">
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
</head>

<body>
    <p>Hi,</p>
    <p>Someone has triggered an alarm on your device.</p>
    <p>{{.Content}}</p>
</body>

</html>
{{end}}
`
