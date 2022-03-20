package handler

import (
	cookie "cookie/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//This var will be used to store session_token
var sessions = map[string]cookie.Session{}

func Login(w http.ResponseWriter, r *http.Request) {
	var credential cookie.Credentials
	//Decode reads from r
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pass, errGetPass := getUserFromDB(credential.Username)
	if errGetPass != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if pass != credential.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//Create new random token
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Second)

	//Set client cookie as session_token generated before
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})

	//Store session_token in session map
	sessions[sessionToken] = cookie.Session{
		Username: credential.Username,
		Expiry:   expiresAt,
	}

}

/** After user login, we have session information of that user stored on their end as cookies.
This info can be used to authenticate subsequent user requests,
Get information about user making the request.
*/
func Auth(w http.ResponseWriter, r *http.Request) (cookie.Session, string) {

	//Every request user made then get the session token from named cookie
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return cookie.Session{}, ""
		}

		w.WriteHeader(http.StatusBadRequest)
		return cookie.Session{}, ""
	}

	sessionToken := c.Value

	userSession, ok := sessions[sessionToken]
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return cookie.Session{}, sessionToken
	}
	if userSession.IsExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return cookie.Session{}, sessionToken
	}
	return userSession, sessionToken
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	auth, sessionToken := Auth(w, r)

	//If authentication is valid then create new session
	delete(sessions, sessionToken)
	newSessionToken := uuid.NewString()
	expiresAt := time.Now().Add(60 * time.Second)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   newSessionToken,
		Expires: expiresAt,
	})

	sessions[newSessionToken] = cookie.Session{
		Username: auth.Username,
		Expiry:   expiresAt,
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	_, sessionToken := Auth(w, r)

	delete(sessions, sessionToken)
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	user, _ := Auth(w, r)

	str := fmt.Sprintf("Hello, %s", user.Username)
	w.Write([]byte(str))
}

//This is just dummy data as if we already get the user info from database
func getUserFromDB(user string) (string, error) {

	var users = map[string]string{
		"username1": "userpassword1",
		"username2": "userpassword2",
	}

	return users[user], nil
}
