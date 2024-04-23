package hitung

func HitungBonusTHR(gaji float64, posisi string) float64 {
	var totalTHR float64
	if posisi == "c-level" {
		totalTHR = gaji * 2.5
	} else if posisi == "manager" {
		totalTHR = gaji * 2.0
	} else if posisi == "senior" {
		totalTHR = gaji * 1.0
	} else {
		totalTHR = 0
	}
	return totalTHR
}
