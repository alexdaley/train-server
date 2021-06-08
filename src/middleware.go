package main

import (
	"context"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func meterRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UnixNano()
		next(w, r)
		endTime := time.Now().UnixNano()

		duration := (endTime - startTime) / (time.Millisecond.Nanoseconds() / time.Nanosecond.Nanoseconds())
		if LogRequestTime {
			log.Printf("Request [%s] duration: %dms\n", getReqId(r), int(duration))
		}
	}
}

func stamp(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "req-id", uuid.NewString())
		next(w, r.WithContext(ctx))
	}
}

func wrap(final http.HandlerFunc) http.HandlerFunc {
	return stamp(CORS(meterRequest(final)))
}
