package schema

import "entgo.io/ent"

// Kratos holds the schema definition for the Kratos entity.
type Kratos struct {
	ent.Schema
}

// Fields of the Kratos.
func (Kratos) Fields() []ent.Field {
	return nil
}

// Edges of the Kratos.
func (Kratos) Edges() []ent.Edge {
	return nil
}
