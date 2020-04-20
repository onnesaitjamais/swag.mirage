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

package mirage

import (
	"github.com/arnumina/swag"
	"github.com/arnumina/swag/service"
	"github.com/arnumina/swag/util/options"
	"github.com/gorilla/mux"

	"github.com/arnumina/swag.mirage/internal/api/systemd"
	v1 "github.com/arnumina/swag.mirage/internal/api/v1"
)

func initialize(r *mux.Router, s *service.Service) {
	systemd.Routes(r.PathPrefix("/api/systemd").Subrouter())
	v1.Routes(r.PathPrefix("/api/v1").Subrouter(), s)
}

// Run AFAIRE
func Run(version, builtAt string) error {
	router := mux.NewRouter()

	service, err := swag.NewService(
		"mirage",
		version,
		builtAt,
		swag.Server(
			"http",
			options.Options{
				"handler":    router,
				"health_URI": "/api/systemd/health",
			},
		),
	)
	if err != nil {
		return err
	}

	defer service.Close()

	initialize(router, service)

	if err := service.Run(); err != nil {
		return err
	}

	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
