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
	mails, emailErr := List(Email, 3)
	if emailErr != nil {
		return inbox.Mail{}, errors.New("inbox is empty")
	} else {
		for _, mail := range mails {
			layout := "2006-01-02 15:04:05 -0700 MST"
			t1, err1 := time.Parse(layout, StartTime)
			t2, err2 := time.Parse(layout, EndTime)
			if err1 == nil && err2 == nil {
				if mail.Date.After(t1) && mail.Date.Before(t2) {
					if strings.Contains(strings.ToLower(mail.Title), strings.ToLower(PassContent)) || strings.Contains(strings.ToLower(mail.Body), strings.ToLower(PassContent)) {
						return mail, nil
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
