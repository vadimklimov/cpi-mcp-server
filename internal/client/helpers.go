package client

import (
	"context"
	"fmt"
	"sync"

	"github.com/panjf2000/ants/v2"
	"github.com/vadimklimov/cpi-mcp-server/internal/config"
)

func Run[I any, O any](
	ctx context.Context,
	values []I,
	task func(context.Context, I) ([]O, error),
) ([]O, []error, error) {
	var (
		results []O
		errs    []error
		wg      sync.WaitGroup
	)

	resultc := make(chan O)
	errc := make(chan error)

	pool, err := ants.NewPoolWithFunc(config.MaxConcurrency(), func(value any) {
		defer wg.Done()

		entries, err := task(ctx, value.(I))
		if err != nil {
			select {
			case errc <- err:
			case <-ctx.Done():
				return
			}

			return
		}

		for _, entry := range entries {
			select {
			case resultc <- entry:
			case <-ctx.Done():
				return
			}
		}
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create worker pool: %w", err)
	}

	defer pool.Release()

	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			select {
			case result, ok := <-resultc:
				if !ok {
					resultc = nil
				} else {
					results = append(results, result)
				}
			case err, ok := <-errc:
				if !ok {
					errc = nil
				} else {
					errs = append(errs, err)
				}
			case <-ctx.Done():
				return
			}

			if resultc == nil && errc == nil {
				return
			}
		}
	}()

	for _, value := range values {
		wg.Add(1)

		_ = pool.Invoke(value)
	}

	wg.Wait()
	close(resultc)
	close(errc)
	<-done

	if ctx.Err() != nil {
		return results, errs, fmt.Errorf("failed to execute parallel tasks: %w", ctx.Err())
	}

	return results, errs, nil
}
