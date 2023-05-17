package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)




func main() {


	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		fmt.Println(Parsing(&text))
	}


}








func Parsing(msg *string) string {

	*msg = strings.ToUpper(*msg)

	list_msg := strings.Split(*msg, " ")
	check_list_operator := []string{}

	roman_numeral := []string{}
	arabic_numeral := []string{} 
	operator := []string{}


	for i := 0;  i < len(list_msg); i++ {
		if len(list_msg[i]) > 0 {
			check_list_operator = append(check_list_operator, list_msg[i])

			Convert(&list_msg[i], &roman_numeral, &arabic_numeral, &operator) 
		}
		
	}

	check_list_operator = append(check_list_operator, "_") //чтобы можно было обралиться к [1]

	var res string

	if len(roman_numeral) == 2 && len(arabic_numeral) == 0 && len(operator) == 1 {
		roman := map[string]string {
			"I": 	"1",
			"II": 	"2",
			"III": 	"3",
			"IV": 	"4",
			"V": 	"5",
			"VI": 	"6",
			"VII": 	"7",
			"VIII": "8",
			"IX": 	"9",
			"X": 	"10",
		}

		type_number := "rom"
		res = Calculate(roman[roman_numeral[0]], roman[roman_numeral[1]], operator[0], type_number)

	} else if len(roman_numeral) == 0 && len(arabic_numeral) == 2 && len(operator) == 1 {
		type_number := "arb"
		res = Calculate(arabic_numeral[0], arabic_numeral[1], operator[0], type_number)


	} else if len(operator) > 1 {
		panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")


	} else if len(roman_numeral) == 1 || len(arabic_numeral) == 1 {
		panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
	
	
	} else {
		panic("Вывод ошибки, так как строка не является математической операцией.")
	}

	return res

}





func Convert(msg *string, roman_numeral *[] string, arabic_numeral *[] string, operator *[] string )  {

	if *msg == "-" {
		*msg = " - "
	}
	if strings.Contains("1 2 3 4 5 6 7 8 9 10 -1 -2 -3 -4 -5 -6 -7 -8 -9 -10", *msg) {
		*arabic_numeral = append(*arabic_numeral, *msg)

	} else if strings.Contains("I II III IV V VI VIII IX IX X", *msg) {
		*roman_numeral = append(*roman_numeral, *msg)

	} else if strings.Contains("+ - * /", *msg) {
		*operator = append(*operator, *msg)

	} else {
		panic("Вывод ошибки, так как строка не является математической операцией.")
	}

}




func ToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}





func Calculate(x string, y string, operator string, type_number string) string {

	var result int
	var result_str string

	switch operator {
	case "+":
		result = ToInt(x) + ToInt(y)
	case " - ":
		result = ToInt(x) - ToInt(y)
	case "*":
		result = ToInt(x) * ToInt(y)
	case "/":
		result = ToInt(x) / ToInt(y)
	}

	if type_number == "rom" {

		roman_numerals := map[int]string {
			1: "I",
			2: "II",
			3: "III",
			4: "IV",
			5: "V",
			6: "VI",
			7: "VII",
			8: "VIII",
			9: "IX",
			10: "X",
			20: "XX",
			30: "XXX",
			40: "XL",
			50: "L",
			60: "LX",
			70: "LXX",
			80: "LXXX",
			90: "XC",
			100: "C",
		}

		if result > 0 && result < 11 {
			result_str = roman_numerals[result]


		} else if result == 0 {
			panic("Вывод ошибки, так как в римской системе нет ноля.")


		} else if result == 100 {
			result_str = roman_numerals[result]
			
		} else if result > 10 && result < 100 {

			one_num_str :=   strconv.Itoa(result)[0:1] + "0" 	// первая цифра
			two_num_str :=   strconv.Itoa(result)[1:2]  		// вторая цифра  

			one_num_int, _ := strconv.Atoi(one_num_str)
			two_num_int, _ := strconv.Atoi(two_num_str)

			result_str = roman_numerals[one_num_int] + roman_numerals[two_num_int]

		} else if result < 0 {
			panic("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		}


	} else {

		result_str = fmt.Sprintln(result)
	}

	return result_str


}




















