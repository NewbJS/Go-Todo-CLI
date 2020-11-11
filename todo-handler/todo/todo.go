package todo

// Todo is the main todo structure.
type Todo struct {
	Desc    string
	TimePub string

	// Only show timeEdited if edited == true.
	TimeEdited string
	Edited     bool
}
