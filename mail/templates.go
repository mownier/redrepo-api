package mail

type EmailContent struct {
	From 	string
	To 		string
	Mime 	string
	Subject	string
	Body	string
}

func GetEmailVerificationTemplate(code string, recipient string) []byte {

	emailContent 			:= new(EmailContent)
	emailContent.From 		= "From: Red Repo Team\n"
	emailContent.To 		= "To: " + recipient + "\n"
	emailContent.Subject	= "Subject: Account Verification\n"
	emailContent.Mime 		= "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	emailContent.Body 		= "<html><body>Welcome to Red Repo!<br><br>Your verification code: " + code + "</body></html>"

	return []byte(emailContent.From + emailContent.To + emailContent.Subject + emailContent.Mime + emailContent.Body)
}