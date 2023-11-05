package widgets

import (
	"image"

	. "github.com/gizak/termui/v3"
)

type ScrollBox struct {
	Block
	ScrollYPosition int
	RowsAmount      int
	Text            string
	TextStyle       Style
	WrapText        bool
}

func NewScrollBox() *ScrollBox {
	return &ScrollBox{
		Block:     *NewBlock(),
		TextStyle: Theme.Paragraph.Text,
		WrapText:  true,
	}
}

func (self *ScrollBox) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	cells := ParseStyles(self.Text, self.TextStyle)
	if self.WrapText {
		cells = WrapCells(cells, uint(self.Inner.Dx()))
	}

	rows := SplitCells(cells, '\n')[self.ScrollYPosition:]

	for y, row := range rows {
		if y+self.Inner.Min.Y >= self.Inner.Max.Y {
			break
		}
		row = TrimCells(row, self.Inner.Dx())
		self.RowsAmount = len(row)
		for _, cx := range BuildCellWithXArray(row) {
			x, cell := cx.X, cx.Cell
			buf.SetCell(cell, image.Pt(x, y).Add(self.Inner.Min))
		}
	}
}

func (self *ScrollBox) ScrollAmount(amount int) {
	self.ScrollYPosition = self.ScrollYPosition + amount
}

func (self *ScrollBox) ScrollUp() {
	if self.ScrollYPosition > 0 {
		self.ScrollAmount(-1)
	}
}

func (self *ScrollBox) ScrollDown() {
	if self.ScrollYPosition < self.RowsAmount-(self.Inner.Max.Y-self.Inner.Min.Y) {
		self.ScrollAmount(1)
	}
}
