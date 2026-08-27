package responses

import (
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
)

type UserPrivateProfileResponse struct {
	Body userPrivateProfile
}

type userPrivateProfile struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	EmailVerified  bool      `json:"emailVerified"`
	ProfilePicture *string   `json:"profilePicture,omitempty"`
	Bio            string    `json:"bio,omitempty"`
	Timezone       string    `json:"timezone,omitempty"`
	Role           string    `json:"role"`
	CreatedAt      string    `json:"createdAt"`
	UpdatedAt      string    `json:"updatedAt"`
}

func ToUserPrivateProfileResponse(dto *user.UserDTO) *UserPrivateProfileResponse {
	return &UserPrivateProfileResponse{
		Body: userPrivateProfile{
			ID:             dto.ID,
			Username:       dto.Username,
			Email:          dto.Email,
			EmailVerified:  dto.EmailVerified,
			ProfilePicture: dto.ProfilePicture,
			Bio:            dto.Bio,
			Timezone:       dto.Timezone,
			Role:           dto.Role,
			CreatedAt:      formatTime(dto.CreatedAt),
			UpdatedAt:      formatTime(dto.UpdatedAt),
		},
	}
}
