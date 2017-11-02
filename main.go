package main

import (
	"fmt"
	"log"
	"os"

	"github.com/YoumnaElghandoor/chatbot"
//	"github.com/YoumnaElghandoor/chatbot-sample/controller"
)

// Autoload environment variables in .env
import _ "github.com/joho/godotenv/autoload"

// func chatbotProcess(session chatbot.Session, message string) (string, error) {
// 	if strings.EqualFold(message, "chatbot") {
// 		return "", fmt.Errorf("This can't be, I'm the one and only %s!", message)
// 	}

// 	var questionMarksCount int
// 	// Try fetching the count of question marks
// 	count, found := session["questionMarksCount"]
// 	// If a count is saved in the session
// 	if found {
// 		// Cast it into an int (since sessions values are generic)
// 		questionMarksCount = count.(int)
// 	} else {
// 		// Otherwise, initialize the count to 1
// 		questionMarksCount = 1
// 	}

// 	// Build the question marks string according to the question marks count
// 	var questionMarks string
// 	for i := 1; i <= questionMarksCount; i++ {
// 		questionMarks += "?"
// 	}

// 	// Save the updated question marks count to the session
// 	session["questionMarksCount"] = questionMarksCount + 1

// 	// Return the response with an extra question mark
// 	return fmt.Sprintf("Hello <b>%s</b>, my name is chatbot. What was yours again%s", message, questionMarks), nil
// }

func main() {
	// Uncomment the following lines to customize the chatbot
	// chatbot.WelcomeMessage = "What's your name?"
	// chatbot.ProcessFunc(chatbotProcess)

	// Use the PORT environment variable
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}
