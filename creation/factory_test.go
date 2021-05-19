package creation

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSimpleFactory(t *testing.T) {
	var (
		c IPeople = NewHelloWorld(China)
		a IPeople = NewHelloWorld(America)
		d IPeople = NewHelloWorld(^(China | America))
	)
	assert.Equal(t, c.SayHello(), (&Chinese{}).SayHello())
	assert.Equal(t, a.SayHello(), (&American{}).SayHello())
	assert.Equal(t, d.SayHello(), (&DefaultPeople{}).SayHello())
}
