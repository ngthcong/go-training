package main

type (
	Phone struct{
		Model string
		Manufacturer string
		Price int64
		Owner string
		Battery Battery
		Display Display
	}
	Battery struct {
		Model string
		Idle  float32
		Talk  float32
	}
	Display struct {
		Size float32
		Colors string
	}

)

func NewPhone(model string,manufacturer string) *Phone{
	return &Phone{
		Model:        model,
		Manufacturer: manufacturer,
	}
}
func NewPhone


func main() {

}
