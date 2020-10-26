import unittest

import src.rover as rover
import src.plateau as plateau

import tests.tdata as tdata


class TestRover(unittest.TestCase):
    def test_rover_run(self):
        for rover_spec in tdata.valid_rovers:
            p = plateau.Plateau(
                rover_spec["tdata"]["plateau_x"], rover_spec["tdata"]["plateau_y"]
            )
            r = rover.Rover(
                rover_spec["tdata"]["x"],
                rover_spec["tdata"]["y"],
                rover_spec["tdata"]["heading"],
                rover_spec["tdata"]["instructions"],
            )

            r.land(p)
            self.assertEqual(r.run(), rover_spec["expected_path"])

    def test_rover_premature_start(self):
        r = rover.Rover(3, 3, "N", "MMR")
        self.assertRaises(rover.RoverNotLanded, r.run)

    def test_rover_off_plateau(self):
        p = plateau.Plateau(2, 2)
        for r in (
            rover.Rover(1, 1, "N", "MM"),
            rover.Rover(1, 1, "E", "MM"),
            rover.Rover(1, 1, "S", "MM"),
            rover.Rover(1, 1, "W", "MM"),
        ):
            r.land(p)
            self.assertRaises(rover.RoverOutOfBounds, r.run)

    def test_rover_bad_landing(self):
        p = plateau.Plateau(2, 2)
        for r in (
            rover.Rover(0, 3, "N", "MM"),
            rover.Rover(3, 1, "E", "MM"),
            rover.Rover(3, 3, "S", "MM"),
            rover.Rover(10000, 500, "W", "MM"),
        ):
            self.assertRaises(rover.RoverOutOfBounds, r.land, p)

    def test_rover_collision(self):
        p = plateau.Plateau(4, 4)

        r1 = rover.Rover(1, 0, "N", "MMRMRM")
        r2 = rover.Rover(3, 3, "S", "MMMRMRM")

        r1.land(p)
        r2.land(p)

        r1.run()

        with self.assertRaises(rover.RoverCollision) as err:
            r2.run()

        self.assertEqual(err.exception.x, 2)
        self.assertEqual(err.exception.y, 1)
