# Rover Challenge

## Running and Testing

There are two versions of the challenge runner. `run_basic.py` which will print the output as described in the and `run_display.py` which will print additional information about the run.

Both take only a single argument which is the path to the file containing the mission plan to execute.

Running `./run_basic.py missions/test-input.txt` should produce the output expected via the challenge description.

All of the mission files should be executable with either runner.

### Testing

Tests exist in the [tests/](tests/) directory and utilize `unittest`. All the tests can be from the mars-rovers directory with: `python3 -m unittest discover -v -s tests/`

## Assumptions

Some assumptions were made about not only the execution environment, but also the implention.

### Execution Assumptions

- Python3 only. Development was done with version 3.8 specifically.
- Development and testing was done on a Linux desktop. Attempts were made to handle Windows line endings but the resources to test on Windows or OSX were not available. OSX should be a non-issue.

### Implementation Assumptions

- Driving or landing a rover off the bounds of a plateau is an error.
- Driving a rover into the same coordinate square where another rover resides (resting position after completing its run) is an error.
- The plateau size is configurable, but only on initial creation. Resizing an existing plateau, while implementible, is not currently possible.
- The plateau is not always a square but is always rectangular.
- Position X:0 Y:0 will ALWAYS be in the bottom left of the plateau. All negative position values are invalid.
- The number of rovers on a plateau is limited only by the number of available coordinate squares.
- The mission plan files may have Unix and/or non-Unix line endings as well as additional line endings (empty rows).

## Implementation

The implementation is comprised of three major components. Classes to represent Rovers and Plateaus as well as a function to parse 'mission plans'. A mission plan in this case is the text input as defined in the [challenge description](CHALLENGE.md).

### Mission Plan

The function [`parse_mission_plan()`](src/missionplan.py#L98) takes a single input; the text (challenge input) and returns an instance of a Plateau as well as a list of Rover instances. Running a rover requires landing a rover on the Plateau (calling the `land()` method on the rover instance passing in the Plateau instance to land on) and calling the `run()` method. This causes the rover to iterate over the instructions saved in an instance variable in the rover instance.

### Plateau

The Plateau class is _very simple_. Its primary responsibility is keeping a list of current rovers and making that information available to the other rovers that have landed on it. This information is required so rovers can detect collisions with one another.

### Rover

The Rover class is where most of the logic resides. It not only keeps track of where it is on the plateau (X and Y coordinates) but it also validates that it's still on the plateau after each move. As part of this validation, it also validates that it has not collided with any other rovers at rest on the plateau. If either of these happen, the appropriate exception is raised and the movement instruction parsing is stopped.

The rover class also contains an instance of a Compass class. The compass is a finite-state machine simply used to keep track of the rover's heading. This was broken out into its own class to keep the rover logic simplifier; however, in its current design, it makes testing purely the rover logic more difficult.