package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/service"
)

func BookingHandler(bookingservice *service.BookingsService) func(message []byte) {
	return func(message []byte) {
		var booking pb.AddBookingReq
		if err := json.Unmarshal(message, &booking); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respbooking, err := bookingservice.AddBooking(context.Background(), &booking)
		if err != nil {
			log.Printf("Cannot create booking via Kafka: %v", err)
			return
		}
		log.Printf("Created booking: %+v",respbooking)
	}
}

func UpdateHandler(bookservice *service.BookingsService) func(message []byte) {
	return func(message []byte) {
		var booking pb.UpdateBookingReq
		if err := json.Unmarshal(message, &booking); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respbooking, err := bookservice.UpdateBooking(context.Background(), &booking)
		if err != nil {
			log.Printf("Cannot create booking via Kafka: %v", err)
			return
		}
		log.Printf("Created booking: %+v",respbooking)
	}
}
func DeleteBookingHandler(bookingservice *service.BookingsService) func(message []byte) {
	return func(message []byte) {
		var booking pb.DeleteBookingReq
		if err := json.Unmarshal(message, &booking); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respbooking, err := bookingservice.DeleteBooking(context.Background(), &booking)
		if err != nil {
			log.Printf("Cannot create booking via Kafka: %v", err)
			return
		}
		log.Printf("Created booking: %+v",respbooking)
	}
}