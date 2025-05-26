package exception

import "log"

func CatchUp() {
	if err := recover(); err != nil {
		log.Println("error internal,", err)
	}
}
