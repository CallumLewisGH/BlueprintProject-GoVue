package user

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	//Standard
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Authentication
	Username      string `gorm:"size:50;uniqueIndex;not null"`
	Email         string `gorm:"size:255;uniqueIndex;not null"`
	EmailVerified bool   `gorm:"default:false"`
	LastLogin     *time.Time
	AuthId        string `gorm:"size:255;uniqueIndex;not null" json:"-"`

	// Refresh session (single active session per user - a fresh login or
	// refresh overwrites whatever was here). Never exposed via UserDTO/API
	// responses - this is sensitive server-side auth state, not profile data.
	RefreshTokenHash      *string    `gorm:"size:64;uniqueIndex" json:"-"`
	RefreshTokenExpiresAt *time.Time `json:"-"`

	// Profile
	ProfilePicture *string `gorm:"text"`
	Bio            string  `gorm:"size:500"`

	// Preferences
	Timezone string `gorm:"size:50"`

	// Status
	IsActive      bool `gorm:"default:true"`
	IsBanned      bool `gorm:"default:false"`
	DeactivatedAt *time.Time

	// Authorization
	Role string `gorm:"size:20;not null;default:'user';index"`
}
