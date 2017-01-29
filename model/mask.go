package model

import (
	"maze/util"
	"os"
	"bufio"
	"fmt"
	"image/png"
	"image/color"
)

type mask struct {
	rows 	int
	columns int

	bits    [][]bool
}

// NOTE: Might be better served in the util package.
func readlines(filename string) []string {
	file, err := os.Open(filename);
	if err != nil {
		panic("Failed to open file.");
	}
	defer file.Close();

	lines := []string{};
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		fmt.Println(scanner.Text());
		lines = append(lines, scanner.Text());
	}

	if err := scanner.Err(); err != nil {
		panic("Something went wrong while we were scanning.");
	}

	return lines;
}

func MaskFromText(filename string) *mask {
	var lines []string = readlines(filename);

	var rows int = len(lines);
	var columns int = len(lines[0]); // yes, yes, this will panic if the file is empty. so sue me.
	var mask *mask = NewMask(rows, columns);

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if string(lines[i][j]) == "X" {
				mask.SetCell(i, j, false);
			} else {
				mask.SetCell(i, j, true);
			}
		}
	}

	return mask;
}

func MaskFromPNG(filename string) *mask {
	f, err := os.Open(filename);
	if err != nil {
		panic("PNG file does not exist.");
	}

	img, err := png.Decode(bufio.NewReader(f));
	if err != nil {
		panic("Could not parse PNG file.")
	}

	columns := img.Bounds().Max.X;
	rows := img.Bounds().Max.Y;

	mask := NewMask(rows, columns)
	for i := 0; i < mask.rows; i++ {
		for j := 0; j < mask.columns; j++ {
			black := color.RGBA{0, 0, 0, 0}
			r, g, b, a := img.At(j, i).RGBA()
			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			if c.R == black.R && c.G == black.G && c.B == black.B {
				mask.SetCell(i, j, false);
			} else {
				mask.SetCell(i, j, true);
			}
		}
	}
	return mask;
}

func NewMask(rows int, columns int) *mask {
	return &mask{
		rows: rows,
		columns: columns,

		bits: prepareBits(rows, columns),
	}
}

func prepareBits(rows int, columns int) [][]bool {
	bits := make([][]bool, rows);
	for i, _ := range(bits) {
		bits[i] = make([]bool, columns);
		for j, _ := range(bits[i]) {
			bits[i][j] = true;
		}
	}
	return bits;
}

func (m *mask) GetCell(row int, column int) bool {
	return m.bits[row][column];
}

func (m *mask) SetCell(row int, column int, val bool) {
	m.bits[row][column] = val;
}

func (m *mask) Count() int {
	var count int = 0;
	for i, _ := range(m.bits) {
		for j, _ := range(m.bits[i]) {
			if m.bits[i][j] {
				count++;
			}
		}
	}
	return count;
}

// Return a random cell that's true
func (m *mask) RandomCell() (row int, column int) {
	for {
		row = util.RANDOM.Intn(m.rows);
		column = util.RANDOM.Intn(m.columns);

		if m.bits[row][column] == true {
			return row, column;
		}
	}
}
