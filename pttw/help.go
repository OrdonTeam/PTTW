package pttw
import "fmt"

var helpParams = map[string]bool{"-h": true, "-help": true, "--h": true, "--help": true, "-?": true, "/?": true}

func ShouldDisplayHelp(params []string) bool {
	for i := 0; i < len(params); i++ {
		if (helpParams[params[i]]) {
			return true
		}
	}
	return false
}

func DisplayHelpScreen(params []string) {
	fmt.Println("-in\tinput_file")
	fmt.Println("\tfile with floating numbers separated with whitespaces")
	fmt.Println("\tif not specify program will use standard input")
	fmt.Println("-out\toutput_file")
	fmt.Println("\toutput file")
	fmt.Println("\tif not specified program will use standard output")
	fmt.Println("-snr\t20")
	fmt.Println("\tsignal to noise ratio in [db] default 20")
	fmt.Print("Ex\techo \"1 1 1 1\" | ")
	fmt.Println(params[0])
	fmt.Println("\t./name -in input.txt -out output.txt -snr 40")
}