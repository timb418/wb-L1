# wb-L1
Each directory represents a correlating task.

Tasks descriptioon:
[RUS]
1. Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской
структуры Human (аналог наследования).
2. Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
4. Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при старте. Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд
программа должна завершаться.
6. Реализовать все возможные способы остановки выполнения горутины.
7. Реализовать конкурентную запись данных в map.
8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
10. Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
11. Реализовать пересечение двух неупорядоченных множеств.
12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
13. Поменять местами два числа без создания временной переменной.
14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
15. К каким негативным последствиям может привести данный фрагмент кода, икак это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
}
func main() {
    someFunc()
}
16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
17. Реализовать бинарный поиск встроенными методами языка.
18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое
значение счетчика.
19. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
20. Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
21. Реализовать паттерн «адаптер» на любом примере.
22. Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
23. Удалить i-ый элемент из слайса.
24. Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
25. Реализовать собственную функцию sleep.
26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
Например:
abcd — true
abCdefAaf — false
aabcd — false

[ENG]
1. The Human structure is given (with an arbitrary set of fields and methods). Implement inlining methods in the Action structure from the parent
structures Human (analogue of inheritance).
2. Write a program that will competitively calculate the value of the squares of numbers taken from the array (2,4,6,8,10) and print their squares to stdout.
3. Given a sequence of numbers: 2,4,6,8,10. Find the sum of their squares (22+32+42….) using competitive calculations.
4. Implement permanent data recording to the channel (main stream). Implement a set of N workers that read arbitrary data from a channel and
output to stdout. You need to be able to select the number of workers at startup. The program should be terminated by pressing Ctrl+C. Select and justify
a way to terminate all workers.
5. Develop a program that will sequentially send values ​​to the channel, and read from the other side of the channel. After N seconds
the program should end.
6. Implement all possible ways to stop the execution of the goroutine.
7. Implement concurrent data entry in map.
8. An int64 variable is given. Write a program that sets the i-th bit to 1 or 0.
9. Develop a pipeline of numbers. Two channels are given: numbers (x) from the array are written to the first, the result of the x * 2 operation is written to the second, after which the data from the second channel should be output to stdout.
10. A sequence of temperature fluctuations is given: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Group these values ​​into groups in 10-degree increments. The sequence in the subsets is not important.
Example: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, ​​20: {24.5}, etc.
11. Implement the intersection of two unordered sets.
12. There is a sequence of strings - (cat, cat, dog, cat, tree) create its own set for it.
13. Swap two numbers without creating a temporary variable.
14. Develop a program that can determine the type of a variable at runtime: int, string, bool, channel from a variable of type interface{}.
15. What negative consequences can this code fragment lead to, and how can I fix it? Give a correct implementation example.
var justString string
func someFunc() {
    v := createHugeString(1 << 10)
    justString = v[:100]
}
func main() {
    someFunc()
}
16. Implement quick sorting of an array (quicksort) using the built-in methods of the language.
17. Implement binary search using the built-in methods of the language.
18. Implement a counter structure that will increment in a competitive environment. Upon completion, the program should output the final
counter value.
19. Develop a program that reverses the string given to the move (for example: "glavryba - abyrvalg"). Characters can be unicode.
20. Develop a program that reverses words in a line.
Example: "snow dog sun - sun dog snow".
21. Implement the "adapter" pattern on any example.
22. Develop a program that multiplies, divides, adds, subtracts two numerical variables a,b, whose value is > 2^20.
23. Remove the i-th element from the slice.
24. Develop a program for finding the distance between two points, which are represented as a Point structure with encapsulated parameters x, y and
constructor.
25. Implement your own sleep function.
26. Develop a program that checks that all characters in a string are unique (true if unique, false etc). The validation function must be case insensitive.
For example:
abcd - true
abCdefAaf-false
aabcd-false
