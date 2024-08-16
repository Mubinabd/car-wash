package pkg

import (
	"log"

	"github.com/Mubinabd/car-wash/kafka"
	"github.com/Mubinabd/car-wash/service"
)

func Reader(brokers []string, kcm *kafka.KafkaConsumerManager, bservice *service.BookingsService, cservice *service.CartService,prservice *service.ProviderService,rservice *service.ReviewService,nservice *service.NotificationService,sservice *service.ServiceService) {
	if err := kcm.RegisterConsumer(brokers, "cr-booking", "product", kafka.BookingHandler(bservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "up-booking", "product", kafka.UpdateHandler(bservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-booking", "product", kafka.DeleteBookingHandler(bservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-cart", "product", kafka.CartHandler(cservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-provider", "product", kafka.ProviderHandler(prservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-provider' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-provider': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "up-provider", "product", kafka.UpdateProviderHandler(prservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'up-provider' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-provider': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-provider", "product", kafka.DeleteProviderHandler(prservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'dl-provider' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-provider': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-review", "product", kafka.ReviewHandler(rservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-review' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-review': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "notif", "product", kafka.NotificationHandler(nservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'notif' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-review': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "up-review", "product", kafka.UpdatereviewHandler(rservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'up-review' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-review': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-review", "product", kafka.DeletereviewHandler(rservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'dk-review' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-review': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-service", "service", kafka.ServiceHandler(sservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-service' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-service': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-service", "service", kafka.DeleteserviceHandler(sservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'dl-service' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'dl-service': %v", err)

		}
	}
}
