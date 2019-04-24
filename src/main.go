package main

import (
	. "./core"
)

func main() {
	router := New()
	router.InitAuth()

	router.LoadHTML()
	//router.LoadAssets()

	router.SetS(Index, GET)
	router.SetS(Upload, POST)

	router.Run("localhost:8089")
}
