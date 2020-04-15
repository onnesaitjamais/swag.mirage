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
	"github.com/arnumina/swag/service"
	"github.com/gorilla/mux"
)

// Routes AFAIRE
func Routes(r *mux.Router, s *service.Service) {
	r.Use(
		loggingMiddleware(s),
	)

	r.HandleFunc("/services", getAllServices(s)).Methods("GET")
	r.HandleFunc("/services/start/{name}", startService(s)).Methods("GET")
	r.HandleFunc("/services/stop/{name}/{instance}", stopService(s)).Methods("GET")
	r.HandleFunc("/services/restart/{name}/{instance}", restartService(s)).Methods("GET")
}

/*
######################################################################################################## @(°_°)@ #######
*/
