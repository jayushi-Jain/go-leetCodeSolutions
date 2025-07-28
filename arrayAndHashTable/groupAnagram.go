func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, str := range strs {
		runeStr := []rune(str) //this will convert the string to rune type i.e. [97 100 101]

		sort.Slice(runeStr, func(i, j int) bool { //sort is in-built function which provides with various functions
			return runeStr[i] < runeStr[j]   //This sort.slice will sort all the words in ascending order.
		})

		sortedStr := string(runeStr) //convert from rune to string again eat=>aet
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str) 
		//this will create array like anagramMap["aet"] = ["eat", "tea", "ate"]
	}

	result := make([][]string, 0,len(anagramMap))
	//create (Slice of string slices, start with 0 elemts, capacity will be uptill length of anagram)

	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}