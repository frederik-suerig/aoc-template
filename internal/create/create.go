package create

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

var (
	//go:embed templates/solution.go.tmpl
	solutionTemplate string
	//go:embed templates/solution_test.go.tmpl
	solutionTestTemplate string
)

type Generator struct {
	day    int
	year   int
	cookie string

	workDir   string
	outputDir string
}

func NewGenerator(day, year int, workdir, cookie string) (*Generator, error) {
	g := &Generator{
		day:     day,
		year:    year,
		workDir: workdir,
		cookie:  cookie,
	}
	if err := g.init(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Generator) init() error {
	if g.day <= 0 || g.day > 25 {
		return fmt.Errorf("invalid day: %d", g.day)
	}
	if g.year <= 0 {
		return fmt.Errorf("invalid year: %d", g.year)
	}

	// From 2025 onwards, there are only 12 challenges per year.
	if g.year >= 2025 && g.day > 12 {
		return fmt.Errorf("invalid day: %d for year %d", g.day, g.year)
	}

	if g.workDir == "" {
		return fmt.Errorf("workdir is required")
	}

	g.outputDir = filepath.Join(
		g.workDir,
		fmt.Sprintf("y%04d", g.year),
		fmt.Sprintf("d%02d", g.day),
	)

	return nil
}

func (g *Generator) Run() error {
	if err := g.createFolderStructure(); err != nil {
		return fmt.Errorf("could not create folder structure: %w", err)
	}

	if err := g.renderTemplates(); err != nil {
		return fmt.Errorf("could not render templates: %w", err)
	}

	if g.cookie != "" {
		if err := g.downloadInput(); err != nil {
			return fmt.Errorf("could not download input: %w", err)
		}
	}

	return nil
}

func (g *Generator) createFolderStructure() error {
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("could not create directory: %w", err)
	}

	return nil
}

func (g *Generator) renderTemplates() error {
	if err := g.renderTemplate(solutionTemplate, "solution.go"); err != nil {
		return fmt.Errorf("could not render solution template: %w", err)
	}

	if err := g.renderTemplate(solutionTestTemplate, "solution_test.go"); err != nil {
		return fmt.Errorf("could not render solution test template: %w", err)
	}

	return nil
}

type templateData struct {
	Day  int
	Year int
}

func (g *Generator) renderTemplate(templateText, filename string) error {
	path := filepath.Join(g.outputDir, filename)

	if fileExists(path) {
		return fmt.Errorf("file already exists: %s", path)
	}

	tmpl, err := template.New(filename).Parse(templateText)
	if err != nil {
		return fmt.Errorf("could not parse template: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer f.Close()

	if err := tmpl.Execute(f, templateData{
		Day:  g.day,
		Year: g.year,
	}); err != nil {
		return fmt.Errorf("could not execute template: %w", err)
	}

	return nil
}

func (g *Generator) downloadInput() error {
	path := filepath.Join(g.outputDir, "testdata", "input.txt")
	if fileExists(path) {
		return fmt.Errorf("file already exists: %s", path)
	}

	if g.cookie == "" {
		return fmt.Errorf("cookie is required")
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", g.year, g.day)
	fmt.Println("Downloading input from:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request: %w", err)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: g.cookie})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not download input: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not download input: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("could not read response body: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("could not create directory: %w", err)
	}

	if err := os.WriteFile(path, body, 0644); err != nil {
		return fmt.Errorf("could not write input: %w", err)
	}

	return nil
}

// fileExists checks whether filename exists.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if err != nil {
		panic(err)
	}
	return !info.IsDir()
}
