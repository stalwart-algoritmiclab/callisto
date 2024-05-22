package filter

const (
	// queryCount - query for counting rows.
	queryCount = "count(*)"
)

// Sort directions enum.
const (
	DirectionAscending  SortDirection = "ASC"
	DirectionDescending SortDirection = "DESC"
)

// directions - map for checking valid direction values.
var directions = map[SortDirection]struct{}{
	DirectionAscending:  {},
	DirectionDescending: {},
}

// Conditions enum.
const (
	ConditionAND Condition = "AND"
	ConditionOR  Condition = "OR"
	NoCondition  Condition = ""
)

// Condition - map for checking valid condition values.
var conditions = map[Condition]struct{}{
	ConditionAND: {},
	ConditionOR:  {},
	NoCondition:  {},
}

type (
	// SortDirection defines direction to sort.
	SortDirection string

	// Condition defines if statements combined with AND or OR operator.
	Condition string

	// searcher defines fields to search for.
	searcher struct {
		fields  []string
		pattern string
	}
)

// String returns string representation of condition.
func (c Condition) String() string {
	return string(c)
}

// Validate sort direction.
func (s SortDirection) Validate() bool {
	_, ok := directions[s]
	return ok
}

// Validate condition field.
func (c Condition) Validate() bool {
	_, ok := conditions[c]
	return ok
}
