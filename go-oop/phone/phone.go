package phone

import (
	"fmt"
	"strings"
	"time"
)

type (
	GSM struct {
		Model        string
		Manufacturer string
		Price        float64
		Owner        string
		battery      Battery
		display      Display
	}
	Battery struct {
		Model string
		Idle  float64
		Talk  float64
	}
	Display struct {
		Size   float64
		Colors float64
	}

	Call struct {
		Time     time.Time
		Number   string
		Duration int
	}
	CallHistory struct {
		CallList []Call
	}
)

func NewGSM(model string, manufacturer string, price float64, owner string, battery Battery, display Display) GSM {
	return GSM{
		Model:        model,
		Manufacturer: manufacturer,
		Price:        price,
		Owner:        owner,
		battery:      battery,
		display:      display,
	}
}
func NewBattery(model string, idle float64, talk float64) Battery {
	return Battery{
		Model: model,
		Idle:  idle,
		Talk:  talk,
	}
}
func NewDisplay(size float64, colors float64) Display {
	return Display{
		Size:   size,
		Colors: colors,
	}
}

func (g GSM) ToString() string {
	s := ""
	s += fmt.Sprintf("Model: %s \n", g.Model)
	s += fmt.Sprintf("Manufacturer: %s \n", g.Manufacturer)
	s += fmt.Sprintf("Price: %v \n", g.Price)
	s += fmt.Sprintf("Owner: %s \n", g.Owner)
	s += fmt.Sprintf("Battery: %v \n", g.battery.ToString())
	s += fmt.Sprintf("Display: %v \n", g.display.ToString())
	return s
}

func (d Display) ToString() string {
	s := ""
	s += fmt.Sprintf("---Size: %v \n", d.Size)
	s += fmt.Sprintf("---Colors: %v \n", d.Colors)
	return s
}
func (b Battery) ToString() string {
	s := ""
	s += fmt.Sprintf("---Model: %v \n", b.Model)
	s += fmt.Sprintf("---Idle: %v \n", b.Idle)
	s += fmt.Sprintf("---Talk: %v \n", b.Talk)
	return s
}
func (ch *CallHistory) ToString() {
	for i, v := range ch.CallList {
		fmt.Println("")
		fmt.Printf("#%d \n", i)
		fmt.Printf("Call number: %s \n", v.Number)
		fmt.Printf("Date: %s \n", v.ToDate())
		fmt.Printf("Time: %s \n", v.ToTime())
		fmt.Printf("Duration: %d second \n", v.Duration)
		fmt.Println("")
	}
}

func (ch *CallHistory) AddCall(call Call) {
	ch.CallList = append(ch.CallList, call)
	fmt.Printf("Call with number %v added \n", call.Number)
}

func (ch *CallHistory) DeleteCall(call Call) {
	index := getIndex(ch.CallList, call)
	if index != -1 {
		ch.CallList = append(ch.CallList[:index], ch.CallList[index+1:]...)
		fmt.Printf("Call with number %v deleted \n", call.Number)
	}
	fmt.Printf("Call with number %v not found \n", call.Number)
}
func (ch *CallHistory) DeleteLongestCall() {
	var longestCall int
	var index int
	for i, v := range ch.CallList {
		if v.Duration > longestCall {
			longestCall = v.Duration
			index = i
		}
	}
	ch.CallList = append(ch.CallList[:index], ch.CallList[index+1:]...)
	fmt.Println("Call with longest duration deleted")

}

func (ch *CallHistory) ClearAll() {
	ch.CallList = nil
	fmt.Println("Call history deleted")
}

func (ch *CallHistory) CalculatePrice(price float64) {
	var total float64
	for _, v := range ch.CallList {
		total += float64(v.Duration/60) * price
	}
	fmt.Printf("Total price is %v \n", total)
}

func getIndex(callList []Call, call Call) int {
	for i, v := range callList {
		if strings.EqualFold(v.Number, call.Number) && v.TimeToString() == call.TimeToString() {
			return i
		}
	}
	return -1
}

func (c Call) ToDate() string {
	return c.Time.Format("dd/MM/yyyy")
}
func (c Call) ToTime() string {
	return c.Time.Format("hh:mm:ss")
}
func (c Call) TimeToString() string {
	return c.Time.Format("dd/MM/yyyy hh:mm:ss")
}
