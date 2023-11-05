package smtp

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Smtp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Err loading .env file")
	}
	emailAppKey := os.Getenv("EMAIL_APP_KEY")

	fmt.Println(emailAppKey)
	// key는 당연히 deprecated
	// auth := smtp.PlainAuth("", "sopt32.palmsprings@gmail.com", "zzdzdvoyfuxgjnwk", "smtp.gmail.com")

	// from := "sopt32.palmsprings@gmail.com"
	// to := []string{"kandy1002@naver.com"}

	// headerSubject := "Subject: 반가워!!\r\n"
	// headerBlank := "\r\n"

	// body := `
	// 	<h1>
	// 		안녕하세요?
	// 	</h1>
	// `

	// msg := []byte(headerSubject + headerBlank + body)

	// err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
