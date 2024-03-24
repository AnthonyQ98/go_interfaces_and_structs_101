package sounds

import "fmt"

// Define unexported structs and fields
// will create getters/setters for the fields,
// and a constructor for building new types of these structs.
type dog struct {
	breed string
	person
}

type phone struct {
	model string
	person
}

type person struct {
	name string
}

// Interface for a function MakeSound() which takes anything that defines a makeNoise() method.
type soundMaker interface {
	makeNoise()
}

// Getters
func (d *dog) Breed() string {
	return d.breed
}

func (p *phone) Model() string {
	return p.model
}

func (p *person) Owner() string {
	return p.name
}

// Setters
func (d *dog) SetBreed(breed string) {
	d.breed = breed
}

func (p *phone) SetModel(model string) {
	p.model = model
}

func (p *person) SetOwner(owner string) {
	p.name = owner
}

// Define stringers
func (p *phone) String() string {
	return fmt.Sprintf("Beep, beep")
}

func (d *dog) String() string {
	return fmt.Sprintf("Woof, woof")
}

// Two common methods between phone and dog
func (p *phone) makeNoise() {
	fmt.Println(p.String())
}

func (d *dog) makeNoise() {
	fmt.Println(d.String())
}

// CallSound Function that accepts interface implemented by dog and phone
func CallSound(sm soundMaker) {
	sm.makeNoise()
}

// NewDog constructs a dog
func NewDog(breed string) *dog {
	d := new(dog)
	if len(breed) > 0 {
		d.breed = breed
	} else {
		d.breed = "Greyhound"
	}
	d.name = "Julio"
	return d
}

// NewPhone constructs a phone
func NewPhone(model string) *phone {
	p := new(phone)
	if len(model) > 0 {
		p.model = model
	} else {
		p.model = "iPhone 13"
	}

	p.name = "John"
	return p
}
