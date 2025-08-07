func encode(strs []string) string {
	var concatString strings.Builder

	for _, str := range strs {
		concatString.WriteString(strconv.Itoa(len(str)))
		concatString.WriteString("#")
		concatString.WriteString(str)
	}

	return concatString.String()
}

func decode(encodedStr string) []string {
	i := 0
	var result []string

	for i < len(encodedStr) {
		j := i
		for j < len(encodedStr) && encodedStr[j] != '#' {
			j++
		}

		length, err := strconv.Atoi(encodedStr[i:j])
		if err != nil {
			fmt.Println("Error decoding length:", err)
			return nil
		}

		i = j + 1
		if i+length > len(encodedStr) {
			fmt.Println("Error: not enough characters for the specified length")
			return nil
		}

		str := encodedStr[i : i+length]
		result = append(result, str)
		i += length
	}

	return result
}