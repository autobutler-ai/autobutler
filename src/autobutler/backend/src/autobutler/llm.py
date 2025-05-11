import sys
from pprint import pprint

import autobutler.config as config
import torch
from pydantic import BaseModel, Field
from transformers import AutoModelForCausalLM, AutoTokenizer


class ChatRequest(BaseModel):
    prompt: str = Field(...)


class LLM:
    def __init__(self):
        """
        Initialize the model and tokenizer.
        This function is called when the module is imported.
        """
        print("Initializing TinyLlama model...")
        self.tokenizer = AutoTokenizer.from_pretrained(
            "TinyLlama/TinyLlama-1.1B-Chat-v1.0"
        )
        print("Loading TinyLlama model...")
        self.model = AutoModelForCausalLM.from_pretrained(
            "TinyLlama/TinyLlama-1.1B-Chat-v1.0"
        )
        if torch.cuda.is_available():
            print("GPU is available. Using GPU for inference.")
            self.device = torch.device("cuda")
        elif torch.backends.mps.is_available():
            print("Metal Performance Shaders is available. Using MPS for inference.")
            self.device = torch.device("mps")
        else:
            print("GPU not available. Using CPU for inference.")
            self.device = torch.device("cpu")
        self.model = self.model.to(self.device)
        print("LLM loaded with the following config: ")
        pprint(config.LLM)

    def chat(self, request: ChatRequest) -> str:
        try:
            print("Preparing input...")
            prompt = config.LLM.PROMPT.format(
                **{
                    "context": "Be cordial, succinct, and helpful. Do not dump huge walls of text.",
                    "prompt": request.prompt.strip(),
                }
            )
            print("Tokenizing input...")
            inputs = self.tokenizer(prompt, return_tensors="pt").to(self.device)

            # Generate response
            print("Generating response...")
            outputs = self.model.generate(
                **inputs,
                max_new_tokens=config.LLM.MAX_TOKENS,
                do_sample=True,
                top_p=config.LLM.TOP_P,
                top_k=config.LLM.TOP_K,
                temperature=config.LLM.TEMPERATURE,
                num_beams=config.LLM.NUM_BEAMS,
                pad_token_id=self.tokenizer.eos_token_id,
            )

            print("Decoding response...")
            response = self.tokenizer.decode(outputs[0], skip_special_tokens=True)

            # Remove the prompt from the response
            response = response[len(prompt) :].strip()
            return response
        except Exception as e:
            print(e, file=sys.stderr)
            return "I'm sorry, I couldn't process your request. Please try again later."
