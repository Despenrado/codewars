package kui

func FindUniq(arr []float32) float32 {
	unique := make(map[float32]int)
	for _, v := range arr {
		unique[v]++
	}
	for k, v := range unique {
		if v == 1 {
			return k
		}
	}
	return arr[0]
}
