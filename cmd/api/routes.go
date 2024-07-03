package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO: Add something like an api key and middleware so only admins can add new items/quests etc.
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
		usersRoutes.GET("/:email", app.getUserHandler)
		usersRoutes.Use(app.authenticate())
		usersRoutes.PATCH("/strength", app.upgradeStrengthHandler)
		usersRoutes.PATCH("/dexterity", app.upgradeDexterityHandler)
		usersRoutes.PATCH("/constitution", app.upgradeConstitutionHandler)
		usersRoutes.PATCH("/intelligence", app.upgradeIntelligenceHandler)
		usersRoutes.POST("/shops/weapon", app.generateWeaponShop)
		usersRoutes.POST("/shops/magic", app.generateMagicShop)
		usersRoutes.PATCH("/shops/weapon/:slotKey", app.buyWeaponShopItem)
		usersRoutes.PATCH("/shops/magic/:slotKey", app.buyMagicShopItem)
		usersRoutes.PATCH("/equip/:slotKey", app.equipItem)
		usersRoutes.PATCH("/sell/:slotKey", app.sellItem)
		usersRoutes.POST("/fight/:email", app.fightPlayerHandler)
	}

	itemsRoutes := r.Group("/items")
	{
		itemsRoutes.POST("/weapons", app.createWeaponHandler)
		itemsRoutes.POST("/accessories", app.createAccessoryHandler)
		itemsRoutes.POST("/armours", app.createArmourHandler)
		itemsRoutes.POST("/shields", app.createShieldHandler)
	}

	questRoutes := r.Group("/quests")
	{
		questRoutes.POST("/create", app.createQuestHandler)
		questRoutes.Use(app.authenticate())
		questRoutes.GET("/generate", app.generateQuestsHandler)
		questRoutes.POST("/set", app.setCurrentQuestHandler)
		questRoutes.GET("/cancel", app.cancelCurrentQuestHandler)
		questRoutes.GET("/collect", app.collectCurrentQuestRewardsHandler)
		questRoutes.GET("/reset", app.resetQuestsHandler)
	}

	workRoutes := r.Group("/work")
	{
		workRoutes.Use(app.authenticate())
		workRoutes.POST("/set", app.setWorkHandler)
		workRoutes.GET("/cancel", app.cancelWorkHandler)
		workRoutes.GET("/collect", app.collectWorkRewardsHandler)
	}

	monsterRoutes := r.Group("/monsters")
	{
		monsterRoutes.POST("/", app.createMonsterHandler)
	}

	bossRoutes := r.Group("/bosses")
	{
		bossRoutes.POST("/", app.createBossHandler)
		bossRoutes.GET("/:position", app.getBossByPositionHandler)
	}

	dungeonRoutes := r.Group("/dungeon")
	{
		dungeonRoutes.Use(app.authenticate())
		dungeonRoutes.PATCH("/", app.fightDungeonBossHandler)
	}

	r.Run(":8080")

	return r
}
