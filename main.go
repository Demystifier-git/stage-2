package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Structure to hold the response data
type MathProperties struct {
	Number          int     `json:"number"`
	IsPrime         bool    `json:"is_prime"`
	IsPerfectSquare bool    `json:"is_perfect_square"`
	SquareRoot      float64 `json:"square_root"`
	Factorial       int     `json:"factorial"`
	Divisors        []int   `json:"divisors"`
	FunFact         string  `json:"fun_fact"`
}

// Function to check if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to check if a number is a perfect square
func isPerfectSquare(n int) bool {
	sqrt := math.Sqrt(float64(n))
	return sqrt == math.Floor(sqrt)
}

// Function to calculate the factorial of a number
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

// Function to find the divisors of a number
func getDivisors(n int) []int {
	divisors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
		}
	}
	return divisors
}

// Function to generate a fun fact about the number
func getFunFact(n int) string {
	if n == 1 {
		return "1 is the only number that is neither prime nor composite!"
	} else if isPrime(n) {
		return fmt.Sprintf("%d is a prime number!", n)
	} else if isPerfectSquare(n) {
		return fmt.Sprintf("%d is a perfect square!", n)
	} else {
		return fmt.Sprintf("%d is a composite number.", n)
	}
}

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// API route to get mathematical properties
	router.GET("/mathproperties", func(c *gin.Context) {
		// Get the 'number' parameter from the query string
		numberStr := c.DefaultQuery("number", "")
		if numberStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": true})
			return
		}

		// Convert the 'number' to an integer
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": true})
			return
		}

		// Generate mathematical properties
		properties := MathProperties{
			Number:          number,
			IsPrime:         isPrime(number),
			IsPerfectSquare: isPerfectSquare(number),
			SquareRoot:      math.Sqrt(float64(number)),
			Factorial:       factorial(number),
			Divisors:        getDivisors(number),
			FunFact:         getFunFact(number),
		}

		// Set content-type as application/json (should be default, but for clarity)
		c.Header("Content-Type", "application/json")

		// Return JSON response
		c.JSON(http.StatusOK, properties)
	})

	// Run the server and handle error if port is already in use
	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
