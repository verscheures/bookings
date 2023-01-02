package forms

type errors map[string][]string

// Add an error message for a field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)

}

// Get Returns the first error message for a field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
