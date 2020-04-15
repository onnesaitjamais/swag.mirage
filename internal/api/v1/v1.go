/*
#######
##                                         _
##        ____    _____ ____ _      __ _  (_)______ ____ ____
##       (_-< |/|/ / _ `/ _ `/ _   /  ' \/ / __/ _ `/ _ `/ -_)
##      /___/__,__/\_,_/\_, / (_) /_/_/_/_/_/  \_,_/\_, /\__/
##                     /___/                       /___/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package v1

import (
	"fmt"
	"net/http"

	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util"
	"github.com/arnumina/swag/util/failure"
	"github.com/arnumina/swag/util/renderer"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/gorilla/mux"
)

const _resultDone = "done"

func getAllServices(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rr := renderer.New(w)

		services, err := s.Registry().List()
		if err != nil {
			rr.Error500(err)
			return
		}

		if err := rr.JSONOk(services); err != nil {
			rr.Error500(err)
			return
		}
	}
}

func startService(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rr := renderer.New(w)
		vars := mux.Vars(r)
		name := vars["name"]

		if name == s.Name() {
			rr.Error(
				http.StatusBadRequest,
				failure.New(nil).
					Set("service", name).
					Msg("this operation is not allowed for this service"), /////////////////////////////////////////////
			)

			return
		}

		conn, err := dbus.New()
		if err != nil {
			rr.Error500(err)
			return
		}

		defer conn.Close()

		sdInstance := util.NewUUID()
		ch := make(chan string, 1)

		_, err = conn.StartUnit(fmt.Sprintf("swag.%s@%s.service", name, sdInstance), "fail", ch)
		if err != nil {
			rr.Error500(err)
			return
		}

		if result := <-ch; result != _resultDone {
			rr.Error500(
				failure.New(nil).
					Set("service", name).
					Set("reason", result).
					Msg("impossible to start this service"), ///////////////////////////////////////////////////////////
			)

			return
		}

		if err := rr.JSONOk(sdInstance); err != nil {
			rr.Error500(err)
			return
		}
	}
}

func stopService(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rr := renderer.New(w)

		conn, err := dbus.New()
		if err != nil {
			rr.Error500(err)
			return
		}

		defer conn.Close()

		vars := mux.Vars(r)
		name := vars["name"]
		sdInstance := vars["instance"]

		var ch chan string

		if name != s.Name() {
			ch = make(chan string, 1)
		}

		_, err = conn.StopUnit(fmt.Sprintf("swag.%s@%s.service", name, sdInstance), "fail", ch)
		if err != nil {
			rr.Error500(err)
			return
		}

		if ch != nil {
			if result := <-ch; result != _resultDone {
				rr.Error500(
					failure.New(nil).
						Set("service", name).
						Set("instance", sdInstance).
						Set("reason", result).
						Msg("impossible to stop this service"), ////////////////////////////////////////////////////////
				)

				return
			}
		}

		rr.NoContent()
	}
}

func restartService(s *service.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rr := renderer.New(w)
		vars := mux.Vars(r)
		name := vars["name"]
		sdInstance := vars["instance"]

		if name == s.Name() && sdInstance != s.SdInstance() {
			rr.Error(
				http.StatusBadRequest,
				failure.New(nil).
					Set("service", name).
					Set("instance", sdInstance).
					Msg("this operation is not allowed for this instance of this service"), ////////////////////////////
			)

			return
		}

		conn, err := dbus.New()
		if err != nil {
			rr.Error500(err)
			return
		}

		defer conn.Close()

		var ch chan string

		if name != s.Name() {
			ch = make(chan string, 1)
		}

		_, err = conn.RestartUnit(fmt.Sprintf("swag.%s@%s.service", name, sdInstance), "fail", ch)
		if err != nil {
			rr.Error500(err)
			return
		}

		if ch != nil {
			if result := <-ch; result != _resultDone {
				rr.Error500(
					failure.New(nil).
						Set("service", name).
						Set("reason", result).
						Msg("impossible to restart this service"), /////////////////////////////////////////////////////
				)

				return
			}
		}

		rr.NoContent()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
