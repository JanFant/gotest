package request

import (
	"fmt"
	"time"

	"github.com/jroimartin/gocui"
	component "github.com/skanehira/gocui-component"
)

type signup struct {
	*component.Form
}

var i int

func Guimain() {
	gui, err := gocui.NewGui(gocui.Output256)

	if err != nil {
		panic(err)
	}
	defer gui.Close()

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	// new form
	signup := &signup{
		component.NewForm(gui, "Vchs and Vas", 0, 0, 0, 0),
	}
	dev := signup.AddSelect("Устройство:", 20, 20)
	dev.AddOptions("vchs", "vas")
	go counter(signup)

	signup.AddButton("count+", signup.countA)
	maxX, maxY := signup.Size()
	v, _ := signup.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2)
	fmt.Fprintf(v, "%v", i)

	// dev := signup.AddSelect("Устройство:", 20, 20)
	// dev.AddOptions("vchs", "vas")
	// // dev.AddHandler(gocui.KeyEnter, signup.addSelector)

	// signup.AddSelect("Подсистема:", 20, 20)

	// // add input field
	// signup.AddInputField("First Name", 20, 18).
	// 	AddValidate("required input", requireValidator)
	// signup.AddInputField("Last Name", 20, 18).
	// 	AddValidate("required input", requireValidator)

	// signup.AddInputField("Password", 20, 18).
	// 	AddValidate("required input", requireValidator).
	// 	SetMask().
	// 	SetMaskKeybinding(gocui.KeyCtrlA)

	// // add checkbox
	// // signup.AddCheckBox("Age 18+", 11)

	// // add select
	// a := signup.AddSelect("Language", 20, 10)

	// for _, str := range vchsdev {
	// 	a.AddOptions(str)
	// }

	// // add radios
	// signup.AddRadio("Country", 20).
	// 	SetMode(component.VerticalMode).
	// 	AddOptions("Japan", "America", "China")

	// add button
	signup.AddButton("Regist", signup.regist)
	signup.AddButton("Cancel", quit)

	signup.Draw()
	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func (s *signup) redraw(g *gocui.Gui, v *gocui.View) error {
	time.Sleep(time.Second)
	s.Draw()
	return nil
}
func (s *signup) countA(g *gocui.Gui, v *gocui.View) error {
	i++
	return nil
}

func (s *signup) myForm(g *gocui.Gui, v *gocui.View) error {

	// dev.AddHandler(gocui.KeyEnter, signup.addSelector)

	s.AddSelect("Подсистема:", 20, 20)
	s.Draw()
	return nil
}

func (s *signup) addSelector(g *gocui.Gui, v *gocui.View) error {
	return nil
}

func counter(s *signup) {
	for {
		time.Sleep(500 * time.Millisecond)
		// if {
		// 	s.Draw()
		// }
		s.Update(func(g *gocui.Gui) error {
			v, err := g.View("Устройство:")
			if err != nil {
				return err
			}
			v.Clear()
			fmt.Fprintln(v, "asd")
			return nil
		})
	}
}

// func (s *signup) upSelector(g *gocui.Gui, v *gocui.View) error {
// 	for dev, str := range s.GetFieldTexts() {
// 		if dev == "Устройство:" && str == "vchs" {

// 		} else if dev == "Устройство:" && str == "vas" {

// 		}
// 	}
// 	return nil
// }

func (s *signup) regist(g *gocui.Gui, v *gocui.View) error {
	if !s.Validate() {
		return nil
	}

	var text string

	for label, ftext := range s.GetFieldTexts() {
		text += fmt.Sprintf("%s: %s\n", label, ftext)
	}

	for label, state := range s.GetCheckBoxStates() {
		text += fmt.Sprintf("%s: %t\n", label, state)
	}

	for label, opt := range s.GetSelectedOpts() {
		text += fmt.Sprintf("%s: %s\n", label, opt)
	}

	text += fmt.Sprintf("radio: %s\n", s.GetSelectedRadios())

	modal := component.NewModal(g, 0, 0, 30).SetText(text)
	modal.AddButton("OK", gocui.KeyEnter, func(g *gocui.Gui, v *gocui.View) error {
		modal.Close()
		s.SetCurrentItem(s.GetCurrentItem())
		return nil
	})

	modal.Draw()

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func requireValidator(value string) bool {
	if value == "" {
		return false
	}
	return true
}
