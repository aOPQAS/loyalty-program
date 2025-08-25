package server

func CalculatePriceAndDiscount(fixedPrice, totalCost, discountPercent int) (finalPrice int, finalDiscount int) {
	if fixedPrice > 0 {
		finalPrice = fixedPrice
		finalDiscount = 100 - (fixedPrice*100)/totalCost
	} else if discountPercent > 0 {
		finalDiscount = discountPercent
		finalPrice = totalCost - (totalCost*discountPercent)/100
	} else {
		finalPrice = totalCost
		finalDiscount = 0
	}
	return
}
