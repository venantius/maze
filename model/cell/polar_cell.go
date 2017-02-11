package cell

type PolarCell struct {
	*baseCell

	links	map[*PolarCell]bool
	CW      Cell
	CCW     Cell
	Inward  Cell
	Outward []Cell
}

func NewPolarCell(row int, column int) *PolarCell {
	return &PolarCell{
		&baseCell{
			row: row,
			column: column,
		},
		make(map[*PolarCell]bool),
		nil,
		nil,
		nil,
		[]Cell{},
	}
}

func (p *PolarCell) Neighbors() []Cell {
	output := []Cell{};
	if p.CW != nil {
		output = append(output, p.CW);
	}
	if p.CCW != nil {
		output = append(output, p.CCW);
	}
	if p.Inward != nil {
		output = append(output, p.Inward);
	}
	output = append(output, p.Outward...);
	return output;
}

func (p *PolarCell) IsLinked(other Cell) bool {
	_, exists := p.links[other.(*PolarCell)]
	return exists
}

func (p *PolarCell) Link(other Cell, bidi bool) {
	p.links[other.(*PolarCell)] = true;
	if (bidi == true) {
		other.Link(p, false);
	}
}

func (p *PolarCell) Unlink(other Cell, bidi bool) {
	delete(p.links, other.(*PolarCell))
	if (bidi == true) {
		other.Unlink(p, false);
	}
}

func (p *PolarCell) Links() []Cell {
	var keys []Cell = make([]Cell, 0, len(p.links))
	for k := range p.links {
		keys = append(keys, k)
	}
	return keys
}

func (p *PolarCell) Distances() *Distances {
	return calculateDistances(p);
}
