func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sr, tr := []rune(s), []rune(t)
    m := map[rune]int{}

    for i, _ := range sr {
        m[sr[i]] += 1
        m[tr[i]] -= 1
    }

    for _, v := range m {
        if v != 0 {
            return false
        }
    }

    return true
}
