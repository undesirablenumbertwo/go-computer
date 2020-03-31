package main

import (
	"strconv"
	"time"

	"github.com/dtylman/gowd"
	"github.com/dtylman/gowd/bootstrap"
)

func (c *computer) clkSpdChnge(sender *gowd.Element, event *gowd.EventElement) {
	c.clkMdl.inputRow.astablSldrLbl.SetText(sender.GetValue())
}

func (c *computer) clkInput(sender *gowd.Element, event *gowd.EventElement) {
	i := 0
	for i = 0; i < len(c.clkMdl.inputRow.astablChkBx.Attributes); i++ {
		if c.clkMdl.inputRow.astablChkBx.Attributes[i].Val == "checked" {
			break
		}
	}
	if i < len(c.clkMdl.inputRow.astablChkBx.Attributes) {
		c.clkMdl.inputRow.monstablChkBx.SetValue("false")
	}
}

func (c *computer) startClk(sender *gowd.Element, event *gowd.EventElement) {
	// adds a text to the body
	text := c.AddElement(gowd.NewStyledText("Running...", gowd.BoldText))

	// clean up - remove the added elements
	defer func() {
		text.SetText("Done.")
		c.Enable()
	}()

	// render the progress bar
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * c.clkMdl.clkDelay)
		c.clkMdl.outputRow.RemoveElement(c.clkMdl.clkOff)
		c.clkMdl.outputRow.AddElement(c.clkMdl.clkOn)
		c.Render()

		time.Sleep(time.Millisecond * c.clkMdl.clkDelay)
		c.clkMdl.outputRow.RemoveElement(c.clkMdl.clkOn)
		c.clkMdl.outputRow.AddElement(c.clkMdl.clkOff)
		c.Render()
	}
}

type clockInputModule struct {
	*gowd.Element

	astablChkBx   *gowd.Element
	monstablChkBx *gowd.Element
	choiceRow     *gowd.Element

	astablSldr    *gowd.Element
	astablSldrLbl *gowd.Element
	ctrlRow       *gowd.Element
}

type clockModule struct {
	*gowd.Element

	clkRunning *gowd.Element
	clkOn      *gowd.Element
	clkOff     *gowd.Element
	clkMode    int
	clkDelay   time.Duration

	outputRow *gowd.Element
	inputRow  clockInputModule
}

type computer struct {
	*gowd.Element
	clkMdl clockModule
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func newComputer() *computer {
	//Initialise computer elements
	c := &computer{}
	c.Element = gowd.NewElement("section")
	c.SetClass("computerapp")
	c.clkMdl.clkDelay = 500
	c.clkMdl.clkOn = gowd.NewElement(`img src="res/lamp_on.png" height="90" width="60" alt="Clock Pulse Indicator On" id="clk_img_on"`)
	c.clkMdl.clkOff = gowd.NewElement(`img src="res/lamp_off.png" height="90" width="60" alt="Clock Pulse Indicator Off" id="clk_img_off"`)
	c.clkMdl.Element = bootstrap.NewElement("div", "well")

	//Output row shows clock module pulse outputs
	c.clkMdl.outputRow = bootstrap.NewRow(c.clkMdl.clkOff)

	//Input row allows choice of clock module & control of parameters i.e. speed
	c.clkMdl.inputRow.astablSldr = bootstrap.NewColumn(bootstrap.ColumnSmall, 4)
	c.clkMdl.inputRow.astablSldr.AddHTML(`<input type="range" min="1" max="1000" value="500" class="slider" id="astable_sldr">`, nil)
	c.clkMdl.inputRow.astablSldrLbl = gowd.NewElement("label")
	c.clkMdl.inputRow.astablSldrLbl.SetText(strconv.Itoa(500))
	c.clkMdl.inputRow.astablSldr.AddElement(c.clkMdl.inputRow.astablSldrLbl)
	c.clkMdl.inputRow.astablSldr.OnEvent(gowd.OnChange, c.clkSpdChnge)

	//Choose clock pulse method: Astable = blinking with speed slider, Monostable = single pulse
	c.clkMdl.inputRow.astablChkBx = bootstrap.NewColumn(bootstrap.ColumnSmall, 4)
	c.clkMdl.inputRow.astablChkBx.AddHTML(`<input type="checkbox" name="astableChk" id="astable_chk"><label for="astable_chk">Astable</label><br>`, nil)
	c.clkMdl.inputRow.astablChkBx.OnEvent(gowd.OnClick, c.clkInput)

	c.clkMdl.inputRow.monstablChkBx = bootstrap.NewColumn(bootstrap.ColumnSmall, 4)
	c.clkMdl.inputRow.monstablChkBx.AddHTML(`<input type="checkbox" name="monostableChk" id="monostable_chk"><label for="monostable_chk">Monostable</label><br>`, nil)
	c.clkMdl.inputRow.monstablChkBx.OnEvent(gowd.OnClick, c.clkInput)

	c.clkMdl.inputRow.ctrlRow = bootstrap.NewRow(c.clkMdl.inputRow.astablSldr)
	c.clkMdl.inputRow.choiceRow = bootstrap.NewRow(c.clkMdl.inputRow.astablChkBx, c.clkMdl.inputRow.monstablChkBx)

	c.clkMdl.AddElement(c.clkMdl.outputRow)
	c.clkMdl.AddElement(c.clkMdl.inputRow.ctrlRow)
	c.clkMdl.AddElement(c.clkMdl.inputRow.choiceRow)
	c.AddElement(c.clkMdl.Element)

	btn := bootstrap.NewButton(bootstrap.ButtonPrimary, "Start")
	btn.OnEvent(gowd.OnClick, c.startClk)
	c.AddElement(bootstrap.NewColumn(bootstrap.ColumnLarge, 4, btn))

	return c
}

func main() {
	gowd.Run(newComputer().Element)
}
