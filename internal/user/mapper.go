package user

import "github.com/YoonDongGeun/go-backend-practice/internal/store"

func toUserEntity(row store.User) UserEntity {
	return UserEntity{
		ID:        row.ID,
		Email:     row.Email,
		Name:      row.Name,
		CreatedAt: row.CreatedAt.Time,
	}
}

func toUserResponseDTO(entity UserEntity) UserResponseDTO {
	return UserResponseDTO{
		ID:        entity.ID,
		Email:     entity.Email,
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt,
	}
}
