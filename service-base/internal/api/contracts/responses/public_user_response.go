package responses

import (
	"github.com/CallumLewisGH/BlueprintProject-GoVue/service-base/internal/domain/user"
	"github.com/google/uuid"
)

type UserPublicProfileResponse struct {
	Body userPublicProfile
}
type UserPublicProfileResponses struct {
	Body []userPublicProfile
}

type userPublicProfile struct {
	ID             uuid.UUID `json:"id"`
	Username       string    `json:"username"`
	ProfilePicture *string   `json:"profilePicture,omitempty"`
	Bio            string    `json:"bio,omitempty"`
	CreatedAt      string    `json:"createdAt"`
}

func ToUserPublicProfileResponse(dto *user.UserDTO) *UserPublicProfileResponse {
	return &UserPublicProfileResponse{
		Body: userPublicProfile{
			ID:             dto.ID,
			Username:       dto.Username,
			ProfilePicture: dto.ProfilePicture,
			Bio:            dto.Bio,
			CreatedAt:      formatTime(dto.CreatedAt),
		},
	}
}

func ToUserPublicProfileResponses(dtos []user.UserDTO) *UserPublicProfileResponses {
	responses := make([]userPublicProfile, len(dtos))
	for i, dto := range dtos {
		responses[i] = userPublicProfile{
			ID:             dto.ID,
			Username:       dto.Username,
			ProfilePicture: dto.ProfilePicture,
			Bio:            dto.Bio,
			CreatedAt:      formatTime(dto.CreatedAt),
		}
	}
	return &UserPublicProfileResponses{
		Body: responses,
	}
}
