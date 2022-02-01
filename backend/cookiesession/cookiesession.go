package cookiesession

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("some-session-key"))
