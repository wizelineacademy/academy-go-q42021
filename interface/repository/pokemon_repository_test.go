package repository_test

import (
	"errors"
	"testing"

	"github.com/hamg26/academy-go-q42021/domain/model"
	ir "github.com/hamg26/academy-go-q42021/interface/repository"
	"github.com/hamg26/academy-go-q42021/testutil"
	ur "github.com/hamg26/academy-go-q42021/usecase/repository"
	"github.com/stretchr/testify/assert"
)

func setupMockRepository(t *testing.T, fakeError error, testCase string) (r ur.PokemonRepository) {
	mycsv := testutil.NewCsvMock(t, fakeError, testCase)
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
				r := setupMockRepository(t, nil, "SUCCESS")
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.Equal(t, 2, len(p))
			},
		},
		"invalid id from csv": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				r := setupMockRepository(t, nil, "INVALID_ID")
				return r
			},
			assert: func(t *testing.T, p []*model.Pokemon, err error) {
				assert.NotEmpty(t, err)
			},
		},
		"error": {
			arrange: func(t *testing.T) ur.PokemonRepository {
				r := setupMockRepository(t, errors.New("fake error"), "NIL")
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
				r := setupMockRepository(t, nil, "SUCCESS")
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
				r := setupMockRepository(t, nil, "SUCCESS")
				id := uint64(0)
				return r, id
			},
			assert: func(t *testing.T, p *model.Pokemon, err error) {
				assert.Nil(t, err)
				assert.Nil(t, p)
			},
		},
		"invalid id from csv": {
			arrange: func(t *testing.T) (ur.PokemonRepository, uint64) {
				r := setupMockRepository(t, nil, "INVALID_ID")
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
				r := setupMockRepository(t, errors.New("fake error"), "NIL")
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
