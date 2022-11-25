package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Site holds the schema definition for the Site entity.
type Site struct {
	ent.Schema
}

// Fields of the Site.
func (Site) Fields() []ent.Field {
	return []ent.Field{
		field.String("domain").Unique(),
		field.String("favicon").Optional(),
	}
}

// Edges of the Site.
func (Site) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("users", User.Type).
			Ref("sites"),
		edge.To("visits", Visit.Type),
	}
}
