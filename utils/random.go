package utils

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func UUID() string {
	return uuid.New().String()
}

func UuidAndTimeStamp() string {
	return fmt.Sprintf("%s-%d", UUID(), time.Now().Unix())
}
