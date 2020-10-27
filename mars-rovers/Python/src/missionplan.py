from .plateau import Plateau
from .rover import Rover


class InvalidMissionPlanFormat(Exception):
    def __init__(self, message="Invalid Mission Plan"):
        self.message = message


class InvalidPlateauSpec(InvalidMissionPlanFormat):
    def __init__(self, message="Invalid Plateau Specification"):
        super(InvalidPlateauSpec, self).__init__(message)


class InvalidRoverSpec(InvalidMissionPlanFormat):
    def __init__(self, message="Invalid Rover Specification"):
        super(InvalidRoverSpec, self).__init__(message)


def parse_plateau_spec(mission_plan_text):
    plateau_text = mission_plan_text.split("\n", 1)[0].strip("\r")

    try:
        raw_x, raw_y = plateau_text.split(" ")
    except ValueError:
        raise InvalidPlateauSpec("Wrong number of values")

    try:
        x = int(raw_x)
    except ValueError:
        raise InvalidPlateauSpec("Bad value for X; non-numatic")

    try:
        y = int(raw_y)
    except ValueError:
        raise InvalidPlateauSpec("Bad value for Y; non-numatic")

    return Plateau(x, y)


def parse_rover_init_coord(coord_text):
    try:
        raw_x, raw_y, direction = coord_text.split(" ")
    except ValueError:
        raise InvalidRoverSpec("Incorrect number of initial coordinate values")

    try:
        x = int(raw_x)
    except ValueError:
        raise InvalidRoverSpec("Bad initial value for X; non-numatic")

    try:
        y = int(raw_y)
    except ValueError:
        raise InvalidRoverSpec("Bad initial value for Y; non-numatic")

    if direction.upper() not in "NESW":
        raise InvalidRoverSpec(f'Bad value of "{direction}" supplied for heading')

    return (x, y, direction.upper())


def parse_rover_instructions(instructions_text):
    valid_char = lambda x: x.upper() in "LRM"

    for i in instructions_text:
        if not valid_char(i):
            raise InvalidRoverSpec(f'Bad instruction "{i}" specified')

    return instructions_text.upper()


def parse_rover_spec(mission_plan_text):
    rovers = []

    # Split on and clean all ctrl chars
    rovers_rows = [i.strip("\r") for i in mission_plan_text.split("\n") if i][1:]
    rovers_rows = [i for i in rovers_rows if i]

    rover_specs = [rovers_rows[i : i + 2] for i in range(0, len(rovers_rows), 2)]
    for spec in rover_specs:
        if len(spec) != 2:
            raise InvalidRoverSpec("Invalid Mission Plan; Rover Spec Malformed!")

    for spec in enumerate(rover_specs):
        try:
            x, y, direction = parse_rover_init_coord(spec[1][0])
            instructions = parse_rover_instructions(spec[1][1])
        except InvalidRoverSpec as e:
            e.message += f"; Check instruction {spec[0]}"
            raise e

        rovers.append(Rover(x, y, direction, instructions))

    return rovers


def parse_mission_plan(mission_plan_text):
    return (parse_plateau_spec(mission_plan_text), parse_rover_spec(mission_plan_text))
