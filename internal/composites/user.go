package composites

import (
	"jwt_registration_api/internal/adapters/api/http_handlers"
	http_handlers_user "jwt_registration_api/internal/adapters/api/http_handlers/user"
	db_user "jwt_registration_api/internal/adapters/db/user"
	domain_user "jwt_registration_api/internal/domain/user"

	"github.com/sirupsen/logrus"
)

type UserComposite struct {
	Handler http_handlers.Handler
	Service http_handlers_user.Service
	Storage domain_user.Storage
}

func NewUserComposite(postgresComposite *PostgresComposite, logger *logrus.Logger) (*UserComposite, error) {
	storage := db_user.NewStorage(postgresComposite.Db)
	service := domain_user.NewService(storage)
	handler := http_handlers_user.NewHandler(service, logger)

	return &UserComposite{
		Handler: handler,
		Service: service,
		Storage: storage,
	}, nil
}
