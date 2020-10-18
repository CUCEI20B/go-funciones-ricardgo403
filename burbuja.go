package main

func Burbuja(mySlice []int64) {
	for i := 0; i < len(mySlice); i++ {
		for j := i + 1; j < len(mySlice); j++ {
			if mySlice[j] < mySlice[i] {
				temp := mySlice[i]
				mySlice[i] = mySlice[j]
				mySlice[j] = temp
			}
		}
	}
}

func main() {

}
