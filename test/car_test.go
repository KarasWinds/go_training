package car

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c, err := New("", 100)
	if err != nil {
		t.Fatal("got errors:", err)
	}

	if c == nil {
		t.Error("car should be nil")
	}
}

func TestNewWithAssert(t *testing.T) {
	c, err := New("", 100)
	assert.NotNil(t, err)
	assert.Error(t, err)
	assert.Nil(t, c)

	c, err = New("foo", 100)
	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "foo", c.Name)
}

func TestCar_SetName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		c    *Car
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "no input name",
			c: &Car{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "",
			},
			want: "foo",
		},
		{
			name: "input name",
			c: &Car{
				Name:  "foo",
				Price: 100,
			},
			args: args{
				name: "bar",
			},
			want: "bar",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			log.Println(tt.args)
			if got := tt.c.SetName(tt.args.name); got != tt.want {
				t.Errorf("Car.SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}
