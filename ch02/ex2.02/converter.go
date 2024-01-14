package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bapturp/converter/lenconv"
	"github.com/bapturp/converter/weiconv"
	"github.com/bapturp/tempconv"
)

func main() {
	if len(os.Args) <= 1 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			s := input.Text()
			for _, val := range strings.Split(s, " ") {
				Conversion(val)
			}
		}
	} else {
		for _, val := range os.Args[1:] {
			Conversion(val)
		}
	}
}

func Conversion(s string) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	tempC := tempconv.Celsius(f)
	tempF := tempconv.Fahrenheit(f)
	tempK := tempconv.Kelvin(f)

	lengthM := lenconv.Meter(f)
	lengthF := lenconv.Foot(f)

	weightK := weiconv.Kilogram(f)
	weightP := weiconv.Pound(f)

	fmt.Printf("%s\t%s\n", tempC.String(), tempconv.CToF(tempC).String())    // Celsius to Fahrenheit
	fmt.Printf("%s\t%s\n", tempF.String(), tempconv.FToC(tempF).String())    // Fahrenheit to Celsius
	fmt.Printf("%s\t%s\n", tempC.String(), tempconv.CToK(tempC).String())    // Celsius to Kelvin
	fmt.Printf("%s\t%s\n", tempK.String(), tempconv.KToC(tempK).String())    // Kevin to Celsius
	fmt.Printf("%s\t%s\n", lengthM.String(), lenconv.MToF(lengthM).String()) // Meter to Foot
	fmt.Printf("%s\t%s\n", lengthF.String(), lenconv.FToM(lengthF).String()) // Foot to Meter
	fmt.Printf("%s\t%s\n", weightK.String(), weiconv.KToP(weightK).String()) // Kilogram to Pound
	fmt.Printf("%s\t%s\n", weightP.String(), weiconv.PToK(weightP).String()) // Kilogram to Pound
}
