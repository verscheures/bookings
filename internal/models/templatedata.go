package models

import (
	"github.com/verscheures/bookings/internal/forms"
)

// TempolateDate holds data from hamdlers to template
type TemplateData struct {
	StringMap 	map[string]string
	IntMap    	map[string]int
	FloatMap  	map[string]float32
	Data      	map[string]interface{}
	CSRFToken 	string
	Flash     	string
	Warning   	string
	Error     	string
	Form		*forms.Form
}
