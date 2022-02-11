package firebase

import (
	"fmt"
	"github.com/appleboy/go-fcm"
	"notify/models"
)

func SendPush(message models.MessageIn, client fcm.Client){
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

		response, err := client.Send(msg)
		if err != nil {
			fmt.Printf("Can't send push message", err)
			return
		}

		fmt.Printf("Message from fcm %s", response)
	}

}
