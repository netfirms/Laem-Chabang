package motd

import (
	"net/http"

	"github.com/portainer/libhttp/response"
	"github.com/netfirms/Laem-Chabang/api"
	"github.com/netfirms/Laem-Chabang/api/crypto"
	"github.com/netfirms/Laem-Chabang/api/http/client"
)

type motdResponse struct {
	Message string `json:"Message"`
	Hash    []byte `json:"Hash"`
}

func (handler *Handler) motd(w http.ResponseWriter, r *http.Request) {

	motd, err := client.Get(portainer.MessageOfTheDayURL, 0)
	if err != nil {
		response.JSON(w, &motdResponse{Message: ""})
		return
	}

	hash := crypto.HashFromBytes(motd)
	response.JSON(w, &motdResponse{Message: string(motd), Hash: hash})
}
