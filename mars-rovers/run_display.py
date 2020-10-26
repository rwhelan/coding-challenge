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
    for _ in range(plateau.y + 1):
        board.append([" "] * (plateau.x + 1))

    for i in path:
        board[i[1]][i[0]] = direction_indicators[i[2]]

    board.reverse()
    display_board(board)


def run_rovers(mission_plan):
    plateau, rovers = missionplan.parse_mission_plan(mission_plan)

    for rover in rovers:
        print(f"\nStarted: X:{rover.x} Y:{rover.y} Heading:{rover.compass.heading}")
        rover.land(plateau)
        path = rover.run()

        print_path(plateau, path)
        print(
            f"Ended: X:{rover.x} Y:{rover.y} Heading:{rover.compass.heading} Commands:{len(path)}\n"
        )


if __name__ == "__main__":
    if len(sys.argv) <= 1:
        print(f"Usage: {sys.argv[0]} <mission spec file>")
        sys.exit(1)

    mission_plan = open(sys.argv[1], "r").read()
    run_rovers(mission_plan)
