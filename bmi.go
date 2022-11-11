package main
/*
*****************************************************************

	@author: Adenugba Adeoluwa
	Date: 9th  November 2022
	Purpose: To test the multiple functions in basic mass index application

********************************************************************
*/
import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("WELCOME TO GOLANG BMI CONVERTER")
	

	for {

		fmt.Println("choose 1 option")
		fmt.Print("[1] Calculate BMI in Metric", "\n", "[2] Calculate BMI in Standard", "\n", "[3] Exit Program", "\n")
		var value int
		fmt.Scan(&value)
		if value == 1 {
			var cm, kg float64
			fmt.Print("enter  height(cm): ")
			fmt.Scan(&cm)
			fmt.Print("enter  weight(kg): ")
			fmt.Scan(&kg)

			bmi_metric := metric_converter(cm, kg)
			category := message(bmi_metric)
			fmt.Printf("Your BMI is %v, %v", bmi_metric, category)

		} else if value == 2 {
			var foot, inches, pounds float64
			fmt.Print("enter  height(foot): ")
			fmt.Scan(&foot)
			fmt.Print("enter  height(inches): ")
			fmt.Scan(&inches)
			fmt.Print("enter  weight(lbs): ")
			fmt.Scan(&pounds)

			bmi_standard := standard_converter(foot, inches, pounds)
			category := message(bmi_standard)
			fmt.Printf("Your BMI is %.1f, %v\n", bmi_standard, category)
		} else if value == 3 {
			os.Exit(1)

		} else {
			fmt.Println("invalid character")
		}

	}
}
func metric_converter(height float64, weight float64) float64 {
	num := float64((weight / height / height) * 10000)
	return math.Round(num*10) / 10

}
func standard_converter(foot float64, inches float64, weight float64) float64 {
	//converting foot to inches
	height := (foot * 12) + inches

	return math.Round(float64((weight/(height*height))*703)*10) / 10

}

func message(input float64) string {

	switch {
	case input < 18.5:
		return "Underweight"
	case 18.5 < input && input < 24.9:
		return "Normal weight"
	case 25 < input && input < 29.9:
		return "Overweight"
	case input > 30:
		return "Obesity"
	default:
		return "failed"
	}
}
