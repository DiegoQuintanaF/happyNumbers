package hn

type HappyNumber struct {
	CalculatedHappyNumbers map[int64]bool
}

func NewHappyNumber() *HappyNumber {
	return &HappyNumber{
		CalculatedHappyNumbers: make(map[int64]bool),
	}
}

func (h *HappyNumber) SquereDigits(n int64) int64 {
	copy := n
	var result int64
	// values := make(map[int64]int64)

	var digit int64

	for i := 0; copy != 0; i++ {
		digit = copy % 10
		result += digit * digit
		copy /= 10
	}

	return result
}

func (h *HappyNumber) IsHappyNumber(n int64) bool {
	values := make(map[int64]bool)
	copy := n

	for copy != 1 {
		if _, exist := values[copy]; exist {
			return false
		}

		values[copy] = true
		copy = h.SquereDigits(copy)

		if _, ok := h.CalculatedHappyNumbers[copy]; ok {
			h.CalculatedHappyNumbers[n] = true
			return true
		}
	}

	for key := range values {
		h.CalculatedHappyNumbers[key] = true
	}

	h.CalculatedHappyNumbers[n] = true
	return true
}

func (h *HappyNumber) CountHappyNumbers(a, b int64) int64 {
	var count int64

	for i := b; i >= a; i-- {
		if h.IsHappyNumber(i) {
			count++
		}
	}

	return count
}
