package schema

import "entgo.io/ent"

// Season holds the schema definition for the Season entity.
type Season struct {
	ent.Schema
}

// Fields of the Season.
func (Season) Fields() []ent.Field {
	return nil
}

// Edges of the Season.
func (Season) Edges() []ent.Edge {
	return nil
}
