package helpers

import (
	"strings"

	"github.com/juju/errors"
)

func FoldErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	ss := make([]string, 0, len(errs))
	for _, e := range errs {
		if e != nil {
			ss = append(ss, errors.ErrorStack(e))
		}
	}
	if len(ss) == 0 {
		return nil
	}
	return errors.New(strings.Join(ss, "\n"))
}
