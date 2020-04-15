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
	"net/http"
	"time"

	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util"
	"github.com/gorilla/mux"
)

func loggingMiddleware(s *service.Service) mux.MiddlewareFunc {
	return func(inner http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			id := r.Header.Get("X-Request-ID")
			if id == "" {
				id = util.NewUUID()
			}

			s.Logger().Trace( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Request",
				"id", id,
				"method", r.Method,
				"uri", r.RequestURI,
			)

			inner.ServeHTTP(w, r)

			s.Logger().Trace( //::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
				"Request",
				"id", id,
				"elapsed", time.Since(startTime).String(),
			)
		})
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
