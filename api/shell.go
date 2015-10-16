package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/megamsys/libgo/errors"
	"github.com/megamsys/megamd/api/context"
	"golang.org/x/net/websocket"
)

func remoteShellHandler(ws *websocket.Conn) {
	var httpErr *errors.HTTP
	defer func() {
		defer ws.Close()
		if httpErr != nil {
			var msg string
			switch httpErr.Code {
			case http.StatusUnauthorized:
				msg = "no token provided or session expired, please login again\n"
			default:
				msg = httpErr.Message + "\n"
			}
			ws.Write([]byte("Error: " + msg))
		}
	}()
	r := ws.Request()
	token := context.GetAuthToken(r)
	if token == nil {
		httpErr = &errors.HTTP{
			Code:    http.StatusUnauthorized,
			Message: "no token provided",
		}
		return
	}
	user, err := token.User()
	if err != nil {
		httpErr = &errors.HTTP{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	appName := r.URL.Query().Get(":appname")
	_, err = getApp(appName)
	if err != nil {
		if herr, ok := err.(*errors.HTTP); ok {
			httpErr = herr
		} else {
			httpErr = &errors.HTTP{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
		return
	}
	unitID := r.URL.Query().Get("unit")
	width, _ := strconv.Atoi(r.URL.Query().Get("width"))
	height, _ := strconv.Atoi(r.URL.Query().Get("height"))
	term := r.URL.Query().Get("term")
	fmt.Printf("%s %s %s %s %s", user, unitID, width, height, term)
	/*opts := provision.ShellOptions{
		Conn:   ws,
		Width:  width,
		Height: height,
		Unit:   unitID,
		Term:   term,
	}
	err = box.Shell(opts)
	if err != nil {
		httpErr = &errors.HTTP{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	*/
}

func getApp(a string) (string, error) {
	return "", nil
}
