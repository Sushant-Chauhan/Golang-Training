// To pass the authenticated user's details
// managing user data (such as user ID and admin status
//can then be accessed by various middlewares and controllers
// useful for actions such as account creation, deposit, withdrawal

package middlewares

import (
	"context"
)

type userContextKey struct{}

// SetUserContext stores the userID and isAdmin in the context.
// function used by the authentication middleware after validating a JWT.
func SetUserContext(ctx context.Context, userID int, isAdmin bool) context.Context {
	return context.WithValue(ctx, userContextKey{}, map[string]interface{}{
		"userID":  userID,
		"isAdmin": isAdmin,
	})
}

// GetUserFromContext retrieves userID and admin status from the context.
func GetUserFromContext(ctx context.Context) (int, bool) {
	userDetails := ctx.Value(userContextKey{}).(map[string]interface{})
	userID := userDetails["userID"].(int)
	isAdmin := userDetails["isAdmin"].(bool)
	return userID, isAdmin
}

// function to simulate checking whether a user is active (for non-admin users)
func CheckUserActive(userID int) bool {
	//i can check in the database whether the user is active.
	// returing  true for now
	return true
}
