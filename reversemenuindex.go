package piscine

func ReverseMenuIndex(menu []string) []string {
	n := len(menu)
	for i := 0; i < n/2; i++ {
		menu[i], menu[n-1-i] = menu[n-1-i], menu[i]
	}
	return menu
}
