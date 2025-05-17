package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/atotto/clipboard"
)

type OpenAIResponse struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	CreatedAt         int    `json:"created_at"`
	Status            string `json:"status"`
	Error             any    `json:"error"`
	IncompleteDetails any    `json:"incomplete_details"`
	Instructions      any    `json:"instructions"`
	MaxOutputTokens   any    `json:"max_output_tokens"`
	Model             string `json:"model"`
	Output            []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Status  string `json:"status"`
		Content []struct {
			Type        string `json:"type"`
			Annotations []any  `json:"annotations"`
			Text        string `json:"text"`
		} `json:"content"`
		Role string `json:"role"`
	} `json:"output"`
	ParallelToolCalls  bool `json:"parallel_tool_calls"`
	PreviousResponseID any  `json:"previous_response_id"`
	Reasoning          struct {
		Effort  any `json:"effort"`
		Summary any `json:"summary"`
	} `json:"reasoning"`
	ServiceTier string  `json:"service_tier"`
	Store       bool    `json:"store"`
	Temperature float64 `json:"temperature"`
	Text        struct {
		Format struct {
			Type string `json:"type"`
		} `json:"format"`
	} `json:"text"`
	ToolChoice string  `json:"tool_choice"`
	Tools      []any   `json:"tools"`
	TopP       float64 `json:"top_p"`
	Truncation string  `json:"truncation"`
	Usage      struct {
		InputTokens        int `json:"input_tokens"`
		InputTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"input_tokens_details"`
		OutputTokens        int `json:"output_tokens"`
		OutputTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"output_tokens_details"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	User     any `json:"user"`
	Metadata struct {
	} `json:"metadata"`
}

func WriteToClipboard(content string) error {
	cmd := exec.Command("xclip", "-selection", "clipboard")
	in, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		// Fallback to xsel
		cmd = exec.Command("xsel", "--clipboard", "--input")
		in, err = cmd.StdinPipe()
		if err != nil {
			return err
		}
		if err := cmd.Start(); err != nil {
			return err
		}
	}

	_, err = io.WriteString(in, content)
	if err != nil {
		return err
	}

	if err := in.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}

var client = &http.Client{}

func main() {
	args := os.Args
	// for i, v := range args {
	// fmt.Printf("arg[%d]: %s\n", i, v)
	// }

	token, foundToken := os.LookupEnv("OPENAI_API_KEY")
	if !foundToken || token == "" {
		log.Fatalln("OpenAI API key not found")
	}

	os := "linux"

	msg := strings.Join(args[1:], " ")
	prompt := fmt.Sprintf(
		"give me a %s command that %s. give me the command only as plain text, do not write any text other than the command itself, do not format with markdown",
		os,
		msg,
	)
	model := "gpt-4.1-mini"

	body, err := json.Marshal(map[string]string{
		"model": model,
		"input": prompt,
	})

	if err != nil {
		log.Fatalln("(×_×;) Error marshaling request body:", err.Error())
	}

	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/responses", payload)
	req.Header.Add("Content-type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	fmt.Println("(⌐■_■) thinking...")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln("(×_×;) Request error:", err.Error())
	}

	response_body_bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln("(×_×;) Error reading response body:", err.Error())
	}
	defer res.Body.Close()

	var response_body OpenAIResponse
	if err := json.Unmarshal(response_body_bytes, &response_body); err != nil {
		log.Fatalln("(×_×;) Error parsing response body:", err.Error())
	}

	result := response_body.Output[0].Content[0].Text

	fmt.Printf("(ง•̀_•́)ง %s\n", result)

	err = clipboard.WriteAll(result)
	if err == nil {
		fmt.Println("        ^ copied to clipboard")
	}
}
