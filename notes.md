# Go Foundations #
- Magesh Kuppan

## Methodology ##
- No powerpoints
- We discuss & code
- Repo : https://github.com/tkmagesh/Cisco-Go-Aug-2023

## Schedule ##
- Commence      : 9:00 AM
- Tea Break     : 10:40 AM (20 mins)
- Lunch Break   : 12:30 PM (1 hour)
- Tea Break     : 2:30 PM (20 mins)
- Wind up       : 5:00 PM

## Golang ##
- Why Golang?
    - Simplicity
    - Performance
        - On par with C++
    - Close to hardware
    - Lightweight concurrency model

## Software Requirements ##
- Go Tools (https://go.dev/dl)
- Visual Studio Code (https://code.visualstudio.com)
- Go Extension for VSCode (https://marketplace.visualstudio.com/items?itemName=golang.Go)

## Introduction ##
- Documentation
    - https://pkg.go.dev/std
- Run a go program
    - > go run <filename.go>
- Create a build
    - > go build <filename.go>
    - > go build -o <output_file> <filename.go>

- To get the list of environment variables used by go tool
    - > go env
- Environment variables for cross compilation
    - GOARCH
    - GOOS
- To get the list of supported platforms/process arch for cross compilation
    - > go tool dist list
- To cross compile
    ```
     GOOS=<target OS> GOARCH=<target processor arch> go build -o <output> <filename.go>
    ```    
    - > ex: GOOS=windows GOARCH=386 go build -o build/hw 01-hello-world.go

## Data Types ##
- bool
- string

- int
- int8
- int16
- int32
- int64

- uint
- uint8
- uint16
- uint32
- uint64

- float32
- float64

- complex64 (real [float32] + imaginary [float32])
- complex128 (real [float64] + imaginary [float64])

- byte (alias for uint8)
- rune (alias for int32)(unicode code point)

## Variables ##
- Function Scope
    - Cannot have unused variables
    - Can use :=
- Package Scope
    - Can have unused variables
    - Cannot use :=
