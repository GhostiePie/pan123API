package APIs

import (
	"encoding/json"
)

type GetUploadDomainResponse struct {
	Response
	Data []string `json:"data"`
}
