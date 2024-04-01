package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	payload := map[string]interface{}{
		"to": "U062a486fb8ca68b0b47e60691f94a477",
		"messages": []map[string]interface{}{
			{
				"type": "text",
				"text": "Hello! botnoi",
			},
			{
				"type":    "template",
				"altText": "Button Template",
				"template": map[string]interface{}{
					"type": "buttons",
					"text": "Click the button:",
					"actions": []map[string]interface{}{
						{
							"type":  "uri",
							"label": "Visit BOTNOI Website",
							"uri":   "https://botnoi.ai/",
						},
					},
				},
			},
			{
				"type":    "template",
				"altText": "Carousel Template",
				"template": map[string]interface{}{
					"type": "carousel",
					"columns": []map[string]interface{}{
						{
							"thumbnailImageUrl": "https://botnoi.ai/assets/etc/botnoi.png",
							"title":             "BOTNOI Faq",
							"text":              "BOTNOI Faq",
							"actions": []map[string]interface{}{
								{
									"type":  "uri",
									"label": "View Details",
									"uri":   "https://botnoi.ai/faq",
								},
							},
						},
						{
							"thumbnailImageUrl": "https://assets.brandinside.asia/uploads/2017/03/botnoi-1111.jpg",
							"title":             "CONTACT",
							"text":              "CONTACT BOTNOI",
							"actions": []map[string]interface{}{
								{
									"type":  "uri",
									"label": "View Details",
									"uri":   "https://botnoi.ai/contact",
								},
							},
						},
					},
				},
			},
			{
				"type": "text",
				"text": "เลือกตัวเลือกด้านล่าง:",
				"quickReply": map[string]interface{}{
					"items": []map[string]interface{}{
						{
							"type": "action",
							"action": map[string]interface{}{
								"type":  "message",
								"label": "ตัวเลือก 1",
								"text":  "เลือกตัวเลือก 1",
							},
						},
						{
							"type": "action",
							"action": map[string]interface{}{
								"type":  "message",
								"label": "ตัวเลือก 2",
								"text":  "เลือกตัวเลือก 2",
							},
						},
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error converting payload to JSON:", err)
		return
	}

	url := "https://api.line.me/v2/bot/message/push"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer j7AfK4GXZ52UMI/Jod2Un1ezJfOwPAMCZ+4pCfDjkqZ9GoVT6xLlO54G4Ddzhv8Hii/usUMfavuA0CbRIHg3fjfSjKyHX9mL2Oq/fQKAwTgcgmPRpUPDm2zwlMeLpswCKvKDieIfzsKMbXFAjqDvmwdB04t89/1O/w1cDnyilFU=")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
