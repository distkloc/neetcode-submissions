func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sr, tr := []rune(s), []rune(t)
    m := map[rune]int{}

    for _, r := range sr {
        m[r] += 1
    }

    for _, r := range tr {
        m[r] -= 1
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}
