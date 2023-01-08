package main

import "log"

/*
https://ru.wikipedia.org/wiki/Вентиль_Тоффоли

C = C XOR (A AND B)
*/

func Toffoli(){
	var a = 12
	var b = 32
	var c = 3

	var r =  c ^ (a & b)

	log.Println(c, r)
	
}


/*
– закон коммутативности
a ∨ b = b ∨ a; a ∧ b = b ∧ a; (6.4)

– закон ассоциативности
(a ∧ b) ∧ c = a ∧ (b ∧ c); (a ∨ b) ∨ c = a ∨ (b ∨ c); (6.5)

– закон поглощения
a ∧ (a ∨ b) = a; a ∨ (a ∧ b) = a; (6.6)

– закон дистрибутивности
a ∧ (b ∨ c) = (a ∧ b) ∨ (a ∧ z); a ∨ (b ∧ c) = (a ∨ b) ∧ (a ∨ c); (6.7)

– закон противоречия
a ∧ a = 0; (6.8)

– закон исключения третьего
a ∨ a = 1; (6.9)
a → b = a ∨ b; a ∼ b = (a ∧ b) ∨ (a ∧ b) (6.10)

*/