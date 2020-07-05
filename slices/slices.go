package slices

/*
Given a slice made up of Strings, verify that a specific
string is present in the slice.
*/
func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}
