package solid

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestMigrationFly_Fly(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Migration Fly Test",
			want: fmt.Sprintf("Migration Fly @%s", time.Now().Format(time.ANSIC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MigrationFly{}
			if got := m.Fly(); got != tt.want {
				t.Errorf("Fly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInjuredFly_Fly(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Injured Fly Test",
			want: fmt.Sprintf("Injured Fly @%s", time.Now().Format(time.ANSIC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InjuredFly{}
			if got := i.Fly(); got != tt.want {
				t.Errorf("Fly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDuck(t *testing.T) {
	type args struct {
		flyBehaviour FlyBehaviour
	}
	tests := []struct {
		name string
		args args
		want *Duck
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				flyBehaviour: &MigrationFly{},
			},
			want: &Duck{
				FlyBehaviour: &MigrationFly{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDuck(tt.args.flyBehaviour); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDuck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuckTest(t *testing.T) {
	type args struct {
		duck Duck
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				duck: Duck{
					FlyBehaviour: &MigrationFly{},
				},
			},
			want: fmt.Sprintf("Migration Fly @%s", time.Now().Format(time.ANSIC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DuckTest(tt.args.duck); got != tt.want {
				t.Errorf("DuckTest() = %v, want %v", got, tt.want)
			}
		})
	}
}
