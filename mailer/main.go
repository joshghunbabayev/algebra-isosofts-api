package mailer

import (
	"context"
	"fmt"
	"os"

	sw "github.com/getbrevo/brevo-go/lib"
)

// SendEmail avtomatik mail göndərmək üçün əsas funksiya
func SendEmail(toEmail, toName, subject, htmlContent string) error {
	// API açarını .env faylından və ya mühit dəyişənlərindən oxuyuruq
	apiKey := os.Getenv("BREVO_API_KEY")

	ctx := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
		Key: apiKey,
	})

	cfg := sw.NewConfiguration()
	client := sw.NewAPIClient(cfg)

	// Göndərən şəxs (Sizin sisteminiz)
	sender := &sw.SendSmtpEmailSender{
		Name:  "Joshghun Babayev",
		Email: "joshghunbabayev@gmail.com",
	}

	// Qəbul edən şəxs
	recipient := sw.SendSmtpEmailTo{
		Email: toEmail,
		Name:  toName,
	}

	// Mail konfiqurasiyası
	emailPayload := sw.SendSmtpEmail{
		Sender:      sender,
		To:          []sw.SendSmtpEmailTo{recipient},
		Subject:     subject,
		HtmlContent: htmlContent, // HTML formatında mesaj
	}

	// Göndərmə prosesi
	_, resp, err := client.TransactionalEmailsApi.SendTransacEmail(ctx, emailPayload)
	if err != nil {
		return fmt.Errorf("mail göndərilərkən xəta baş verdi: %v", err)
	}

	fmt.Printf("Mail uğurla göndərildi! Status: %s\n", resp.Status)
	return nil
}
