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
	jpn       time.Time
	lax       time.Time
	utc       time.Time
}

func init() {
	// setup the http handlers
	http.HandleFunc("/", handler_redirect)
	http.HandleFunc("/time", handler_time)
}

// get the current time in JPN, LAX, UTC and return them in a struct of 3x t.Time
func getTime() times {
	var tzu times
	t := time.Now()
	// locations
	loc_lax, _ := time.LoadLocation("America/Los_Angeles")
	loc_jpn, _ := time.LoadLocation("Japan")
	loc_utc, _ := time.LoadLocation("UTC")
	// times
	tzu.lax = t.In(loc_lax)
	tzu.jpn = t.In(loc_jpn)
	tzu.utc = t.In(loc_utc)
	return tzu
}

// redirect
func handler_redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/time", 302)
}

// print the current time in Japan, LA, London/UTC
func handler_time(w http.ResponseWriter, r *http.Request) {
	var mytimes times
	mytimes = getTime()
	place := r.URL.Query().Get("place")
	// user := r.URL.Query().Get("user_name")
	switch place {
	case "lax":
		fmt.Fprintf(w, mytimes.lax.Format(shortformat)+" in Los Angeles ("+mytimes.jpn.Format(shortformat)+" in Japan,  "+mytimes.utc.Format(shortformat)+" in London/UTC)\n")
	case "jpn":
		fmt.Fprintf(w, mytimes.jpn.Format(shortformat)+" in Japan ("+mytimes.lax.Format(shortformat)+" in Los Angeles, "+mytimes.utc.Format(shortformat)+" in London/UTC)\n")
	case "utc":
		fmt.Fprintf(w, mytimes.utc.Format(shortformat)+" in London/UTC ("+mytimes.lax.Format(shortformat)+" in Los Angeles, "+mytimes.jpn.Format(shortformat)+" in Japan)\n")
	default:
		fmt.Fprintf(w, mytimes.lax.Format(shortformat)+" in Los Angeles ("+mytimes.jpn.Format(shortformat)+" in Japan, "+mytimes.utc.Format(shortformat)+" in London/UTC)\n")
	}
}
