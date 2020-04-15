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

package main

import (
	"os"

	"github.com/arnumina/swag.mirage/cmd/mirage"
)

var version, builtAt string

func main() {
	if mirage.Run(version, builtAt) != nil {
		os.Exit(-1)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
