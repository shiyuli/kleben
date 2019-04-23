package main

import ."./core"

func main() {
	router := New()
	router.InitAuth()

	//core.LoadHTMLGlob("./templates")
	//engine.LoadHTMLFiles("./templates/index.tmpl")
	//core.Static("/", "./templates")

	router.LoadAssets()

	router.SetS(Index, GET)
	//router.Set(Login, GET)

	router.Run("localhost:8089")
}
