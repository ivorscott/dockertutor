package tutor

import (
	"testing"
)

func TestNewTutorial(t *testing.T) {
	tut, _ := NewTutorial("docker")
	t.Logf("===========%v", tut)

	//if tut.Category != "docker" {
	//	t.Errorf("Category not set correctly")
	//}
}
