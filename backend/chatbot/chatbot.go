package chatbot

import (
	"backend/gemini"
	"backend/portfolio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type RequestBody struct {
	Balance               int    `json:"balance"`
	Experience            string `json:"experience"`
	Preference            string `json:"preference"`
	Liquidity             string `json:"liquidity"`
	RiskBearing           string `json:"risk_bearing"`
	MinimumFreezingPeriod int    `json:"minimum_freezing_period"`
	QueryType             string `json:"query_type"`
}

func GenerateResponse(req RequestBody) (string, error) {
	// Simulate fetching a user's portfolio
	userPortfolio := portfolio.Portfolio{
		Balance:               req.Balance,
		Experience:            req.Experience,
		Preference:            req.Preference,
		Liquidity:             req.Liquidity,
		RiskBearing:           req.RiskBearing,
		MinimumFreezingPeriod: req.MinimumFreezingPeriod,
	}

	// Load prompt template
	promptPath := getPromptPath("prompt.txt")
	promptBytes, err := os.ReadFile(promptPath)
	if err != nil {
		return "", fmt.Errorf("failed to load prompt template from %s: %v", promptPath, err)
	}

	// Fill prompt template with user's data
	prompt := strings.ReplaceAll(string(promptBytes), "{balance}", fmt.Sprintf("%d", userPortfolio.Balance))
	prompt = strings.ReplaceAll(prompt, "{experience}", userPortfolio.Experience)
	prompt = strings.ReplaceAll(prompt, "{preference}", userPortfolio.Preference)
	prompt = strings.ReplaceAll(prompt, "{liquidity}", userPortfolio.Liquidity)
	prompt = strings.ReplaceAll(prompt, "{risk_bearing}", userPortfolio.RiskBearing)
	prompt = strings.ReplaceAll(prompt, "{minimum_freezing_period}", fmt.Sprintf("%d", userPortfolio.MinimumFreezingPeriod))

	// Send the filled prompt to Gemini
	return gemini.QueryGemini(prompt)
}

func Analyze_Portfolio(req map[string]interface{}) (string, error) {

	promptPath := getPromptPath("analyze_prompt.txt")
	prompt, err := os.ReadFile(promptPath)
	if err != nil {
		return "", fmt.Errorf("failed to load prompt template from %s: %v", promptPath, err)
	}
	fmt.Println(req)
	prompt = append(prompt, []byte(fmt.Sprintf("%+v", req))...)

	response, err := gemini.QueryGemini(string(prompt))
	fmt.Println(response)

	return response, err

}

// getPromptPath searches for prompt files in multiple possible locations
func getPromptPath(filename string) string {
	// Try current directory first
	if _, err := os.Stat(filename); err == nil {
		return filename
	}

	// Try backend directory (for when running from project root)
	backendPath := filepath.Join("backend", filename)
	if _, err := os.Stat(backendPath); err == nil {
		return backendPath
	}

	// Try absolute path from executable location
	execPath, err := os.Executable()
	if err == nil {
		execDir := filepath.Dir(execPath)
		absPath := filepath.Join(execDir, filename)
		if _, err := os.Stat(absPath); err == nil {
			return absPath
		}
	}

	// Return original filename as fallback (will trigger error in caller)
	return filename
}
