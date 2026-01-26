package efw

import "fmt"

func (that *eframework) printBanner() {
	fmt.Println(`
    ______ ____                            
   / ____// __/_____ ____ _ ____ ___   ___ 
  / __/  / /_ / ___// __  // __  __ \ / _ \
 / /___ / __// /   / /_/ // / / / / //  __/
/_____//_/  /_/    \__,_//_/ /_/ /_/ \___/ 

eframe v1.0.13 start`)
}
