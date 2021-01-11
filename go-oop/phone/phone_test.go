package phone


import (
"fmt"
"testing"
"time"
)

func TestGSM(t *testing.T) {
	GSMs := []GSM{
		{
			Model:        "GM-9000",
			Manufacturer: "SAMSUNG",
			Price:        20000000.00,
			Owner:        "John",
			battery: Battery{
				Model: "TTA-2",
				Idle:  30,
				Talk:  10,
			},
			display: Display{
				Size:   6.99,
				Colors: 32000000.00,
			},
		},
		{
			Model:        "IPHONE 6s",
			Manufacturer: "APPLE",
			Price:        20000000.00,
			Owner:        "Jinny",
			battery: Battery{
				Model: "TTA-3",
				Idle:  30,
				Talk:  10,
			},
			display: Display{
				Size:   6.99,
				Colors: 32000000.00,
			},
		},
	}

	for _, v := range GSMs {
		fmt.Println(v.ToString())
	}
}

func TestCallHistory(t *testing.T) {
	x := time.Date(2021, 1, 11, 12, 22, 33, 0, time.UTC)
	callHistory := CallHistory{
		CallList: []Call{{
			Number:   "111111",
			Time:     x.Add(time.Minute),
			Duration: 22,
		},
			{
				Number:   "111112",
				Time:     x.Add(time.Hour),
				Duration: 111,
			},
			{
				Number:   "111113",
				Time:     x.Add(time.Minute),
				Duration: 111,
			},
			{
				Number:   "111114",
				Time:     x,
				Duration: 111,
			},
		},
	}

	callHistory.ToString()

	callHistory.CalculatePrice(0.37)
	callHistory.DeleteLongestCall()
	callHistory.CalculatePrice(0.37)
	callHistory.ClearAll()
	callHistory.ToString()
}
