package models

import "github.com/operandinc/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
