package piscine

type food struct {
	food     string
	preptime int
}

func FoodDeliveryTime(order string) int {
	var menu []food
	burger := food{"burger", 15}
	chips := food{"chips", 10}
	nuggets := food{"nuggets", 12}

	menu = append(menu, burger, chips, nuggets)

	sum := 0

	for _, f := range menu {
		if f.food == order {
			sum += f.preptime
		}
	}

	return sum
}
