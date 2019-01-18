package cmd

import (
	"fmt"

	"github.com/fatih/color"

	"github.com/sakib0hasan/yogo/inbox"
)

func RenderInboxMail(in *inbox.Inbox) []inbox.Mail {
	if in.Count() == 0 {
		info("Inbox is empty")
	}
	return in.GetAll()
}

func RenderMail(mail *inbox.Mail) {
	output("---\n")
	output(fmt.Sprintf("From  : %s <%s>\n", color.MagentaString(mail.Sender.Name), color.MagentaString(mail.Sender.Mail)))
	output(fmt.Sprintf("Title : %s\n", color.YellowString(mail.Title)))
	output(fmt.Sprintf("Date  : %s\n", color.BlueString(mail.Date.Format("2006-01-02 15:04"))))
	output("---\n")
	output(color.CyanString(mail.Body))
	output("\n---\n")
}
