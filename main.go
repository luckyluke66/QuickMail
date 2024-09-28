package main

import (
    mail "QuickMail/mailUtils"
	"fmt"
	"os"
	"strings"

    "github.com/luckyluke66/FuncSlice/slices"
    "github.com/joho/godotenv"
)

var filepath = "emails.txt"

func main() {
    godotenv.Load()

    creds := mail.Creds{
        Email: os.Getenv("EMAIL_ADDR"),
        Password: os.Getenv("EMAIL_PWD"),
    }

    data, err := os.ReadFile(filepath)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }

    emails := strings.Split(string(data), "\n")
    emails = slices.Map(emails, func(s string) string {return strings.TrimSpace(s)})
    message := mail.Message{Subject: "zkouska", Body: "toto je zkusebni email z me nove aplikace"}

    mail.SendMultiple(creds, emails, message)
}