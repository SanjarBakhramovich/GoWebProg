package authenticate

import "net/http"

func Authenticate(w http.ResponseWriter, r *http.Request) {
	r.ParseFiles
	user, _ := data.UserByEmail(r.PostFormValue("email"))
	if user.Password == data.Encrypt(r.PostFromValue("password")) {
		session := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// 31.5 - 32 str
