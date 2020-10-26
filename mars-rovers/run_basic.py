#!/usr/bin/env python3

import sys

import src.missionplan as missionplan


def run_rovers(mission_plan):
    plateau, rovers = missionplan.parse_mission_plan(mission_plan)

    for rover in rovers:
        rover.land(plateau)
        path = rover.run()

        print("{} {} {}\n".format(*path[-1]))


if __name__ == "__main__":
    if len(sys.argv) <= 1:
        print(f"Usage: {sys.argv[0]} <mission spec file>")
        sys.exit(1)

    mission_plan = open(sys.argv[1], "r").read()
    run_rovers(mission_plan)
