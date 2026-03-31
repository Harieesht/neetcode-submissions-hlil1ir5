func isPalindrome(s string) bool {
    runes := []rune(s)
    left := 0
    right := len(s) - 1

    for { 
        if left > right {
            break
        }
        if !isAlphaNum(runes[left]) {
            left++
            continue
        } 
        if !isAlphaNum(runes[right]) {
            right--
            continue
        }  

        if toLower(runes[left]) != toLower(runes[right]) {
            return false
        }
        left++
        right--
    }

    return true

}

func isAlphaNum(r rune) bool {
    return (r >= 'A' && r <= 'Z') ||
    (r >= 'a' && r <= 'z') ||
    (r >= '0' && r <= '9')
}

func toLower(r rune) rune {
    if r >= 'A' && r <= 'Z' {
        return r + 32
    }
    return r
}
