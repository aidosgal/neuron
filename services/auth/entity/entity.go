package entity

import "github.com/aidosgal/neuron/ent"

type (
	Admin struct {
		ID       int    `json:"id"`
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token string `json:"token"`
		Admin Admin `json:"admin"`
	}
)

func MakeStorageAdminToEntity(admin *ent.Admin) *Admin {
	return &Admin{
		ID:       admin.ID,
		Login:    admin.Login,
		Password: admin.Password,
	}
}
