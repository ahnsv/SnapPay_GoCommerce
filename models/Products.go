package models

// Products strcut
type Products struct {
	ID        int
	Title     string
	Subtitle  string
	Inventory string
	Options   []string
	Price     string
	Image     []string
}
