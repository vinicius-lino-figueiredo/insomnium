package insomnium

import (
	"context"
	"errors"
	"fmt"

	"github.com/vinicius-lino-figueiredo/gedb"
	"github.com/vinicius-lino-figueiredo/gedb/domain"
)

var (
	// ErrNoParent TODO
	ErrNoParent = errors.New("should have parent ID")
)

// Model represents any type of data that is stored by insomnium.
type Model interface {
	GetID() string
	GetParentID() string
}

// Repository represents a datafile.
type Repository[T Model] struct {
	db        gedb.GEDB
	hasParent bool
	size      int
}

// NewRepository returns a new instance of [*Repository].
func NewRepository[T Model](ctx context.Context, path string, hasParent bool, size int) (*Repository[T], error) {
	db, err := gedb.NewDB(gedb.WithFilename(path))
	if err != nil {
		return nil, err
	}
	err = db.LoadDatabase(ctx)
	if err != nil {
		return nil, err
	}
	return &Repository[T]{
		hasParent: hasParent,
		size:      size,
		db:        db,
	}, nil
}

// Load loads the datafile to refresh the cached data.
func (r *Repository[T]) Load() error {
	return r.db.LoadDatabase(context.Background())
}

// Get returns the given record, finding by ID.
func (r *Repository[T]) Get(ctx context.Context, key string) (res T, err error) {
	err = r.db.FindOne(ctx, map[string]any{"_id": key}, &res)
	return res, err
}

// GetChildren returns all children by parent ID.
func (r *Repository[T]) GetChildren(ctx context.Context, key string) ([]T, error) {
	cur, err := r.db.Find(ctx, map[string]any{"parentId": key})
	if err != nil {
		return nil, err
	}

	return r.readCursor(ctx, cur)
}

// GetChild returns the first child of a parent ID. Fails if it finds zero.
func (r *Repository[T]) GetChild(ctx context.Context, key string) (res T, err error) {
	err = r.db.FindOne(ctx, map[string]any{"parentId": key}, &res)
	return res, err
}

// Create inserts a new record.
func (r *Repository[T]) Create(ctx context.Context, t T) error {
	if r.hasParent && t.GetParentID() == "" {
		return fmt.Errorf("%w: %v", ErrNoParent, t)
	}
	_, err := r.db.Insert(ctx, t)
	return err
}

// Remove removes the given record
func (r *Repository[T]) Remove(ctx context.Context, t T) error {
	_, err := r.db.Remove(ctx, t, domain.WithRemoveMulti(false))
	return err
}

// Find returns a slice version of the cursor for a query.
func (r *Repository[T]) Find(ctx context.Context, query any, opts ...gedb.FindOption) ([]T, error) {
	cur, err := r.db.Find(ctx, query, opts...)
	if err != nil {
		return nil, err
	}

	return r.readCursor(ctx, cur)
}

// All returns all cached records.
func (r *Repository[T]) All(ctx context.Context) ([]T, error) {
	cur, err := r.db.GetAllData(ctx)
	if err != nil {
		return nil, err
	}
	return r.readCursor(ctx, cur)

}

func (r *Repository[T]) readCursor(ctx context.Context, cur gedb.Cursor) ([]T, error) {
	res := make([]T, 0, r.size)
	for cur.Next() {
		var t T
		if err := cur.Scan(ctx, &t); err != nil {
			return nil, err
		}
		res = append(res, t)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
