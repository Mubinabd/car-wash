package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/service"
)

func CartHandler(cartservice *service.CartService) func(message []byte) {
	return func(message []byte) {
		var cart pb.CreateCartReq
		if err := json.Unmarshal(message, &cart); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respcart, err := cartservice.CreateCart(context.Background(), &cart)
		if err != nil {
			log.Printf("Cannot create cart via Kafka: %v", err)
			return
		}
		log.Printf("Created cart: %+v",respcart)
	}
}
