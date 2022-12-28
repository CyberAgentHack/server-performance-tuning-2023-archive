package schema

import "entgo.io/ent"

// Series holds the schema definition for the Series entity.
type Series struct {
	ent.Schema
}

// Fields of the Series.
func (Series) Fields() []ent.Field {
	return nil
}

// Edges of the Series.
func (Series) Edges() []ent.Edge {
	return nil
}
