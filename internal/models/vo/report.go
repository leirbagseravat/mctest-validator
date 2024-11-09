package vo

import (
	"os"
	"time"
)


type Report struct {
	ID string 
	CreatedAt time.Time
	File *os.File
}