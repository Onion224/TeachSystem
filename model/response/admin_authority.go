package response

import "server/model"

type AuthorityResponse struct {
	Authority model.Authority `json:"authority"`
}

type AuthorityCopyResponse struct {
	Authority      model.User `json:"authority"`
	OldAuthorityId string     `json:"old_authority_id"`
}
