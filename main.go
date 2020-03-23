package main

import (
	"time"

	"github.com/dtylman/gowd"
	"github.com/dtylman/gowd/bootstrap"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (c *computer) btnClicked(sender *gowd.Element, event *gowd.EventElement) {
	// adds a text to the body
	text := c.AddElement(gowd.NewStyledText("Running...", gowd.BoldText))

	// clean up - remove the added elements
	defer func() {
		text.SetText("Done.")
		c.Enable()
	}()

	// render the progress bar
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 500)
		c.div.RemoveElement(c.clkOff)
		c.div.AddElement(c.clkOn)
		c.Render()

		time.Sleep(time.Millisecond * 500)
		c.div.RemoveElement(c.clkOn)
		c.div.AddElement(c.clkOff)
		c.Render()
	}
}

type computer struct {
	*gowd.Element
	clkOn  *gowd.Element
	clkOff *gowd.Element
	div    *gowd.Element
}

func newComputer() *computer {
	c := &computer{}
	c.Element = gowd.NewElement("section")
	c.SetClass("computerapp")

	c.div = bootstrap.NewElement("div", "well")
	row := bootstrap.NewRow(bootstrap.NewColumn(bootstrap.ColumnMedium, 6, c.div))
	c.clkOn = gowd.NewElement(`img src="res/lamp_on.png" height="90" width="60" alt="Clock Pulse Indicator On" id="clk_img_on"`)
	c.clkOff = gowd.NewElement(`img src="res/lamp_off.png" height="90" width="60" alt="Clock Pulse Indicator Off" id="clk_img_off"`)

	c.div.AddElement(c.clkOff)
	c.AddElement(row)

	btn := bootstrap.NewButton(bootstrap.ButtonPrimary, "Start")
	btn.OnEvent(gowd.OnClick, c.btnClicked)
	c.AddElement(bootstrap.NewColumn(bootstrap.ColumnLarge, 4, btn))

	return c
}

func main() {
	gowd.Run(newComputer().Element)
}
