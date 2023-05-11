package ctx

import (
	"fmt"
	"time"
)

type Gateway interface {
	Send(userID, message string)
}

type Handler struct {
	Gateway           Gateway
	RateLimits        map[string]time.Duration
	RecipientLastSent map[string]map[string]time.Time
}

func (s *Handler) Send(notificationType, userID, message string) error {
	maxInterval, ok := s.RateLimits[notificationType]
	if !ok {
		return fmt.Errorf("Invalid notification type")
	}

	currentTime := time.Now()
	lastSentTime, exists := s.RecipientLastSent[userID][notificationType]
	if exists && currentTime.Sub(lastSentTime) < maxInterval {
		return fmt.Errorf("Rate limit exceeded for recipient: %s", userID)
	}

	if s.RecipientLastSent[userID] == nil {
		s.RecipientLastSent[userID] = make(map[string]time.Time)
	}
	s.RecipientLastSent[userID][notificationType] = currentTime

	s.Gateway.Send(userID, message)

	return nil
}

func NewHandler(gateway Gateway) *Handler {
	rateLimits := map[string]time.Duration{
		"status":    time.Second / 2, // 2 por minuto
		"news":      time.Hour * 24,  // 1 por día
		"marketing": time.Hour * 3,   // 3 por hora
	}

	return &Handler{
		Gateway:           gateway,
		RateLimits:        rateLimits,
		RecipientLastSent: make(map[string]map[string]time.Time),
	}
}
