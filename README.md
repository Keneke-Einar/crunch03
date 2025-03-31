# Conway's Game of Life

This is a Go implementation of Conway's Game of Life, a cellular automaton devised by mathematician John Conway. The program simulates the evolution of a grid of cells following specific rules, with various customization options through command-line flags.

## Features

- Implements core Game of Life rules
- Supports multiple input methods (manual, file, random generation)
- Customizable display with optional colors and footprints
- Adjustable grid size and animation speed
- Terminal-aware fullscreen mode
- Portal edges option

## Prerequisites

- Go 1.16 or higher
- A terminal that supports ANSI escape codes (for colored output)

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd <repository-directory>
```

2. Run the program:
```bash
go run main.go [options]
```

## Usage

The program accepts a grid of '#' (live cells) and '.' (dead cells) through:

- manual input,
- file input,
- random generation, or
- built-in template.

### Basic Usage
```bash
go run main.go
```
Then enter dimensions (height width) followed by the grid rows:
```
4 4
....
.#..
.##.
....
```

### Command Line Flags
```
--help                    Show usage information
--verbose                 Display simulation details (tick, size, live cells, delay)
--delay-ms=NUM            Set animation delay in milliseconds (default: 2500)
--random=WxH              Generate random grid of width W and height H
--file=FILENAME           Load grid from specified file
--fullscreen              Adjust grid to terminal size
--footprints              Show traces of previously live cells (displayed as '∘')
--colored                 Add color to live cells (cyan) and footprints (yellow)
--edges-portal            Enable portal edges (cells wrap around to opposite side)
--template=TEMPLATE       Load one of the existing templates
--use-unicode             Use unicode characters to display the cells
```

### Input Rules

- Minimum grid size: 3x3
- Only accepts '#' (live) and '.' (dead) characters
- All rows must match specified width

## Game Rules

- Live cell with < 2 live neighbors dies (underpopulation)
- Live cell with 2-3 live neighbors survives
- Live cell with > 3 live neighbors dies (overpopulation)
- Dead cell with exactly 3 live neighbors becomes alive (reproduction)
- Simulation runs until no live cells remain

## Display

- × represents live cells
- · represents dead cells
- ∘ represents footprints (when enabled)

### Unicode Characters

If the `--use-unicode` option is used, then the following characters are used:

- ⬛ (U+2B1B) for the live cells
- ⬜ (U+2B1C) for the dead cells
- 🟨 (U+1F7E8) for the footprints

## Examples

### Basic Input
```bash
go run main.go
4 4
....
.#..
.##.
....
```
Initial output:
```
· · · ·
· × · ·
· × × ·
· · · ·
```

### Verbose Mode
```bash
go run main.go --verbose
6 6
......
..##..
.##...
..##..
..##..
......
```
Output:
```
Tick: 1
Grid Size: 6x6
Live Cells: 8
DelayMs: 2500ms

· · · · · ·
· · × × · ·
· × × · · ·
· · × × · ·
· · × × · ·
· · · · · ·
```

### Random Grid with Unicode Characters
```bash
go run main.go --random=5x5 --use-unicode
```
Possible output:

```
⬜⬜⬛⬜⬜
⬛⬜⬛⬜⬜
⬜⬛⬛⬜⬜
⬛⬜⬜⬜⬜
⬜⬜⬜⬛⬛
```

### File Input

Use the `gosper-glider-gun.txt` file:

Run:

```bash
go run main.go --file=gosper-glider-gun.txt --use-unicode --verbose
```

Output:

```
Tick: 1
Grid Size: 38x20
Live Cells: 36
DelayMs: 2500 ms

⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜
⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬛⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬛⬜⬛⬛⬜⬜⬜⬜⬛⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬜⬜⬜⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬛⬛⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜⬜
```