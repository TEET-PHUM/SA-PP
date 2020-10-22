// Code generated by entc, DO NOT EDIT.

package mealplan

const (
	// Label holds the string label denoting the mealplan type in the database.
	Label = "mealplan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldFoodID holds the string denoting the food_id field in the database.
	FieldFoodID = "food_id"
	// FieldMealID holds the string denoting the meal_id field in the database.
	FieldMealID = "meal_id"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeEatinghistorys holds the string denoting the eatinghistorys edge name in mutations.
	EdgeEatinghistorys = "eatinghistorys"

	// Table holds the table name of the mealplan in the database.
	Table = "mealplans"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "mealplans"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"
	// EatinghistorysTable is the table the holds the eatinghistorys relation/edge.
	EatinghistorysTable = "eatinghistories"
	// EatinghistorysInverseTable is the table name for the Eatinghistory entity.
	// It exists in this package in order to avoid circular dependency with the "eatinghistory" package.
	EatinghistorysInverseTable = "eatinghistories"
	// EatinghistorysColumn is the table column denoting the eatinghistorys relation/edge.
	EatinghistorysColumn = "mealplan_id"
)

// Columns holds all SQL columns for mealplan fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldFoodID,
	FieldMealID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Mealplan type.
var ForeignKeys = []string{
	"owner_id",
}

var (
	// DateValidator is a validator for the "date" field. It is called by the builders before save.
	DateValidator func(string) error
	// FoodIDValidator is a validator for the "food_id" field. It is called by the builders before save.
	FoodIDValidator func(int) error
	// MealIDValidator is a validator for the "meal_id" field. It is called by the builders before save.
	MealIDValidator func(int) error
)
