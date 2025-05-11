import os
from dataclasses import dataclass

# Configuration
HA_URL = os.getenv(
    "HA_URL", "http://homeassistant.local:8123"
)  # Your Home Assistant URL
HA_TOKEN = os.getenv("HA_TOKEN", None)  # Your Home Assistant access token


@dataclass
class LLMConfig:
    PROMPT: str
    MAX_TOKENS: int
    TOP_P: float
    TOP_K: int
    TEMPERATURE: float
    NUM_BEAMS: int


LLM = LLMConfig(
    PROMPT="User: {prompt}",  # "System: You are AutoButler, an automated home assistant. {context}\n\nUser: {prompt}",
    MAX_TOKENS=int(os.getenv("LLM_MAX_TOKENS", 1024)),  # Max tokens for LLM response
    TOP_P=float(os.getenv("LLM_TOP_P", 0.95)),  # Top P sampling for LLM
    TOP_K=int(os.getenv("LLM_TOP_K", 50)),  # Top K sampling for LLM
    TEMPERATURE=float(os.getenv("LLM_TEMPERATURE", 0.7)),  # Temperature for LLM
    NUM_BEAMS=int(os.getenv("LLM_NUM_BEAMS", 1)),  # Number of beams for LLM
)
