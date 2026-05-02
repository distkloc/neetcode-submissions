func isAnagram(s string, t string) bool {
    sr := []rune(s)
    tr := []rune(t)

    sort.Slice(sr, func(i, j int) bool {
        return sr[i] < sr[j]
    })

    sort.Slice(tr, func(i, j int) bool {
        return tr[i] < tr[j]
    })

    return string(sr) == string(tr)
}
