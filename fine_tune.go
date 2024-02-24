package main

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

func fine_tune() {
	client := openai.NewClient("sk-3WxBIzVgRdljWdgu4cNa8i3XHgVlZlRWMwTbvmEgfezaqG4G")
	ctx := context.Background()

	// create a .jsonl file with your training data for conversational model
	// {"prompt": "<prompt text>", "completion": "<ideal generated text>"}
	// {"prompt": "<prompt text>", "completion": "<ideal generated text>"}
	// {"prompt": "<prompt text>", "completion": "<ideal generated text>"}

	// chat models are trained using the following file format:
	// {"messages": [{"role": "system", "content": "Marv is a factual chatbot that is also sarcastic."}, {"role": "user", "content": "What's the capital of France?"}, {"role": "assistant", "content": "Paris, as if everyone doesn't know that already."}]}
	// {"messages": [{"role": "system", "content": "Marv is a factual chatbot that is also sarcastic."}, {"role": "user", "content": "Who wrote 'Romeo and Juliet'?"}, {"role": "assistant", "content": "Oh, just some guy named William Shakespeare. Ever heard of him?"}]}
	// {"messages": [{"role": "system", "content": "Marv is a factual chatbot that is also sarcastic."}, {"role": "user", "content": "How far is the Moon from Earth?"}, {"role": "assistant", "content": "Around 384,400 kilometers. Give or take a few, like that really matters."}]}

	// you can use openai cli tool to validate the data
	// For more info - https://platform.openai.com/docs/guides/fine-tuning

	file, err := client.CreateFile(ctx, openai.FileRequest{
		FilePath: "D:/Project/src/gpt-fofa/training_prepared.jsonl",
		Purpose:  "fine-tune",
	})
	if err != nil {
		fmt.Printf("Upload JSONL file error: %v\n", err)
		return
	}

	// create a fine tuning job
	// Streams events until the job is done (this often takes minutes, but can take hours if there are many jobs in the queue or your dataset is large)
	// use below get method to know the status of your model
	fineTuningJob, err := client.CreateFineTuningJob(ctx, openai.FineTuningJobRequest{
		TrainingFile: file.ID,
		Model:        "fofa-gpt-002", // gpt-3.5-turbo-0613, babbage-002.
	})
	if err != nil {
		fmt.Printf("Creating new fine tune model error: %v\n", err)
		return
	}

	fineTuningJob, err = client.RetrieveFineTuningJob(ctx, fineTuningJob.ID)
	if err != nil {
		fmt.Printf("Getting fine tune model error: %v\n", err)
		return
	}
	fmt.Println(fineTuningJob.FineTunedModel)

	// once the status of fineTuningJob is `succeeded`, you can use your fine tune model in Completion Request or Chat Completion Request

	// resp, err := client.CreateCompletion(ctx, openai.CompletionRequest{
	//	 Model:  fineTuningJob.FineTunedModel,
	//	 Prompt: "your prompt",
	// })
	// if err != nil {
	//	 fmt.Printf("Create completion error %v\n", err)
	//	 return
	// }
	//
	// fmt.Println(resp.Choices[0].Text)
}
