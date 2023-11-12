package middleware

type CulversData struct {
	Town string
	Flavor string
}

type Data interface {
	IsBlank() bool
}

func (c *CulversData) IsBlank() bool {
	return c.Flavor == ""
}