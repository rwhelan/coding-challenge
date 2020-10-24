#!/usr/bin/env python3

import sys

import src.missionplan as missionplan

def display_board(board):
    for r in board:
        print("|", " ".join(r), "|")

def print_path(plateau, path):
    direction_indicators = {
        "N": "^",
        "E": ">",
        "S": "v",
        "W": "<",
    }

    board = []
    for _ in range(plateau.y):
        board.append([' '] * (plateau.x+1))

    r = [i for i in range(plateau.y -1,-1, -1)]
    for i in path:
        board[r[i[1]]][i[0]] = direction_indicators[i[2]]

    display_board(board)

def run_rovers(mission_plan):
    plateau, rovers = missionplan.parse_mission_plan(mission_plan)

    for rover in rovers:
        print(f"\nStarted at: X:{rover.x} Y:{rover.y} Heading:{rover.compass.heading}")
        rover.land(plateau)
        path = rover.run()

        print_path(plateau, path)
        print(f"Ended at: X:{rover.x} Y:{rover.y} Heading:{rover.compass.heading}\n")



if __name__ == "__main__":
    if len(sys.argv) <= 1:
        print(f"Usage: {sys.argv[0]} <mission spec file>")
        sys.exit(1)

    mission_plan = open(sys.argv[1], 'r').read()
    run_rovers(mission_plan)

 