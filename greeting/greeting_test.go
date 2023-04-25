package greeting_test

import (
	"fmt"
	"testing"

	"github.com/lkeix/inspect-coverage-saas/greeting"
)

func TestNewHoge(t *testing.T) {
	testcases := []struct {
		name string
	}{{name: ""}}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			fmt.Println("write your unit test!")
		})
	}
}
func TestHogeGreeting(t *testing.T) {
	testcases := []struct {
		name string
	}{{name: ""}}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			fmt.Println("write your unit test!")
		})
	}
}
func TestNewFuga(t *testing.T) {
	testcases := []struct {
		name string
	}{{name: ""}}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			fmt.Println("write your unit test!")
		})
	}
}
func TestFugaGreeting(t *testing.T) {
	testcases := []struct {
		name string
	}{{name: "test"}}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			f := greeting.NewFuga("hoge")
			f.Greeting()
		})
	}
}
