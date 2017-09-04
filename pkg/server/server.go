package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// StartHTTPServer start the builtin http server that servers static
// configuration
func StartHTTPServer() {
	http.HandleFunc("/static/anaconda", anacondaConfigHandler)
	http.HandleFunc("/static/cloudinit", cloudInitConfigHandler)

	logrus.Infof("Starting HTTP Server on port 8080")
	if err := http.ListenAndServe(":1", nil); err != nil {
		logrus.Fatalf("HTTP Server failed: %#v", err)
	}
}

func anacondaConfigHandler(w http.ResponseWriter, r *http.Request) {

}

func cloudInitConfigHandler(w http.ResponseWriter, r *http.Request) {

}
