package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	Id        uuid.UUID   `json:"id" db:"id"`
	Picture   pgtype.Text `json:"picture" db:"picture"`
	FullName  string      `json:"full_name" db:"full_name"`
	Email     string      `json:"email" db:"email"`
	Password  string      `json:"password" db:"password"`
	Address   string      `json:"address" db:"address"`
	Phone     string      `json:"phone" db:"phone"`
	RoleId    int         `json:"role_id" db:"role_id"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}

// import (
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/jackc/pgx/v5/pgtype"
// )

// // ========================================================================= REQUEST

// type LoginInput struct {
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required"`
// }

// type UpdateInput struct {
// 	Email    *string `json:"email"`
// 	Picture  *string `json:"picture"`
// 	FullName *string `json:"full_name"`
// 	Password *string `json:"password"`
// 	Address  *string `json:"address"`
// 	RoleId   *int    `json:"role_id"`
// 	Phone    *string `json:"phone"`
// }

// // ============================================================================= RESPONSE

// // Picture pgtype.Text

// type UserResponse struct {
// 	Id       uuid.UUID `json:"id"`
// 	Picture  string    `json:"picture"`
// 	FullName string    `json:"full_name"`
// 	Email    string    `json:"email"`
// 	RoleId   int       `json:"role_id"`
// 	Address  string    `json:"address"`
// 	Phone    string    `json:"phone"`
// }
