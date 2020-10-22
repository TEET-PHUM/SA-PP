package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Foodmenu holds the schema definition for the Foodmenu entity.
type Foodmenu struct {
	ent.Schema
}

// Fields of the Foodmenu.
func (Foodmenu) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("menuing").NotEmpty().Unique(),
		field.String("addedtime").NotEmpty().Unique(),
		field.Int("groupid").Positive().Unique(),
		field.Int("typeid").Positive().Unique(),
		field.Int("calories").Positive().Unique(),
	}
}

// Edges of the Foodmenu.
func (Foodmenu) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("foodmenus").Unique(),
		edge.To("eatinghistorys", Eatinghistory.Type).StorageKey(edge.Column("foodmenu_id")),
	}
}
