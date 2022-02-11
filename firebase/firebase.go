package firebase

import (
	"fmt"
	"github.com/appleboy/go-fcm"
	"notify/models"
	"os"
	"log"
)

func SendPush(message models.MessageIn){
	// Create message first
	//log.Printf("Send push...", message)

	for _, token := range message.FcmTokens {
		msg := &fcm.Message{
			To: token,
			Notification: &fcm.Notification{
				Title:        message.Title,
				Body:         message.Body,
				Sound:        "default",

			},
			Data: map[string]interface{}{
				"extra_uid":message.ExtraUID,
				"action": message.Action,
			},
		}

		// Create fcm client
		client, err := fcm.NewClient(os.Getenv("FCM_API_KEY"))
		if err != nil {
			fmt.Print("FCM credentials is not correct")
			return
		}

		response, err := client.Send(msg)
		if err != nil {
			fmt.Print("Can't send push message")
			return
		}

		fmt.Printf("Message from fcm %s", response)
	}

}
