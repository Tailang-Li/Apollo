package leetcode

/*
The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

countAndSay(1) = "1"
countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.
To determine how you "say" a digit string, split it into the minimal number of groups so that each group is a contiguous section all of the same character. Then for each group, say the number of characters, then say the character. To convert the saying into a digit string, replace the counts with a number and concatenate every saying.
*/

func countAndSay(n int) string {
	result := "1"
	for i := 2; i <= n; i++ {

		var temp string
		cnt := 1
		for j := 1; j < len(result); j++ {
			if result[j] == result[j-1] {
				cnt++
			} else {
				temp += fmt.Sprintf("%d%c", cnt, result[j-1])
				cnt = 1
			}
		}

		temp += fmt.Sprintf("%d%c", cnt, result[len(result)-1])

		result = temp
	}

	return result
}
