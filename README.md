# Conway's Game of Life in Go

This project implements [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life) using the Go programming language. It simulates an infinite two-dimensional grid using 64-bit signed integers (`int64`) to support extremely large coordinate ranges.

Live cells are input using the **Life 1.06** file format. The program reads the input, runs the simulation for 10 generations, and writes the final state to an output file. Each generation is also visualized in the terminal for easier debugging and observation.

---

## 🧠 Approach & Thought Process

- **Sparse grid representation**: Instead of allocating memory for a huge or infinite 2D grid, I used a `map[Coord]bool` to store only the live cells. This is both memory-efficient and performant, especially given that the grid can contain extremely large values in the 64-bit space.

- **Neighbor tracking**: For each live cell, I incremented the neighbor count for all 8 surrounding positions. This allows us to determine which cells might become alive or remain alive based on the Game of Life rules.

- **Rule application**: I applied the standard rules of Conway's Game of Life:
  - A live cell stays alive if it has 2 or 3 neighbors.
  - A dead cell becomes alive if it has exactly 3 neighbors.

- **Separation of concerns**: I structured the code so that reading input, stepping generations, and writing output are clearly separated. This makes testing and debugging easier.

---

## 🧪 How I Tested It

1. **Visual inspection**: I printed each generation to the terminal within a small fixed window (`-5 to 5`) to observe known patterns like gliders or blinkers and verify their expected movement or transformation.

2. **Edge case testing**: I tested with:
   - Sparse patterns with extremely large positive and negative coordinates (e.g., `-2000000000000 -2000000000000`)
   - Dense patterns close to origin to check neighbor counting accuracy

3. **File-based I/O**: I manually verified the correctness of the final output (`output.txt`) by comparing it with expected states for known patterns.

4. **Pattern reference**: I used the initial glider pattern from the Game of Life wiki to confirm that the simulation worked as expected across generations.

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
0 1
1 2
2 0
2 1
2 2
-2000000000000 -2000000000000
-2000000000001 -2000000000001
-2000000000000 -2000000000001
