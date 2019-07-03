package middleware

import (
	"fmt"
	"os"

	"github.com/hecjhs/api-go/api/models"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestMain(m *testing.M) {
	models.DB_init()
	os.Exit(m.Run())
}

func TestReadFromDB(t *testing.T) {
	repo := &Queue{}
	domains := []string{"alpha", "omega", "beta"}

	t.Run("Returning domains", func(t *testing.T) {
		for _, d := range domains {
			domain := repo.ReadFromDB(d)
			fmt.Printf("dominio %+v expected: %v", domain, d)
			assert.Equal(t, domain.Domain, d)
		}
	})

}
