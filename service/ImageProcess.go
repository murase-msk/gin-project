/*
サービス
*/
package service

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

//画像処理関係
type ImageProcess struct {
	x int
	y int
}

//画像ファイルからテキストを取得する
func (p *ImageProcess) GetTextFromImage(filePath string) []string {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Println(err)
	}

	f, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	image, err := vision.NewImageFromReader(f)
	if err != nil {
		log.Println(err)
	}
	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Println(err)
	}

	if len(annotations) == 0 {
		fmt.Println("No text found.")
	} else {
		fmt.Println("Text:")
		for _, annotation := range annotations {
			fmt.Println(annotation.Description)
		}
	}

	return nil
}

//画像のファイルパスから画像解析しラベルを取得する
func (p *ImageProcess) GetLabelFromImage(filePath string) []string {
	//
	// Google Vision API.
	//
	ctx := context.Background()
	// Creates a client.
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Sets the name of the image file to annotate.
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	labelList := []string{}
	//fmt.Println("Labels:")
	for _, label := range labels {
		//fmt.Println(label.Description)
		labelList = append(labelList, label.Description)
	}
	return labelList
}
