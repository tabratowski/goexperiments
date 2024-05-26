package validation

type testStruct struct {
	name      string `name:"name"`
	surname   string
	someVal   int
	someFloat float64
	sub       *testSub
	arr       []int
	m         map[int]string
}

type testSub struct {
	subName string
}
