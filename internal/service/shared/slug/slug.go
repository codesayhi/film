package slug

import (
	"regexp"
	"strconv"

	"github.com/google/uuid"
)

type Generator interface {
	Make(input string) string
}

const DefaultBase = "item"

// MakeUnique MakeUnique: nhận base + danh sách slugs đã có (query từ repo), rồi trả slug unique.
// Không cần ctx/ignoreID vì không đụng DB.
func MakeUnique(base string, slugs []string) string {
	if base == "" {
		return DefaultBase + "-" + uuid.NewString()
	}

	// match đúng pattern: ^base-(\d+)$
	re := regexp.MustCompile("^" + regexp.QuoteMeta(base) + `-(\d+)$`)

	used := make(map[int]bool, len(slugs))

	for _, s := range slugs {
		if s == base {
			used[0] = true
			continue
		}

		m := re.FindStringSubmatch(s)
		if len(m) == 2 {
			if n, err := strconv.Atoi(m[1]); err == nil && n > 0 {
				used[n] = true
			}
		}
	}

	// base chưa dùng -> dùng base
	if !used[0] {
		return base
	}

	// base đã dùng -> tìm suffix nhỏ nhất chưa dùng
	for i := 1; ; i++ {
		if !used[i] {
			// Trả về chuỗi
			return base + "-" + strconv.Itoa(i)
		}
	}
}
