package auth

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/aleksrutins/litelytics/dbutil"
	"github.com/aleksrutins/litelytics/ent"
	"github.com/aleksrutins/litelytics/ent/user"
	"github.com/gofiber/fiber/v2"
)

func checkCredentials(ctx context.Context, email string, password string) (*ent.User, bool) {
	hash := sha256.Sum256([]byte(password))
	record, err := dbutil.Client.User.
		Query().
		Where(user.EmailEQ(email)).
		First(context.Background())

	if err != nil {
		return nil, false
	}
	return record, (string(hash[:]) == string(record.Password))
}

func authenticateRequest(c *fiber.Ctx, user *ent.User) {
	c.Cookie(&fiber.Cookie{
		Name:  "userId",
		Value: fmt.Sprint(user.ID),
	})
	c.Cookie(&fiber.Cookie{
		Name:  "userEmail",
		Value: user.Email,
	})
}

func GetUser(c *fiber.Ctx) *UserInfo {
	if c.Cookies("userId") != "" && c.Cookies("userEmail") != "" {
		return &UserInfo{
			ID:    c.Cookies("userId"),
			Email: c.Cookies("userEmail"),
		}
	} else {
		return nil
	}
}
