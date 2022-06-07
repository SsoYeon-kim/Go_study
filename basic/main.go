package main

import (
	"fmt"
	//"strings"
	"github.com/ssoyeon/learngo/accounts"
	"github.com/ssoyeon/learngo/mydict"
)

/*


//파라미터의 타입, return의 타입을 정의해줘야함
// (a, b int)로 적어줘도 둘다 int형으로 인식함
//항상 return 값은 존재해야함
func multiply(a int, b int) int {
	return a * b
}

//여러개 return 가능
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//원하는만큼 argument 받기 (타입 앞에 ...)
//array로 반환됨
func repeatMe(words ...string) {
	fmt.Println(words)
}

//naked return : return값을 지정하지 않아도됨
//defer : func이 return을 반환한 후 실행
func nakedlenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done")
	//:= 다시 생성 , = 값 업데이트
	//위 함수에서 length int로 이미 생성을 해준것이기 때문에 = 사용
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

//go에는 for밖에 없음
//range : array에 loop를 적용할 수 있게 해줌
func superAdd(numbers ...int) int {
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	fmt.Println("ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ")

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("ㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡㅡ")

	//인덱스를 사용하고 싶지 않을때
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

//??
func canIDrink(age int) bool {
	//조건 확인 전에 변수 생성가능 (;뒤로 변수 사용 가능)
	//koreanAge := age + 2와 같음
	//아래와 같이 생성 : if 조건에만 사용하기 위해 생성했음을 알 수 있음
	if koreanAge := age + 2; koreanAge < 20 {
		return false
	}
	return true
}

func s_canIDrink(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 18:
		return false
	case 20:
		return true
	}
	return false

	/*
		switch {
		case age < 18:
			return false
		case age >20:
			return true
		}
		return false

		switch age{
		case 18:
			return false
		case 20:
			return true
		}
		return false
*/
//}

/*
func main() {
	//const : 값 변경할 수 없음
	const name string = "soyeon"

	//var : 값 변경 가능
	var sy string = "soyeon"
	sy = "kim"
	fmt.Println(sy)

	//함수 안에 이렇게 정의할 수 있음 (알아서 type을 찾아줌)
	//함수 밖에서는 위와 같이 정의해야만 쓸 수 있음
	//type은 바꿀 수 없음, 변수에만 적용 가능
	ksy := "kimsoyeon"
	fmt.Println(ksy)

	fmt.Println(multiply(3, 3))

	totalLen, upperName := lenAndUpper("soyeon")
	fmt.Println(totalLen, upperName)

	// _ : 변수를 무시하는 값
	lenght, _ := lenAndUpper("kim")
	fmt.Println(lenght)

	repeatMe("kim", "so", "yeon")
	l, u := nakedlenAndUpper("I'm soyeon")
	fmt.Println(l, u)

	sum := superAdd(1, 2, 3, 4)
	fmt.Println(sum)

	fmt.Println(canIDrink(16))
	fmt.Println(s_canIDrink(18))

}

*/

/*
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	a := 2
	b := &a
	a = 5
	*b = 20
	fmt.Println(&a, b)
	fmt.Println(*b)
	fmt.Println(a)

	//array : 크기 제한
	names := [3]string{"k", "s"}
	names[2] = "y"
	fmt.Println(names)

	//slice : 길이 제한 없음
	names_slice := []string{"k", "s"}
	//slice 자체를 수정하진 않음, 새로운 값이 추가된 slice를 return함
	names_slice = append(names_slice, "y")
	fmt.Println(names_slice)

	//map[키]값
	ksy := map[string]string{"name": "so", "age": "3"}

	for key, value := range ksy {
		fmt.Println(key, value)
	}

	favFood := []string{"k", "food"}
	//아래 두가지 형식을 섞어 쓸 순 없음
	//soyeon := person{"sso", 3, favFood}
	soyeon := person{name: "sso", age: 3, favFood: favFood}
	fmt.Println(soyeon)
}
*/

func main() {
	account := accounts.NewAccount("ksy")
	account.Deposit(10)

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(account.Balance(), account.Owner())

	//struct의 주소 : &{ksy 10} 이 나오는데
	// String method를 사용해 struct의 출력 형태를 바꿈
	fmt.Println(account)

	/*
		dictionary := mydict.Dictionary{}
		dictionary["hello"] = "hello"
		fmt.Println(dictionary)
	*/

	/*
		dictionary := mydict.Dictionary{"first": "First word"}
		definition, err := dictionary.Search("second")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(definition)
		}
	*/

	dictionary := mydict.Dictionary{}
	baseword := "hello"
	definition := "greeting"
	err2 := dictionary.Add(baseword, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
	hello, _ := dictionary.Search(baseword)
	fmt.Println("found", baseword, "definition:", hello)

	err3 := dictionary.Add(baseword, definition)
	if err3 != nil {
		fmt.Println(err3)
	}

	err4 := dictionary.Update(baseword, "Second")
	if err4 != nil {
		fmt.Println(err4)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)

	dictionary.Delete(baseword)
	word, err5 := dictionary.Search(baseword)
	if err5 != nil {
		fmt.Println(err5)
	} else {
		fmt.Println(word)
	}
}
