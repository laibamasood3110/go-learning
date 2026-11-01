# Go Practice

My journey learning the Go programming language. This repo holds the exercises I write while learning, and I add to it as I go.

## Topics covered so far
Variables and types, printing and input, operators (including bitwise), if/else, switch, loops, range, arrays and slices, maps, functions, recursion, structs, methods, pointers, error handling, and interfaces.

## How to run
Each exercise is its own folder with its own `go.mod`.

    cd <folder-name>
    go run .

## Exercises

### Basics: conditions, switch and loops
- `simplegradecalculator.go`: takes marks, works out the percentage, and prints a grade with an if/else chain
- `ticket-pricing-system`: movie ticket prices with nested if/else and weekend logic
- `traffic-light`: a traffic light simulator using slices and switch
- `library-book-checker`: slice capacity practice and switch conditions with `&&` and `||`
- `loop-practice`: for loops that print even and odd numbers with `break` at 15, a loop that skips 5 with `continue`, and a `range` loop over a slice of fruits
- `range-practice`: ranges over a slice of favourite books and prints each book's index and name

### Maps and structs
- `mapspractice`: a map of employee names to salaries, looped over with `range` to print each key and value
- `maps-struct`: a map of employee structs with a nested address struct, using a `calculateavg` method and a `printemployee` function, and writing the updated copy back into the map
- `employee-struct`: a slice of employee structs; sums each employee's salaries and prints the monthly average, dividing by the real number of salary entries
- `struct-student`: a slice of student structs; sums marks with nested loops, stores each percentage in the struct, and prints it with `%.02f`
- `student-result-management-system`: totals, percentages, grades, strongest subject, class average and pass/fail counts, built on a 2D slice

### Recursion
- `fictorial`: recursion practice: a recursive Fibonacci function, plus a recursive power function (commented out)

### Methods and pointers
- `employee-management-system`: employees in a map with a nested address struct, plus `calculateavg` and `describe` methods
- `pointers-practice`: shows that a function taking a struct only changes a copy, while a pointer receiver changes the original
- `employe-management-system-pointers`: the same program rewritten with pointer receivers

### Error handling
- `error-handling-practice`: a `divide` function that returns a result and an error, and handles division by zero
- `error-handlinginemployeemanagementsystem`: the employee system rejecting empty, all-zero and negative salary lists
- `minibanksystem`: a bank program with deposit, withdraw and transfer, using error handling for problems like a missing account or an insufficient balance

### Interfaces
- `interface`: `dog`, `cat` and `cow` each have `speak()`, and one function accepts all of them
- `interface-practice`: employees, accounts and products in one slice of a `describer` interface, printed by one loop
- `interface-errorpractice`: a custom error type with an `Error()` method

## What I learned about interfaces
- A type satisfies an interface by having methods with the same name, inputs and return type, with no explicit declaration needed.
- Adding a new type needs no changes to the interface or the function that uses it.
- `error` is a built-in interface.

## About me
I'm a BS Computer Science student learning Go and building my fundamentals step by step.