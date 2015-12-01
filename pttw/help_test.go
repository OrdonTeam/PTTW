package pttw
import (
	"testing"
	"github.com/assertgo/assert"
)

func TestHelpShouldReturnFalseIsNotAskedForHelp(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{})
	assert.ThatBool(result).IsFalse()
}

func TestHelpShouldReturnTrue1(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{"-h"})
	assert.ThatBool(result).IsTrue()
}

func TestHelpShouldReturnTrue2(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{"-help"})
	assert.ThatBool(result).IsTrue()
}

func TestHelpShouldReturnTrue3(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{"--h"})
	assert.ThatBool(result).IsTrue()
}

func TestHelpShouldReturnTrue4(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{"--help"})
	assert.ThatBool(result).IsTrue()
}

func TestHelpShouldReturnTrue5(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{"-?"})
	assert.ThatBool(result).IsTrue()
}

func TestHelpShouldReturnTrue6(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{"/?"})
	assert.ThatBool(result).IsTrue()
}

func TestHelpShouldSearchWholeParamsList(t *testing.T) {
	assert := assert.New(t)
	result := ShouldDisplayHelp([]string{"", "", "", "-help"})
	assert.ThatBool(result).IsTrue()
}

func ExampleHelpShouldDisplayHelpScreen() {
	DisplayHelpScreen([]string{"./name"})
	// Output:
	//-in	input_file
	//	file with floating numbers separated with whitespaces
	//	if not specify program will use standard input
	//-out	output_file
	//	output file
	//	if not specified program will use standard output
	//-snr	20
	//	signal to noise ratio in [db] default 20
	//Ex	echo "1 1 1 1" | ./name
	//	./name -in input.txt -out output.txt -snr 40
}

func ExampleHelpShouldDisplayHelpScreenWithDifferentFileName() {
	DisplayHelpScreen([]string{"./main"})
	// Output:
	//-in	input_file
	//	file with floating numbers separated with whitespaces
	//	if not specify program will use standard input
	//-out	output_file
	//	output file
	//	if not specified program will use standard output
	//-snr	20
	//	signal to noise ratio in [db] default 20
	//Ex	echo "1 1 1 1" | ./main
	//	./name -in input.txt -out output.txt -snr 40
}