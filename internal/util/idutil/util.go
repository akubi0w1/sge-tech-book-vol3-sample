package idutil

import (
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"github.com/google/uuid"
)

func UUID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", terror.Wrapf(terror.CodeInternal, err, "failed to generate new uuid.")
	}

	return uuid.String(), nil
}
