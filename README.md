# timebot-simple
####  a simple go Google appengine app that returns a nice, human-readable string with the current time in Los Angeles, Japan & London/UTC. intended for use as an example Slack "/slash" command, delivered via simple http responder running in Google's appengine go runtime. 

this app is currently up & running at [http://timebot-simple.appspot.com](http://timebot-simple.appspot.com).

there are 2x URL endpoints:
* `/` automatic redirect to `/time`
* `/time` returns a nice, human-readable string with the current time (in Los Angeles)

note that the `/time` URL endpoint accepts a `?place=` parameter:
* `/time?place=lax` returns the current time in Los Angeles
* `/time?place=jpn` returns the current time in Japan
* `/time?place=utc` returns the current time in UTC

-RMT

