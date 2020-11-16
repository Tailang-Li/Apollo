package leetcode

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for i := 1; i < len(strs); i++ {
		temp := result
		result = ""
		for j := 0; j < len(strs[i]) && j < len(temp); j++ {
			if strs[i][j] != temp[j] {
				break
			}
			result += string(temp[j])
		}
		if result == "" {
			break
		}
	}

	return result
}
