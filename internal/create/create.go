package create

import (
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/frederik-suerig/advent-of-code/internal/ui"
)

var (
	//go:embed templates/solution.go.tmpl
	solutionTemplate string
	//go:embed templates/solution_test.go.tmpl
	solutionTestTemplate string
)

// Config holds the configuration for creating a new Advent of Code challenge
type Config struct {
	Year    int
	Day     int
	WorkDir string
	Cookie  string
}

type Generator struct {
	day    int
	year   int
	cookie string

	workDir   string
	outputDir string
}

func NewGenerator(cfg Config) (*Generator, error) {
	g := &Generator{
		day:     cfg.Day,
		year:    cfg.Year,
		workDir: cfg.WorkDir,
		cookie:  cfg.Cookie,
	}
	if err := g.init(); err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Generator) init() error {
	if g.day <= 0 || g.day > 25 {
		return fmt.Errorf("%w: %d", ErrInvalidDay, g.day)
	}
	if g.year <= 0 {
		return fmt.Errorf("%w: %d", ErrInvalidYear, g.year)
	}

	// From 2025 onwards, there are only 12 challenges per year.
	if g.year >= 2025 && g.day > 12 {
		return fmt.Errorf("%w: %d for year %d", ErrInvalidDay, g.day, g.year)
	}

	if g.workDir == "" {
		return ErrWorkdirRequired
	}

	g.outputDir = filepath.Join(
		g.workDir,
		fmt.Sprintf("y%04d", g.year),
		fmt.Sprintf("d%02d", g.day),
	)

	return nil
}

func (g *Generator) Run() error {
	ui.Header("Creating Advent of Code %d - Day %d", g.year, g.day)

	// Check if directory or files already exist
	if g.directoryOrFilesExist() {
		relPath := ui.MakeRelative(g.outputDir)
		shouldDelete, err := ui.ConfirmOverwrite(g.year, g.day, relPath)
		if err != nil {
			return err
		}
		if !shouldDelete {
			return fmt.Errorf("operation cancelled by user")
		}
		if err := g.deleteDirectory(); err != nil {
			return fmt.Errorf("failed to delete existing directory: %w", err)
		}
	}

	if err := g.createFolderStructure(); err != nil {
		return err
	}

	if err := g.renderTemplates(); err != nil {
		return err
	}

	if g.cookie != "" {
		if err := g.downloadInput(); err != nil {
			return err
		}
	} else {
		ui.Warning("No cookie provided - skipping input download")
		ui.DimText("  You can download the input manually or provide a cookie with --cookie")
	}

	return nil
}

func (g *Generator) createFolderStructure() error {
	// Check if directory already exists
	dirExists := false
	if info, err := os.Stat(g.outputDir); err == nil && info.IsDir() {
		dirExists = true
	}

	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Only show success message if we actually created the directory
	if !dirExists {
		ui.DirCreated(g.outputDir)
	}
	return nil
}

func (g *Generator) renderTemplates() error {
	if err := g.renderTemplate(solutionTemplate, "solution.go"); err != nil {
		return err
	}

	if err := g.renderTemplate(solutionTestTemplate, "solution_test.go"); err != nil {
		return err
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
		return NewFileExistsError(path)
	}

	tmpl, err := template.New(filename).Parse(templateText)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			// Log but don't fail if close fails after successful write
			_, _ = fmt.Fprintf(os.Stderr, "warning: failed to close file %s: %v\n", path, closeErr)
		}
	}()

	if err := tmpl.Execute(f, templateData{
		Day:  g.day,
		Year: g.year,
	}); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	ui.FileCreated(path)
	return nil
}

func (g *Generator) downloadInput() error {
	path := filepath.Join(g.outputDir, "testdata", "input.txt")
	if fileExists(path) {
		return NewFileExistsError(path)
	}

	if g.cookie == "" {
		return ErrCookieRequired
	}

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", g.year, g.day)
	ui.Download("Downloading input from adventofcode.com")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return NewDownloadError(fmt.Sprintf("failed to create request: %v", err), 0)
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: g.cookie})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return NewDownloadError(fmt.Sprintf("network error: %v", err), 0)
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		// Provide more specific error messages based on status code
		switch resp.StatusCode {
		case http.StatusUnauthorized, http.StatusForbidden:
			return NewDownloadError("authentication failed - check your session cookie", resp.StatusCode)
		case http.StatusNotFound:
			return NewDownloadError("input not available - puzzle may not be released yet", resp.StatusCode)
		default:
			return NewDownloadError(resp.Status, resp.StatusCode)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if closeErr := resp.Body.Close(); closeErr != nil {
		return fmt.Errorf("failed to close response body: %w", closeErr)
	}
	if err != nil {
		return NewDownloadError(fmt.Sprintf("failed to read response: %v", err), 0)
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(path, body, 0644); err != nil {
		return fmt.Errorf("failed to write input file: %w", err)
	}

	ui.FileCreated(path)
	return nil
}

// directoryOrFilesExist checks if the output directory or any expected files exist
func (g *Generator) directoryOrFilesExist() bool {
	// Check if directory exists
	if info, err := os.Stat(g.outputDir); err == nil && info.IsDir() {
		return true
	}

	// Check if any of the expected files exist
	expectedFiles := []string{
		filepath.Join(g.outputDir, "solution.go"),
		filepath.Join(g.outputDir, "solution_test.go"),
		filepath.Join(g.outputDir, "testdata", "input.txt"),
	}

	for _, file := range expectedFiles {
		if fileExists(file) {
			return true
		}
	}

	return false
}

// deleteDirectory removes the entire output directory
func (g *Generator) deleteDirectory() error {
	if err := os.RemoveAll(g.outputDir); err != nil {
		return fmt.Errorf("failed to remove directory: %w", err)
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
