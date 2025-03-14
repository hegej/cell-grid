# 🧬 Cell-Grid

 **Cell-Grid**! is a fun and interactive Go application built using the [Raylib](https://www.raylib.com/) library. It simulates falling cells on a grid, with features like toggling cell states, applying gravity, and enabling a rainbow color mode. Perfect for experimenting with cellular physics and visualizations! 🎨

---

## 🚀 Features

- **Interactive Grid**: Click on any cell in the grid to toggle it on or off. 🖱️
- **Gravity Simulation**: Active cells fall towards the bottom of the screen, stacking on top of each other. 📉
- **Rainbow Mode**: Toggle rainbow colors for active cells by pressing the `R` key. 🌈
- **Smooth Physics**: Uses delta time for consistent falling behavior across different frame rates. ⏱️
- **Customizable**: Easily modify constants like `screenWidth`, `screenHeight`, `cellSize`, and `accelerationRate` to tweak the simulation. ⚙️

---

## 🛠️ Prerequisites

To run Cell-Grid, you need the following installed on your system:

- [Go](https://golang.org/dl/) (version 1.16 or later) 🐹
- [Raylib-Go](https://github.com/gen2brain/raylib-go) (Raylib bindings for Go) 🎮
- A C compiler (like `gcc` or `clang`) for compiling Raylib dependencies 🛠️

---

## 🎮 How to Use
Once Cell-Grid is running, you can interact with it as follows:
Left-Click: Toggle a cell on or off in the grid. 

R Key: Toggle rainbow mode on or off. When enabled, active cells get random colors; when disabled, all cells revert to white. 

Close Window: Press ESC or click the window's close button to exit. 

