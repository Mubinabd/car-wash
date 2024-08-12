package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/service"
)

func ServiceHandler(serviceservice *service.ServiceService) func(message []byte) {
	return func(message []byte) {
		var service pb.AddServiceReq
		if err := json.Unmarshal(message, &service); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respservice, err := serviceservice.AddService(context.Background(), &service)
		if err != nil {
			log.Printf("Cannot create service via Kafka: %v", err)
			return
		}
		log.Printf("Created service: %+v",respservice)
	}
}

func UpdateserviceHandler(serviceservice *service.ServiceService) func(message []byte) {
	return func(message []byte) {
		var service pb.UpdateServiceReq
		if err := json.Unmarshal(message, &service); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respservice, err := serviceservice.UpdateService(context.Background(), &service)
		if err != nil {
			log.Printf("Cannot create service via Kafka: %v", err)
			return
		}
		log.Printf("Created service: %+v",respservice)
	}
}

func DeleteserviceHandler(serviceservice *service.ServiceService) func(message []byte) {
	return func(message []byte) {
		var service pb.DeleteServiesReq
		if err := json.Unmarshal(message, &service); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respservice, err := serviceservice.DeleteService(context.Background(), &service)
		if err != nil {
			log.Printf("Cannot create service via Kafka: %v", err)
			return
		}
		log.Printf("Created service: %+v",respservice)
	}
}
