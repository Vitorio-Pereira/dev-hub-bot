package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vitorio-Pereira/dev-hub-bot/internal/project"
)

type RepoInfo struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URL         string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type Client struct {
	token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetRepo(ctx context.Context, owner string, repo string) (RepoInfo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return RepoInfo{}, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RepoInfo{}, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()

	var info RepoInfo

	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return RepoInfo{}, fmt.Errorf("decoding response: %w", err)
	}
	return info, nil
}

func (c *Client) GetLanguages(ctx context.Context, owner string, repo string) (map[string]int, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", owner, repo)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}
	defer resp.Body.Close()
	var languages map[string]int
	if err := json.NewDecoder(resp.Body).Decode(&languages); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	return languages, nil
}

func ClassifyLanguages(langs map[string]int) project.Stack {
	var stack project.Stack

	for lang := range langs {
		switch lang {
		case "Go", "Python", "TypeScript", "JavaScript", "Rust", "Java", "C", "C++", "C#", "Ruby", "PHP", "Kotlin", "Swift", "Lua", "HTML", "CSS":
			stack.Languages = append(stack.Languages, lang)
		case "Dockerfile", "HCL", "Shell", "Makefile", "PowerShell", "Jinja", "Smarty", "YAML", "Procfile":
			stack.Infra = append(stack.Infra, lang)
		default:
			stack.Other = append(stack.Other, lang)
		}
	}

	return stack
}
