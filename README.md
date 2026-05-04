# Basics of Go

A professional, beginner-friendly guide to learning the fundamental concepts of the Go programming language. This repository covers variables, data types, functions, and string formatting through practical examples.

## 🚀 Getting Started

### Prerequisites
- [Go](https://go.dev/doc/install) installed on your machine.

### Running the Code
To run the main application and see the output, execute the following command in your terminal:
```bash
go run main.go
```

### 🖥️ Expected Output
After running the command, you should see the following output in your terminal:
```text
This is my coffee Making black coffee coffee with sugar...... 
 and this is your bill 25 
This is my coffee Making Latte coffee without sugar...... 
 and this is your bill 20 
```

---

## 📚 Key Concepts Covered

### 1. Variables & Declarations
Go offers several ways to declare variables. This project demonstrates:
- **Short Hand Declaration (`:=`)**: The most common way to declare and initialize variables inside functions.
- **`var` Keyword**: Used for explicit type declaration or when declaring variables outside of functions.
- **Grouped Variables**: A clean way to declare multiple variables at once.
- **Constants**: Variables whose values cannot be changed after declaration.

### 2. Data Types & Zero Values
Understanding how Go handles data types is crucial:
- **Core Types**: `string`, `int`, `float64`, and `bool`.
- **Zero Values**: In Go, variables declared without an initial value are given a default "zero value":
  - `int`: `0`
  - `string`: `""` (empty string)
  - `bool`: `false`
  - `float`: `0`

### 3. Functions
The project showcases the flexibility of functions in Go:
- **Single Return**: Returning a single value.
- **Multiple Returns**: Functions can return more than one value, often used for returning a result along with an error.
- **Named Returns**: You can name the return variables in the function signature, allowing for cleaner code and "naked" returns.

### 4. String Formatting
Examples of how to format strings using the `fmt` package:
- `fmt.Printf`: Printing formatted text to the console.
- `fmt.Sprintf`: Returning a formatted string for later use.

---

## 🛠️ Code Breakdown

### Coffee Maker Example
The project uses a practical `makeCoffe` function to demonstrate:
- Function parameters (`kind`, `isSugar`).
- Conditional logic (`if/else`).
- Returning both a status message and a price.

```go
func makeCoffe (kind string, isSugar bool) (coffee string, price int) {
    // Logic for making coffee and calculating bill
}
```

## 👨‍💻 Learning Notes
- Use `:=` for quick declarations within functions.
- Remember that Go is statically typed; once a variable's type is set, it cannot change.
- Zero values ensure that your variables are never uninitialized.

---
*Happy Coding!*
