package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (app *application) setupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           0,
	}))

	usersRoutes := r.Group("/users")
	{
		usersRoutes.POST("/register", app.registerUserHandler)
		usersRoutes.POST("/login", app.loginUserHandler)
		usersRoutes.Use(app.authenticate())
		usersRoutes.GET("/:email", app.getUserHandler)
		usersRoutes.PATCH("/strength", app.upgradeStrengthHandler)
		usersRoutes.PATCH("/dexterity", app.upgradeDexterityHandler)
		usersRoutes.PATCH("/constitution", app.upgradeConstitutionHandler)
		usersRoutes.PATCH("/intelligence", app.upgradeIntelligenceHandler)
	}

	itemsRoutes := r.Group("/items")
	{
		//should add something like an api key and middleware so only admins can add new items
		itemsRoutes.POST("/weapons", app.createWeaponHandler)
		itemsRoutes.POST("/accessories", app.createAccessoryHandler)
		itemsRoutes.POST("/armours", app.createArmourHandler)
		itemsRoutes.POST("/shields", app.createShieldHandler)
	}

	questRoutes := r.Group("/quests")
	{
		//should add something like an api key and middleware so only admins can add new quests
		questRoutes.POST("/create", app.createQuestHandler)
		questRoutes.Use(app.authenticate())
		questRoutes.GET("/generate", app.generateQuestsHandler)
		questRoutes.POST("/set", app.setCurrentQuestHandler)
	}

	r.Run(":8080")

	return r
}
