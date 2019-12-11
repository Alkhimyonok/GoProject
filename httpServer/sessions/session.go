package sessions

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("S3CR3TK3Y"))

func Flash(r *http.Request, w http.ResponseWriter) string {
	var message string = ""
	session, _ := Store.Get(r, "session")
	untypedMessage := session.Values["Message"]
	message, ok := untypedMessage.(string)
	if !ok {
		return ""
	}
	delete(session.Values, "MESSAGE")
	session.Save(r, w)
	return message
}
