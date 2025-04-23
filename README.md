# Conway's Game of Life — Go Implementation

Hi there! 👋

This is my implementation of Conway’s Game of Life in Go, designed to handle an *infinite* 2D grid by using 64-bit signed integers. The goal was to simulate the evolution of live and dead cells over time, even when the cells are spaced extremely far apart — think trillions of units apart.

---

## 🧠 How I Thought About the Problem

The biggest challenge here was supporting coordinates anywhere in the 64-bit integer space. Allocating a full 2D grid obviously wouldn't work — it would blow up memory instantly. So instead, I leaned on a **sparse representation** using a `map` where keys are coordinates and values are booleans representing whether the cell is alive.

I broke the problem down into a few key parts:

1. **Reading the input**: I read cells from a file in Life 1.06 format, skipping comments and parsing coordinates as `int64`s.
2. **Tracking neighbors**: For every live cell, I looped over its 8 neighbors and counted how many times each cell was “visited” by a neighbor.
3. **Applying the rules**: Using the classic Game of Life rules:
   - Alive cells with 2 or 3 neighbors stay alive
   - Dead cells with exactly 3 neighbors come alive
4. **Writing output**: After simulating 10 generations, I wrote the final set of live cells back out in the same Life 1.06 format.

---

## 🧪 How I Tested It

To make sure everything was working correctly:

- I printed out each generation in a small window (from `-5` to `5`) to visually confirm the patterns evolved correctly.
- I used a **glider pattern** to make sure it moved diagonally as expected — a great sanity check.
- I added cells with very large coordinates (like `-2000000000000`) to make sure nothing broke at scale.
- And finally, I manually compared the input and output files after 10 generations to confirm the simulation was evolving realistically.

---

## 💡 Features

- 🧮 Handles 64-bit signed coordinates (yes, really)
- 🧼 Clean and idiomatic Go
- 🧭 Fully decoupled input/output/logic for readability and reusability
- 📺 Visualized grid output in the terminal using ASCII (live cells show up as `■`, dead cells as `.`)

---

## 📥 Input Format

The program reads from `input.txt` in the [Life 1.06 format](https://conwaylife.com/wiki/Life_1.06):

```txt
#Life 1.06
0 1
1 2
2 0
2 1
2 2