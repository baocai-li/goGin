package main

import "GinHello/initRouter"

func main() {
	//router := initRouter.SetupRouterNoGroup()
	router := initRouter.SetupRouter()
	_ = router.Run()
}

