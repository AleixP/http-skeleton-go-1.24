package fruit

type Fruit struct {
	ID    *int   `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

func CreateFromPrimitive(name string, color string) *Fruit {
	return &Fruit{
		Name:  name,
		Color: color,
	}
}
