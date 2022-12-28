package schema

import "entgo.io/ent"

// Episode holds the schema definition for the Episode entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
	return nil
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return nil
}
