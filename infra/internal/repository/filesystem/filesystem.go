package filesystem

import (
	"fmt"
	"github.com/biosvos/markadr/flow/adr"
	"github.com/biosvos/markadr/flow/repository"
	"github.com/pkg/errors"
	"os"
	"strings"
)

var _ repository.Repository = &Repository{}

type Repository struct {
	workDir string
}

func NewRepository(workDir string) *Repository {
	return &Repository{workDir: workDir}
}

func (r *Repository) Get(title string) (*adr.ADR, error) {
	contents, err := os.ReadFile(fmt.Sprintf("%v/%v.md", r.workDir, title))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	document, err := adr.NewDocument(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	section := adr.DivideSection(document)
	ret, err := adr.NewADR(section)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}

func (r *Repository) ListSummaries() ([]*adr.Summary, error) {
	files, err := os.ReadDir(r.workDir)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var ret []*adr.Summary
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasSuffix(file.Name(), ".md") {
			continue
		}

		summary, err := r.newSummary(file.Name())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret = append(ret, summary)
	}
	return ret, nil
}

func (r *Repository) newSummary(filename string) (*adr.Summary, error) {
	title := strings.TrimSuffix(filename, ".md")

	contents, err := os.ReadFile(fmt.Sprintf("%v/%v", r.workDir, filename))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	status, err := parseStatus(contents)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ret, err := adr.NewSummary(title, status)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}

func (r *Repository) RawData(title string) ([]byte, error) {
	contents, err := os.ReadFile(fmt.Sprintf("%v/%v.md", r.workDir, title))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return contents, nil
}
