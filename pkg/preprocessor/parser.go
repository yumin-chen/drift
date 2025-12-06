package preprocessor

// TODO: Define the D2 AST struct.
type AST struct{}

// TODO: Define the MotionGraph struct, which represents the internal graph
// format from MotionPlatform.
type MotionGraph struct{}

// NormalizeGraph converts an internal MotionPlatform graph into a canonical D2 AST.
// The conversion must be deterministic, with stable sorting of nodes and edges.
func NormalizeGraph(graph *MotionGraph) (*AST, error) {
	// TODO: Implement the logic to convert the MotionGraph to a D2 AST.
	// This will involve iterating over nodes and edges, applying stable
	// sorting, and generating the corresponding D2 AST representation.
	return &AST{}, nil
}
