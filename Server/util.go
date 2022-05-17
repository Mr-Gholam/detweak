package main

import "log"

func handleError(e error) {
	if e != nil {
		log.Println(e)
	}
}
