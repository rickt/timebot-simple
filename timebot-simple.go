// http://timebot-simple.appspot.com/time
// http://timebot-simple.appspot.com/time?place=lax (etc)

package timebotsimple

import (
	"fmt"
	"net/http"
	"time"
)

// constants
const (
	shortformat = "3:04pm"
)

// types
type times struct {
	jpn time.Time
	lax time.Time
	utc time.Time
}

func init() {
	// setup the http handlers
	http.HandleFunc("/", handler_redirect)
	http.HandleFunc("/time", handler_time)
}

// redirect
func handler_redirect(w http.ResponseWriter, r *http.Request) {
	// redirect / to /time
	http.Redirect(w, r, "/time", 302)
}

// print the current time in Japan, LA, London/UTC
func handler_time(w http.ResponseWriter, r *http.Request) {
	// example call from slack
	// /time?place=jpn&token=REDACTED&team_id=REDACTED&team_domain=REDACTED&channel_id=REDACTED&channel_name=directmessage&user_id=REDACTED&user_name=REDACTED&command=%2Fjapantime&text=&response_url=REDACTED
	place := r.URL.Query().Get("place")
	// user := r.URL.Query().Get("user_name")
	switch place {
	case "healthcheck":
		fmt.Fprintf(w, "ok\n")
	default:
		loc, _ := time.LoadLocation(place)
		fmt.Fprintf(w, time.Now().In(loc).Format(shortformat)+" in "+place+"\n")
	}
}
