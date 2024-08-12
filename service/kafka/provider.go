package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/service"
)

func ProviderHandler(providerservice *service.ProviderService) func(message []byte) {
	return func(message []byte) {
		var provider pb.RegisterProviderReq
		if err := json.Unmarshal(message, &provider); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respprovider, err := providerservice.RegisterProvider(context.Background(), &provider)
		if err != nil {
			log.Printf("Cannot create provider via Kafka: %v", err)
			return
		}
		log.Printf("Created provider: %+v",respprovider)
	}
}

func UpdateProviderHandler(providerservice *service.ProviderService) func(message []byte) {
	return func(message []byte) {
		var provider pb.UpdateProviderReq
		if err := json.Unmarshal(message, &provider); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respprovider, err := providerservice.UpdateProvider(context.Background(), &provider)
		if err != nil {
			log.Printf("Cannot create provider via Kafka: %v", err)
			return
		}
		log.Printf("Created provider: %+v",respprovider)
	}
}

func DeleteProviderHandler(providerservice *service.ProviderService) func(message []byte) {
	return func(message []byte) {
		var provider pb.DeleteProviderReq
		if err := json.Unmarshal(message, &provider); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respprovider, err := providerservice.DeleteProvider(context.Background(), &provider)
		if err != nil {
			log.Printf("Cannot create provider via Kafka: %v", err)
			return
		}
		log.Printf("Created provider: %+v",respprovider)
	}
}
