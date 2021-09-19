package constants

// The main SignUp url
const URL = "https://spclient.wg.spotify.com/signup/public/v1/account"

// The headers to pass with the request
var HEADERS = map[string]string{
	"authority":          "spclient.wg.spotify.com",
	"sec-ch-ua":          `"Google Chrome";v="93", " Not;A Brand";v="99", "Chromium";v="93"`,
	"sec-ch-ua-mobile":   "?0",
	"user-agent":         `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36`,
	"sec-ch-ua-platform": `"Windows"`,
	"content-type":       "application/x-www-form-urlencoded",
	"accept":             "*/*",
	"origin":             "https://www.spotify.com",
	"sec-fetch-site":     "same-site",
	"sec-fetch-mode":     "cors",
	"sec-fetch-dest":     "empty",
	"referer":            "https://www.spotify.com/",
	"accept-language":    "en-US,en;q=0.9",
}
