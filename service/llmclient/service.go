package llmclient

import (
	"encoding/json"
)

func ParseEmailAndName(text string) (ResumeProfile, error) {
	geminRequest := GeminiRequestBody{
		SystemInstruction: map[string]interface{}{
			"parts": map[string]string{
				"text": ExtractEmailAndNameSystemMessage,
			},
		},
		Contents: []Message{
			{
				Role: USER,
				Parts: []Part{
					{
						Text: text,
					},
				},
			},
		},
		GenerationConfig: GenerationConfig{
			ResponseMimeType: "application/json",
			ResponseSchema: map[string]interface{}{
				"type":"object",
				"properties": map[string]interface{}{
					"email": map[string]string{
						"type": "string",
					},
					"name": map[string]string{
						"type": "string",
					},
				},
			},
		},

	}
	res, err := GeminiClient(geminRequest)
	if err != nil {
		return ResumeProfile{}, err
	}
	processedEmail := res.Candidates[0].Content.Parts[0].Text

	var profile ResumeProfile;
	err = json.Unmarshal([]byte(processedEmail), &profile)
	if err != nil {
		return ResumeProfile{}, err
	}
	return profile, nil
}