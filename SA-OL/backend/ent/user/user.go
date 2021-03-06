// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"

	// EdgeFoodmenus holds the string denoting the foodmenus edge name in mutations.
	EdgeFoodmenus = "foodmenus"
	// EdgeMealplans holds the string denoting the mealplans edge name in mutations.
	EdgeMealplans = "mealplans"
	// EdgeEatinghistorys holds the string denoting the eatinghistorys edge name in mutations.
	EdgeEatinghistorys = "eatinghistorys"

	// Table holds the table name of the user in the database.
	Table = "users"
	// FoodmenusTable is the table the holds the foodmenus relation/edge.
	FoodmenusTable = "foodmenus"
	// FoodmenusInverseTable is the table name for the Foodmenu entity.
	// It exists in this package in order to avoid circular dependency with the "foodmenu" package.
	FoodmenusInverseTable = "foodmenus"
	// FoodmenusColumn is the table column denoting the foodmenus relation/edge.
	FoodmenusColumn = "owner_id"
	// MealplansTable is the table the holds the mealplans relation/edge.
	MealplansTable = "mealplans"
	// MealplansInverseTable is the table name for the Mealplan entity.
	// It exists in this package in order to avoid circular dependency with the "mealplan" package.
	MealplansInverseTable = "mealplans"
	// MealplansColumn is the table column denoting the mealplans relation/edge.
	MealplansColumn = "owner_id"
	// EatinghistorysTable is the table the holds the eatinghistorys relation/edge.
	EatinghistorysTable = "eatinghistories"
	// EatinghistorysInverseTable is the table name for the Eatinghistory entity.
	// It exists in this package in order to avoid circular dependency with the "eatinghistory" package.
	EatinghistorysInverseTable = "eatinghistories"
	// EatinghistorysColumn is the table column denoting the eatinghistorys relation/edge.
	EatinghistorysColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldEmail,
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
)
