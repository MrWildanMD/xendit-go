package promotion_test

import (
	"fmt"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/promotion"
	"log"
	"time"
)

func ExampleCreatePromotion() {
	xendit.Opt.SecretKey = "examplesecretkey"

	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)

	createPromotionData := promotion.CreatePromotionParams{
		ReferenceID: "BRI_20_JAN",
		Description: "20% discount applied for all BRI cards",
		BinList: []string{
			"400000",
			"460000",
		},
		DiscountPercent:   20,
		Currency:          "IDR",
		StartTime:         &startTime,
		EndTime:           &endTime,
		ChannelCode:       "BRI",
		MinOriginalAmount: 25000,
		MaxDiscountAmount: 5000,
	}

	promotionResp, err := promotion.CreatePromotion(&createPromotionData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created promotion: %+v\n", promotionResp)
}
