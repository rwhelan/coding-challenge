class CompassError(Exception):
    pass


class Compass:
    def __init__(self, initial_direction):
        self.ALL_DIRECTIONS = ["N", "E", "S", "W"]

        try:
            self.direction_idx = self.ALL_DIRECTIONS.index(initial_direction.upper())

        except ValueError:
            raise CompassError(
                f"Bad initial direction value of {initial_direction} not in known values of {self.ALL_DIRECTIONS}"
            )

        except AttributeError:
            raise CompassError(
                f"Bad initial direction value type: {type(initial_direction)}"
            )

    def turn_right(self):
        if self.direction_idx == len(self.ALL_DIRECTIONS) - 1:
            self.direction_idx = 0

        else:
            self.direction_idx += 1

        return self.ALL_DIRECTIONS[self.direction_idx]

    def turn_left(self):
        if self.direction_idx == 0:
            self.direction_idx = len(self.ALL_DIRECTIONS) - 1

        else:
            self.direction_idx -= 1

        return self.ALL_DIRECTIONS[self.direction_idx]

    @property
    def heading(self):
        return self.ALL_DIRECTIONS[self.direction_idx]


class RoverError(Exception):
    def __init__(self, message="Rover Error"):
        self.message = message

    def __str__(self):
        return self.message


class RoverOutOfBounds(RoverError):
    def __init__(self, message="Rover off plateau"):
        super(RoverOutOfBounds, self).__init__(message)


class RoverNotLanded(RoverError):
    def __init__(self, message="Rover not on the ground"):
        super(RoverNotLanded, self).__init__(message)


class RoverCollision(RoverError):
    def __init__(self, x, y):
        self.x, self.y = x, y
        super(RoverCollision, self).__init__(f"Rovers have collided at X:{x} Y:{y}")


class Rover:
    def __init__(self, x, y, direction, instructions):
        self.x = x
        self.y = y
        self.instructions = instructions
        self.compass = Compass(direction)
        self.path_taken = []
        self.plateau = None

    def _validate_position(self):
        if self.x > self.plateau.x:
            raise RoverOutOfBounds("Off plateau to the east")

        if self.x < 0:
            raise RoverOutOfBounds("Off plateau to the west")

        if self.y > self.plateau.y:
            raise RoverOutOfBounds("Off plateau to the north")

        if self.y < 0:
            raise RoverOutOfBounds("Off plateau to the south")

        for plateau_rover in self.plateau.all_rovers():
            if (
                plateau_rover.x == self.x
                and plateau_rover.y == self.y
                and plateau_rover is not self
            ):
                raise RoverCollision(self.x, self.y)

    def _move_forward(self):
        # fmt: off
        def m_north(): self.y += 1
        def m_east():  self.x += 1
        def m_south(): self.y -= 1
        def m_west():  self.x -= 1

        directional_moves = {
            "N": m_north,
            "E": m_east,
            "S": m_south,
            "W": m_west,
        }
        # fmt: on

        directional_moves[self.compass.heading]()

    def land(self, plateau):
        self.plateau = plateau
        self.plateau.add_rover(self)
        self._validate_position()

    def run(self):
        self.path_taken = [(self.x, self.y, self.compass.heading)]

        if not self.plateau:
            raise RoverNotLanded

        ops = {
            "M": self._move_forward,
            "R": self.compass.turn_right,
            "L": self.compass.turn_left,
        }

        for i in self.instructions:
            ops[i.upper()]()
            self.path_taken.append((self.x, self.y, self.compass.heading))
            self._validate_position()

        return self.path_taken
