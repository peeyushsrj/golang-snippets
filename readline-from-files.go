// ReadLines from file & return array of strings in lines (without \n)
func ReadLinesFromFile(path string) ([]string, error) {
	var arr []string

	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return arr, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr = append(arr, strings.TrimSpace(scanner.Text()))
	}
	if scanner.Err() != nil {
		return arr, scanner.Err()
	}
	return arr, nil
}
