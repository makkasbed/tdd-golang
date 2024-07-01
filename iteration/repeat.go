package iteration

//Takes a string and returns the num of times
func Repeat(str string, num int) string {
	var results string
	for i := 0; i < num; i++ {
		results = results + str
	}
	return results
}
