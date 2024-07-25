package main

import "fmt"

var (
	name string = "Dima"
	age  int    = 23
)

/*
Напишите программу, которая последовательно делает следующие операции с введённым числом:

Число умножается на 2;
Затем к числу прибавляется 100.
*/
func stepic_1_5_11() {
	var a int
	fmt.Scan(&a)
	a *= 2
	a += 100
	fmt.Println(a)
}

/*
Петя торопился в школу и неправильно написал программу,
которая сначала находит квадраты двух чисел, а затем их суммирует. Исправьте его программу.
*/
func stepic_1_5_12() {

	var a int
	var b int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли

	a = a * a
	b = b * b
	c := a + b
	fmt.Println(c)
}

/*
По данному целому числу, найдите его квадрат.
*/

func stepic_1_5_13() {

	var b int
	fmt.Scan(&b)
	b = b * b
	fmt.Println(b)

}

//Дано натуральное число, выведите его последнюю цифру.

func stepic_1_5_14() {

	var b int
	fmt.Scan(&b)
	b %= 10
	fmt.Println(b)
}

//Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).

func stepic_1_5_15() {

	var b int
	fmt.Scan(&b)

	fmt.Println(b % 100 / 10)
}

//Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

func stepic_1_5_16() {

	var angle int
	fmt.Scan(&angle)

	fmt.Println("It is", angle/30, "hours", angle*2%60, "minutes.")
}

/*
На ввод подается целое число. Если число положительное - вывести сообщение
"Число положительное", если число отрицательное - "Число отрицательное".
Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.
*/

func stepic_1_9_5() {

	var a int
	fmt.Scan(&a)
	if a < 0 {
		fmt.Println("Число положительное")
	} else if a < 0 {
		fmt.Println("Число отрицательное")
	} else if a == 0 {
		fmt.Println("Ноль")
	}
}

// По данному трехзначному числу определите, все ли его цифры различны.
func stepic_1_9_6() {
	var a int
	fmt.Scan(&a)
	if a%10 != a%100/10 && a%10 != a/100 && a%100/10 != a/100 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

//Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

func stepic_1_9_7() {
	var a int
	fmt.Scan(&a)
	if a == 10000 {
		fmt.Println(a / 10000)
	} else if a < 10000 && a >= 1000 {
		fmt.Println(a / 1000)
	} else if a < 1000 && a >= 100 {
		fmt.Println(a / 100)
	} else if a < 100 && a >= 10 {
		fmt.Println(a / 10)
	} else if a < 10 {
		fmt.Println(a)
	}

}

func stepic_1_9_8() {

	var a string
	fmt.Scan(&a)

	if int((a[0]))+int(a[1])+int(a[2]) == int(a[3])+int(a[4])+int(a[5]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func stepic_1_9_8_v2() {
	var s string
	fmt.Scan(&s)
	if s[0]+s[1]+s[2] == s[3]+s[4]+s[5] {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}

/*
Требуется определить, является ли данный год високосным, напомним:
Год является високосным если он соответствует хотя бы одному из нижеперечисленных условий:
- кратен 400;
- кратен 4, но не кратен 100.
*/
func stepic_1_9_9() {
	var a int
	fmt.Scan(&a)
	if a%400 == 0 || (a%4 == 0 && a%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func stepic_1_10_2() {

	for a := 1; a <= 10; a++ {
		fmt.Println(a * a)
	}
}

func stepic_1_10_3() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	var sum = 0
	for ; a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}
