package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Visit holds the schema definition for the Visit entity.
type Visit struct {
	ent.Schema
}

// Fields of the Visit.
func (Visit) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.String("referer"),
		field.Time("timestamp"),
		field.String("ip"),
	}
}

// Edges of the Visit.
func (Visit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("site", Site.Type).
			Ref("visits").
			Unique().
			Required(),
	}
}
