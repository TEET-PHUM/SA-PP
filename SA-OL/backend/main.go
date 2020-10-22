package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Teeth/app/controllers"
	"github.com/Teeth/app/ent"
	"github.com/Teeth/app/ent/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	USERNAME  string
	USEREMAIL string
}

type Mealplans struct {
	Mealplan []Mealplan
}

type Mealplan struct {
	MEALDATE      string
	FOODID        int
	MEALID        int
	MEALPLANOWNER int
}

type Foodmenus struct {
	Foodmenu []Foodmenu
}

type Foodmenu struct {
	FOODTIME      string
	FOODMENUNAME  string
	FOODMENUING   string
	CALORIES      int
	FOODTYPE      int
	FOODGROUP     int
	FOODMENUOWNER int
}

type Tastes struct {
	Taste []Taste
}

type Taste struct {
	TASTENAME string
}

// @title SUT SA Example API Eatinghistory
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewFoodmenuController(v1, client)
	controllers.NewTasteController(v1, client)
	controllers.NewMealplanController(v1, client)
	controllers.NewEatinghistoryController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"Chao Pramong", "salt@gmail.com"},
			User{"Moung Najae", "rainbow@hotmail.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetEmail(u.USEREMAIL).
			SetName(u.USERNAME).
			Save(context.Background())
	}

	// Set Taste Data
	tastes := Tastes{
		Taste: []Taste{
			Taste{"Delicious"},
			Taste{"Good"},
			Taste{"Normal"},
			Taste{"Bad"},
			Taste{"Can not eat"},
		},
	}

	for _, t := range tastes.Taste {
		client.Taste.
			Create().
			SetTaste(t.TASTENAME).
			Save(context.Background())
	}

	// Set Mealplan Data
	mealplans := Mealplans{
		Mealplan: []Mealplan{
			Mealplan{"2020-02-15", 1, 3, 1},
			Mealplan{"2020-02-16", 2, 2, 1},
			Mealplan{"2020-02-17", 3, 1, 1},
			Mealplan{"2020-02-18", 4, 5, 2},
			Mealplan{"2020-02-19", 5, 4, 2},
			Mealplan{"2020-02-20", 6, 6, 2},
		},
	}

	for _, m := range mealplans.Mealplan {

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(m.MEALPLANOWNER))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Mealplan.
			Create().
			SetMealID(m.MEALID).
			SetFoodID(m.FOODID).
			SetDate(m.MEALDATE).
			SetOwner(u).
			Save(context.Background())
	}

	// Set Foodmenu Data
	foodmenus := Foodmenus{
		Foodmenu: []Foodmenu{
			Foodmenu{"2020-02-09 15:44", "Bread", "flour wheat ", 230, 2, 3, 1},
			Foodmenu{"2020-02-10 15:44", "Cup Noodle", "flour wheat MSG water", 670, 3, 3, 1},
			Foodmenu{"2020-02-11 15:44", "Seafood", "chille onion shrimp dquid ", 1100, 1, 2, 1},
			Foodmenu{"2020-02-12 15:44", "rice", "wheat ", 100, 1, 1, 2},
			Foodmenu{"2020-02-13 15:44", "soup", "water meat vegetable ", 50, 1, 1, 2},
			Foodmenu{"2020-02-14 15:44", "Burger", "flour wheat meat vegetable", 5000, 4, 4, 2},
		},
	}

	for _, f := range foodmenus.Foodmenu {

		u, err := client.User.
			Query().
			Where(user.IDEQ(int(f.FOODMENUOWNER))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Foodmenu.
			Create().
			SetAddedtime(f.FOODTIME).
			SetName(f.FOODMENUNAME).
			SetMenuing(f.FOODMENUING).
			SetCalories(f.CALORIES).
			SetTypeid(f.FOODTYPE).
			SetGroupid(f.FOODGROUP).
			SetOwner(u).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
