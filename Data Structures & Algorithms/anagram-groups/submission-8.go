import (
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    res := make(map[string][]string)

    for _, s := range strs {
        r := []rune(s)
        slices.Sort(r)
        sorted := string(r)
        res[sorted] = append(res[sorted], s) 
    }

    result := [][]string{} 

    for _, group := range res {
        result = append(result, group)
    }
    return result
}
