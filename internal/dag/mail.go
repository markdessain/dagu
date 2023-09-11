package dag

type NftyConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type SmtpConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type MailConfig struct {
	From   string
	To     string
	Prefix string
}

type TopicConfig struct {
	Topic string
}
