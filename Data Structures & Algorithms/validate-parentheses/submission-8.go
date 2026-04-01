func isValid(s string) bool {
	sruned := []rune(s)
	paranStack := []rune{}

	for _,value := range sruned {
		if value == '(' || value == '{' || value == '[' {
			paranStack = append(paranStack,value)
		} else if value == ')' {
			if len(paranStack) == 0 {
				return false
			} 
			top := paranStack[len(paranStack)-1]
			if top == '(' {
				paranStack = paranStack[:len(paranStack)-1]
			}else {
				return false
			}
		} else if value == '}' {
			if len(paranStack) == 0 {
				return false
			} 
			top := paranStack[len(paranStack)-1]
			if top == '{' {
				paranStack = paranStack[:len(paranStack)-1]
			}else {
				return false
			} 
		}else if value == ']' {
			if len(paranStack) == 0 {
				return false
			} 
			top := paranStack[len(paranStack)-1]
			if top == '[' {
				paranStack = paranStack[:len(paranStack)-1]
			}else {
				return false
			}
		}
	}

	if len(paranStack) > 0 {
		return false
	} else {
		return true
	}


}
