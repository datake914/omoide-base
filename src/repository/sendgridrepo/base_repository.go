package sendgridrepo

import (
	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
)

type Config struct {
	SendGridEndpoint   string
	SendGridApiKey     string
	SendGridTemplateId string
}

type BaseRepository struct {
	Cfg Config
}

func NewBaseRepository(cfg Config) BaseRepository {
	return BaseRepository{
		Cfg: cfg,
	}
}

func (repo BaseRepository) GetSendRequest() rest.Request {
	return sendgrid.GetRequest(repo.Cfg.SendGridApiKey, "/v3/mail/send", repo.Cfg.SendGridEndpoint)
}
