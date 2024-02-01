package generator

import (
	"fmt"
	"strings"
	"unicode"
)

func normalizeName(name string) string {
	var s strings.Builder
	s.Grow(len(name))

	upper := true
	for _, c := range name {
		if upper {
			s.WriteRune(unicode.ToUpper(c))
			upper = false
			continue
		}
		if c == '-' {
			upper = true
			continue
		}
		s.WriteRune(c)
	}
	return s.String()
}

func normalizeTypeRefName(name string) string {
	return fmt.Sprintf("%s", normalizeName(name))
}

func variantInterfaceName(variantName string) string {
	return fmt.Sprintf("variantValue_%s", normalizeName(variantName))
}

func variantInterfaceMethodName(variantName string) string {
	return fmt.Sprintf("isVariant_%s", normalizeName(variantName))
}

func variantCaseTypeName(variantName, caseName string) string {
	return fmt.Sprintf("%s%s", normalizeName(variantName), normalizeName(caseName))
}
