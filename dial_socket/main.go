package main

type user struct {
	name  string
	email string
}

type user1 struct {
	name   string
	email  string
	email1 string
	email2 string
	email3 string
	email4 string
}

type IntType interface {
	test()
}

func (u user) test() {
	println("u in test", &u)
}

//go:noinline
func (u user1) test() {
	println("u1 in test", &u)
}

//go:noinline
func testIntType(i IntType) {
	u1 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	u2 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	u3 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	u4 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	u5 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	u6 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	u7 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	u8 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	println("V2", &u1, &u2, &u3, &u4, &u5, &u6, &u7, &u8, &i)
	i.test()
}

func main() {
	a := 1
	println(&a)
	test1(&a)
	println(&a)
	//createUserV1()
	//createUserV2()
	//u := user{}
	//println("base V2", &u)
	//testIntType(u)
	//u1 := user1{}
	//println("base V2 u1", &u1)
	//testIntType(u1)
	//u2 := user1{}
	//println("base V2 u2", &u2)
	//testIntType(u2)
	//createUserV1()
	//u1 := createUserV1()
	//u2 := createUserV1()
	//u3 := createUserV3()
	//_ = u3
	//u4 := createUserV3()
	//_ = u4
	//u3 := createUserV2()
	//u4 := createUserV2()
	//u5 := createUserV1()

	//println("u1", &u1, "u2", &u2)
}

//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)
	inside(u)
	return u
}

//go:noinline
func createUserV3() user {
	u2 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	_ = u2
	u1 := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	_ = u1
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V1", &u)
	println("V2", &u1)
	println("V3", &u2)
	return u
}

//go:noinline
func createUserV2() user {
	u := user{
		name:  "Bill111",
		email: "bill@ardanlabs.com111",
	}

	println("V2", &u)
	inside(u)
	return u
}

func test1(a *int) {
	println(&a)
}

//go:noinline
func inside(u IntType) {
	println("inside u:", &u)
	u.test()
}
