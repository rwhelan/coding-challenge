import unittest

import src.rover as rover


class TestCompass(unittest.TestCase):
    def test_compass_turn_right(self):
        compass = rover.Compass("N")
        self.assertEqual(compass.turn_right(), "E")
        self.assertEqual(compass.turn_right(), "S")
        self.assertEqual(compass.turn_right(), "W")
        self.assertEqual(compass.turn_right(), "N")
        self.assertEqual(compass.turn_right(), "E")

    def test_compass_turn_left(self):
        compass = rover.Compass("N")
        self.assertEqual(compass.turn_left(), "W")
        self.assertEqual(compass.turn_left(), "S")
        self.assertEqual(compass.turn_left(), "E")
        self.assertEqual(compass.turn_left(), "N")
        self.assertEqual(compass.turn_left(), "W")

    def test_compass_bad_initial_value(self):
        self.assertRaises(rover.CompassError, rover.Compass, "F")
        self.assertRaises(rover.CompassError, rover.Compass, 3)
