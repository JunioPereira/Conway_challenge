# Conway's Game of Life in Go

This project implements [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) using the Go programming language. It simulates an infinite two-dimensional grid using 64-bit signed integers to support extremely large coordinate ranges.

Live cells are input using the **Life 1.06** file format. The program reads the input, runs the simulation for a number of generations, and outputs the final state to a new file. Each generation is also visualized in the terminal for easier debugging and observation.

---

## 🔧 Features

- Supports **arbitrary coordinates** using `int64` (effectively infinite grid)
- Uses a **sparse map-based representation** for performance
- Reads and writes using the **Life 1.06 format**
- **Terminal visualization** of generations using ASCII characters
- Clean and idiomatic Go implementation

---

## 📥 Input Format (`input.txt`)

```txt
#Life 1.06
1 0
1 1
1 2
