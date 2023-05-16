package ctx

import (
	"fmt"
	"time"
)

type Gateway interface {
	Send(userID, message string)
}

type Handler struct {
	Gateway    Gateway
	RateLimits map[string]struct {
		Interval     time.Duration
		MaxMensagges int
	}
	RecipientLastSent map[string]map[string][]time.Time
}

func (s *Handler) Send(notificationType, userID, message string) error {
	rateLimit, ok := s.RateLimits[notificationType]
	if !ok {
		return fmt.Errorf("Invalid notification type")
	}

	currentTime := time.Now()
	lastSentTimes, exists := s.RecipientLastSent[userID][notificationType]
	if exists && len(lastSentTimes) >= rateLimit.MaxMensagges {
		timeDiff := currentTime.Sub(lastSentTimes[len(lastSentTimes)-rateLimit.MaxMensagges])
		if timeDiff < rateLimit.Interval {
			return fmt.Errorf("Rate limit exceeded for recipient: %s", userID)
		}
	}

	if s.RecipientLastSent[userID] == nil {
		s.RecipientLastSent[userID] = make(map[string][]time.Time)
	}
	s.RecipientLastSent[userID][notificationType] = append(s.RecipientLastSent[userID][notificationType], currentTime)

	s.Gateway.Send(userID, message)

	return nil
}

func NewHandler(gateway Gateway) *Handler {

	rateLimits := map[string]struct {
		Interval     time.Duration
		MaxMensagges int
	}{
		"status": {
			Interval:     time.Second * 10,
			MaxMensagges: 2,
		},
		"news": {
			Interval:     time.Hour * 24,
			MaxMensagges: 20,
		},
		"marketing": {
			Interval:     time.Hour,
			MaxMensagges: 10,
		},
	}

	return &Handler{
		Gateway:           gateway,
		RateLimits:        rateLimits,
		RecipientLastSent: make(map[string]map[string][]time.Time),
	}
}
