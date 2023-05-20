// {{ headerComment }}
package master

import (
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/master"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/enums"
	pbmaster "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/master"
)

{{- range .MasterMessages }}
func to{{ .Name }}MasterList(lst master.{{ .Name }}Slice) []*pbmaster.{{ .Name }} {
	result := make([]*pbmaster.{{ .Name }}, 0, len(lst))
	for _, v := range lst {
		result = append(result, &pbmaster.{{ .Name }}{
            {{- range .Fields }}
            {{- if .IsEnum }}
            {{ .Name }}: enums.{{ .Type }}(v.{{ goFieldName .Name }}),
            {{- else }}
            {{ .Name }}: v.{{ goFieldName .Name }},
            {{- end }}
            {{- end }}
		})
	}

	return result
}
{{ end }}