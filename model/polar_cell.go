package model

// TODO -- potentially find a way to ditch North, South, East, West, etc.
type polarCell struct {
	*BaseCell

	cw Cell
	ccw Cell
	inward Cell
	outward []Cell
}

func NewPolarCell(row int, column int) *polarCell {
	return &polarCell{
		NewBaseCell(row, column),
		nil,
		nil,
		nil,
		nil,
	}
}

func (p *polarCell) Neighbors() []Cell {
	output := make([]Cell, 0, 4);
	if p.cw != nil {
		output = append(output, p.cw);
	}
	if p.ccw != nil {
		output = append(output, p.ccw);
	}
	if p.inward != nil {
		output = append(output, p.inward);
	}
	output = append(output, p.outward...);
	return output;
}
