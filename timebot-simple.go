// http://timebot-simple.appspot.com/time?tz=America/Los_Angeles
// http://timebot-simple.appspot.com/time?tz=Asia/Tokyo
// http://timebot-simple.appspot.com/time?tz=Asia/UTC (etc)

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

// there's no main() func in google app engine
func init() {
	// setup the http handlers
	http.HandleFunc("/", handler_redirect)
	http.HandleFunc("/time", handler_time)
}

// redirect / to /time
func handler_redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/time", 302)
}

// print the current time in Japan, LA, London/UTC
func handler_time(w http.ResponseWriter, r *http.Request) {
	// example call from slack
	// /time?tz=Asia/Tokyo&token=REDACTED&team_id=REDACTED&team_domain=REDACTED&channel_id=REDACTED&channel_name=directmessage&user_id=REDACTED&user_name=REDACTED&command=%2Fjapantime&text=&response_url=REDACTED
	tz := r.URL.Query().Get("tz")
	switch tz {
	case "healthcheck":
		fmt.Fprintf(w, "ok\n")
		return
	case "":
		loc, _ := time.LoadLocation("America/Los_Angeles")
		fmt.Fprintf(w, time.Now().In(loc).Format(shortformat)+" in "+tz+"\n")
		return
	default:
		loc, err := time.LoadLocation(tz)
		if err != nil {
			fmt.Fprintf(w, "I'm sorry, but I was not able to find a timezone matching \""+tz+"\"! :confused: \n")
			return
		}
		fmt.Fprintf(w, time.Now().In(loc).Format(shortformat)+" in "+tz+"\n")
		return
	}
}
