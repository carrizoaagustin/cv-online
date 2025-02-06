package auth

import "fmt"

func GenerateRedirectURL(providerName string) string {
	return fmt.Sprintf("/auth/%s/callback", providerName)
}
