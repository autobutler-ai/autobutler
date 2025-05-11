import sys

import autobutler.llm as llm


def usage():
    print("Usage: butlerctl", file=sys.stderr)
    print("  --help            Show this help message", file=sys.stderr)


def main() -> int:
    is_running = True
    if "--help" in sys.argv:
        usage()
        is_running = False
    model = llm.LLM()
    while is_running:
        query = input("Ask AutoButler something...")
        if query.lower() == "exit":
            is_running = False
            continue
        question = question.strip()
        response = model.chat(llm.ChatRequest(prompt=question))
        print(response)
    print("It was a joy serving you. :)")
    return 0
