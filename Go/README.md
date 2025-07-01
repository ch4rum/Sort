# 🐹 Go - Sorting Algorithms

This Go implementation provides the same sorting algorithms: **Bubble**, **Selection**, **Insertion**, **Merge**, and **Quick Sort**.  
It reads data from a file, sorts it using the selected algorithm, and displays execution time.

## ⚙️ Requirements

This Go project depends on several external libraries, you'll need:
- **Go 1.20** installed - [Download Go]( https://go.dev/)
- `github.com/gookit/color` (v1.5.4) — for colored terminal output.
- `github.com/spf13/pflag` (v1.0.6) — for enhanced command-line flag parsing.

Install then with:
```bash
go get github.com/gookit/color
go get github.com/spf13/pflag
```

## 🚀 Usage & Build

- Clone the repository and install the dependencies, after run or build the program.

```sh
# Clone and navigate to folder
git clone https://github.com/ch4rum/Sort.git
cd Sort/Go

# Run directly
go run main.go [options]

# Or build and run
go build -o sortapp main.go
./sortapp [options]
