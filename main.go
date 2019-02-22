package yogo

import (
	"errors"
	"fmt"
	"github.com/sakib0hasan/yogo/inbox"
	"strings"
	"time"
)

func List(identifier string, offset int) ([]inbox.Mail, error) {
	in, err := inbox.ParseInboxPages(identifier, offset)
	if err != nil {
		fmt.Println(err)
	}
	if in.Count() == 0 {
		return nil, errors.New("inbox is empty")
	}
	mails := in.GetAll()
	for idx, _ := range mails {
		in.Parse(idx)
	}
	return mails, nil
}

func CheckEmail(Email string, PassContent string, StartTime string, EndTime string) (inbox.Mail, error) {
	Email = strings.Replace(Email, "@yopmail.com", "", -1)
	PassContentParts := strings.Split(PassContent, "|")
	mails, emailErr := List(Email, 3)
	if emailErr != nil {
		return inbox.Mail{}, errors.New("inbox is empty")
	} else {
		for _, mail := range mails {
			layout := "2006-01-02 15:04:05 -0700 MST" //-0800 PST if not working
			t1, err1 := time.Parse(layout, StartTime)
			t2, err2 := time.Parse(layout, EndTime)

			fmt.Println("t1", t1.String())
			fmt.Println("t2", t2.String())
			fmt.Println("mail", mail.Date.String())
			fmt.Println(mail.Title)

			if mail.Date.After(t1) && mail.Date.Before(t2) {
				fmt.Println("TIME IN RANGE")
			}

			if err1 == nil && err2 == nil {
				if mail.Date.After(t1) && mail.Date.Before(t2) {
					for _, passContent := range PassContentParts{
						if strings.Contains(strings.ToLower(mail.Title), strings.ToLower(passContent)) || strings.Contains(strings.ToLower(mail.Body), strings.ToLower(passContent)) {
							return mail, nil
						}
					}
				}
			}
		}
	}
	return inbox.Mail{}, errors.New("pass content not found")
}

func CheckOffset(count int, offset int) error {
	if count == 0 {
		return errors.New("inbox is empty")
	}
	if count < offset-1 {
		return errors.New("lower your offset value")
	}
	return nil
}
