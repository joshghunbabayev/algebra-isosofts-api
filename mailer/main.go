package mailer

import (
	"context"
	"fmt"
	"os"

	sw "github.com/getbrevo/brevo-go/lib"
)

// EmailContact alıcı məlumatlarını (Email və Ad) saxlamaq üçündür
type EmailContact struct {
	Email string
	Name  string
}

// SendEmail birdən çox TO və CC (istəyə bağlı) alıcıya mail göndərir
func SendEmail(to []EmailContact, cc []EmailContact, subject, htmlContent string) error {
	apiKey := os.Getenv("BREVO_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("BREVO_API_KEY mühit dəyişəni tapılmadı")
	}

	ctx := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: apiKey,
	})

	cfg := sw.NewConfiguration()
	client := sw.NewAPIClient(cfg)

	// Göndərən şəxs (Brevo panelində təsdiqlənmiş olmalıdır)
	sender := &sw.SendSmtpEmailSender{
		Name:  "Joshghun Babayev",
		Email: "joshghunbabayev@gmail.com",
	}

	// 1. TO (Əsas Alıcılar) Siyahısını Hazırlayırıq
	if len(to) == 0 {
		return fmt.Errorf("ən azı bir əsas alıcı (TO) qeyd edilməlidir")
	}

	var brevoTo []sw.SendSmtpEmailTo
	for _, t := range to {
		name := t.Name
		brevoTo = append(brevoTo, sw.SendSmtpEmailTo{
			Email: t.Email,
			Name:  name,
		})
	}

	// Mail konfiqurasiyası (Payload)
	emailPayload := sw.SendSmtpEmail{
		Sender:      sender,
		To:          brevoTo,
		Subject:     subject,
		HtmlContent: htmlContent,
	}

	// 2. CC Siyahısını Hazırlayırıq (Əgər parametr olaraq göndərilibsə)
	if len(cc) > 0 {
		var brevoCc []sw.SendSmtpEmailCc
		for _, c := range cc {
			ccName := c.Name
			brevoCc = append(brevoCc, sw.SendSmtpEmailCc{
				Email: c.Email,
				Name:  ccName,
			})
		}
		// CC siyahısını payload-a əlavə edirik
		emailPayload.Cc = brevoCc
	}

	// Göndərmə prosesi
	_, resp, err := client.TransactionalEmailsApi.SendTransacEmail(ctx, emailPayload)
	if err != nil {
		// 400 Bad Request olduqda Brevo-nun qaytardığı dəqiq səbəbi görmək üçün resp-i də yazdırırıq
		return fmt.Errorf("mail göndərilərkən xəta baş verdi: %v | Brevo Cavabı: %+v", err, resp)
	}

	fmt.Printf("Mail uğurla göndərildi! Status koda uyğun cavab: %s\n", resp.Status)
	return nil
}
