package repository_test

import (
	"testing"

	"github.com/hamg26/academy-go-q42021/domain/model"
	ir "github.com/hamg26/academy-go-q42021/interface/repository"
	"github.com/hamg26/academy-go-q42021/testutil"
	ur "github.com/hamg26/academy-go-q42021/usecase/repository"
	"github.com/stretchr/testify/assert"
)

func setupSuccess(t *testing.T) (r ur.PokemonRepository) {
	records := [][]string{
		{"1", "name1", "type1"},
		{"2", "name2", "type2"},
	}
	return setupRepository(t, records)
}

func setupInvalidId(t *testing.T) (r ur.PokemonRepository) {
	records := [][]string{
		{"INVALID_ID", "name1", "type1"},
		{"2", "name2", "type2"},
	}
	return setupRepository(t, records)
}

func setupError(t *testing.T) (r ur.PokemonRepository) {
	return setupRepository(t, nil)
}

func setupRepository(t *testing.T, records [][]string) (r ur.PokemonRepository) {
	mycsv := testutil.NewCsvMock(t, records)
	return ir.NewPokemonRepository(mycsv)
}

func TestPokemonRepository_FindAll(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) ur.PokemonRepository
		assert  func(t *testing.T, p []*model.Pokemon, err error)
	}{
		"success": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				r := setupSuccess(t)
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 2, len(p))
			},
		},
		"invalid data from csv": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				r := setupInvalidId(t)
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.NotEmpty(t, err)
			},
		},
		"error": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				r := setupError(t)
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.NotEmpty(t, err)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			r := tt.arrange(t)
			err, p := r.FindAll()
			tt.assert(t, p, err)
		})
	}
}

func TestPokemonRepository_FindOne(t *testing.T) {
	t.Helper()

	cases := map[string]struct {
		arrange func(t *testing.T) (ur.PokemonRepository, uint64)
		assert  func(t *testing.T, p *model.Pokemon, err error)
	}{
		"success": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				r := setupSuccess(t)
				id := uint64(1)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, p)
			},
		},
		"id not found": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				r := setupSuccess(t)
				id := uint64(0)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.Nil(t, p)
			},
		},
		"invalid data from csv": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				r := setupInvalidId(t)
				id := uint64(2)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.NotEmpty(t, err)
				assert.Nil(t, p)
			},
		},
		"error": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				r := setupError(t)
				id := uint64(1)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.NotEmpty(t, err)
				assert.Nil(t, p)
			},
		},
	}

	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			r, id := tt.arrange(t)
			err, p := r.FindOne(id)
			tt.assert(t, p, err)
		})
	}
}
