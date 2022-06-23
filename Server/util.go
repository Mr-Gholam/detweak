package main

import "log"

func handleError(e error) {
	if e != nil {
		log.Println(e)
	}
}
func reverseArray(arr []map[string]interface{}) []map[string]interface{} {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
