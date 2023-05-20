// {{ headerComment }}

package master

import (
	"context"
	{{ range .MasterMessages }}
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/{{ lower .Name }}"
	{{- end }}
	pb "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/service"
)

type handler struct {
	{{- range .MasterMessages }}
	{{ lowerCamel .Name }}Repository {{ lower .Name }}.Repository
	{{- end }}
}

func New(
	{{- range .MasterMessages }}
	{{ lowerCamel .Name }}Repository {{ lower .Name }}.Repository,
	{{- end }}
) pb.MasterServiceServer {
	return &handler{
		{{- range .MasterMessages }}
		{{ lowerCamel .Name }}Repository: {{ lowerCamel .Name }}Repository,
		{{- end }}
	}
}

{{ range .Methods }}
func (h *handler) {{ .Name }}(ctx context.Context, _ *pb.Empry) (*pb.{{ .OutputMessage.Name }}, error) {
	{{- range .OutputMessage.Fields }}
	{{ lowerCamel .Name }}, err := h.{{ lowerCamel .Type }}Repository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}
	{{- end }}

	return &pb.{{ .OutputMessage.Name }}{
		{{- range .OutputMessage.Fields }}
		{{ .Name }}: to{{ .Type }}MasterList({{ lowerCamel .Name }}),
		{{- end }}
	}, nil
}
{{ end }}