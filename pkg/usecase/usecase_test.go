package usecase

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CyberAgentHack/server-performance-tuning-2023/pkg/repository"
)

func TestNew(t *testing.T) {
	db := &repository.Database{}
	want := &UsecaseImpl{
		db: db,
	}
	got := NewUsecase(db)
	got.validate = nil
	require.Equal(t, want, got)
}

func TestInterfaceCheck(t *testing.T) {
	require.Implements(t, new(Usecase), new(UsecaseImpl))
}
