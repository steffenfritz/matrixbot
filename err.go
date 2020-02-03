package matrixbot

import "log"

func e(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
