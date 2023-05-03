package dtos

import "user-services/domain/entities"

type GetUsersDto struct {
	Meta MetadataDto     `json:"metadata"`
	Data *entities.Users `json:"data"`
}

type GetUserDto struct {
	Meta MetadataDto    `json:"meta"`
	Data *entities.User `json:"data"`
}
