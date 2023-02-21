package cookies

import (
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func ConfigCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, token string) error {
	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	codificatedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    codificatedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
