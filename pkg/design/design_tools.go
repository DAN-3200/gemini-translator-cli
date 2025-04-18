// import with command .
package design

import (
	"fmt"
	"strings"
)

// DynamicLayout
func Div(elements ...string) string {
	var n_elements = len(elements)
	if n_elements == 0 {
		return ""
	}

	var layout strings.Builder
	var all_elements = make([]any, n_elements)

	for index, value := range elements {
		all_elements[index] = value

		layout.WriteString("%s")
		if index < n_elements-1 {
			layout.WriteString("\n")
		}
	}

	return fmt.Sprintf(layout.String(), all_elements...)
}
