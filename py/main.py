import requests
import os


def dl_puzzle_input(filename):
    if os.path.exists(filename):
        return
    url = "https://codingmystery.com/assets/puzzle-input/the-beginning/" + filename
    r = requests.get(url)
    with open(filename, "w+", encoding="utf-8") as out:
        out.write(r.content)


def main():
    print("https://codingmystery.com/the-beginning")
    dl_puzzle_input("blank-sheet-of-paper.txt")
    dl_puzzle_input("shredded-sheet-of-paper.txt")
