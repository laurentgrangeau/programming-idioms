package pigae

import "google.golang.org/appengine"

// Env encapsulates a Programming-Idioms webapp environment.
type Env struct {
	// Host depends on the target environment.
	// It should not have a trailing slash.
	IsDev           bool
	Host            string
	UseAbsoluteUrls bool
	UseMinifiedCss  bool
	UseMinifiedJs   bool
}

//
// Prod
//

var envProd = Env{
	IsDev:           false,
	Host:            "http://www.programming-idioms.org",
	UseAbsoluteUrls: true,
	UseMinifiedCss:  true,
	UseMinifiedJs:   true,
}

//
// Dev
//

var envDev = Env{
	IsDev:           true,
	Host:            "http://localhost:8080",
	UseAbsoluteUrls: true,
	UseMinifiedCss:  false,
	UseMinifiedJs:   false,
}

// Which one is used ?

var env Env

func initEnv() {
	if appengine.IsDevAppServer() {
		env = envDev
	} else {
		env = envProd
	}
}
