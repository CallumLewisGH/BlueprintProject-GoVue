package integrations

import (
	"fmt"
	"os"
	"strconv"
	"time"

	mail "github.com/wneessen/go-mail"
)

// SendEmail sends an HTML email over plain SMTP via go-mail
// (https://github.com/wneessen/go-mail) - a client built on the Go standard
// library's net/smtp, not a third-party SaaS API. Sending mail only depends
// on whatever SMTP server/credentials are configured
// (SMTP_HOST/SMTP_PORT/SMTP_USERNAME/SMTP_PASSWORD/SMTP_FROM_EMAIL), so any
// mailbox or transactional-email provider that exposes an SMTP endpoint
// (Gmail, Zoho Mail, Fastmail, SES, Postmark, etc.) works without a
// provider-specific integration.
//
// replyTo is optional - pass "" to omit it. It lets a message appear to come
// from Boondock's own address while still letting the recipient hit "Reply"
// and land in the actual sender's inbox (used by the contact form, where the
// visitor's email becomes the Reply-To rather than the From, since SMTP
// providers reject a From address they don't own).
func SendEmail(toEmail string, subject string, htmlBody string, replyTo string) error {
	host := os.Getenv("SMTP_HOST")
	if host == "" {
		return fmt.Errorf("SMTP_HOST is not set")
	}
	username := os.Getenv("SMTP_USERNAME")
	if username == "" {
		return fmt.Errorf("SMTP_USERNAME is not set")
	}
	password := os.Getenv("SMTP_PASSWORD")
	if password == "" {
		return fmt.Errorf("SMTP_PASSWORD is not set")
	}
	fromEmail := os.Getenv("SMTP_FROM_EMAIL")
	if fromEmail == "" {
		return fmt.Errorf("SMTP_FROM_EMAIL is not set")
	}

	port := 587
	if raw := os.Getenv("SMTP_PORT"); raw != "" {
		parsed, err := strconv.Atoi(raw)
		if err != nil {
			return fmt.Errorf("SMTP_PORT must be a number: %w", err)
		}
		port = parsed
	}

	msg := mail.NewMsg()

	if err := msg.FromFormat("Boondock", fromEmail); err != nil {
		return fmt.Errorf("invalid from address: %w", err)
	}
	if err := msg.To(toEmail); err != nil {
		return fmt.Errorf("invalid to address: %w", err)
	}
	if replyTo != "" {
		if err := msg.ReplyTo(replyTo); err != nil {
			return fmt.Errorf("invalid reply-to address: %w", err)
		}
	}
	msg.Subject(subject)
	msg.SetBodyString(mail.TypeTextHTML, htmlBody)

	client, err := mail.NewClient(host,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover),
		mail.WithUsername(username),
		mail.WithPassword(password),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithTimeout(10*time.Second),
	)
	if err != nil {
		return fmt.Errorf("failed to build SMTP client: %w", err)
	}

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
