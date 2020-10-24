import unittest

import src.missionplan as missionplan

import tests.tdata as tdata


class TestMissionPlanParsePlateau(unittest.TestCase):
    def test_parse_valid_plateau_spec(self):
        for spec in tdata.valid_plateau_specs:
            p = missionplan.parse_plateau_spec(spec["test_data"])
            self.assertEqual(p.x, spec["x"], "Invalid X value!")
            self.assertEqual(p.y, spec["y"], "Invalid Y value!")

    def test_parse_invalid_plateau_spec(self):
        for spec in tdata.invalid_plateau_specs:
            with self.assertRaises(
                missionplan.InvalidPlateauSpec, msg=spec["test_label"]
            ):
                missionplan.parse_plateau_spec(spec["test_data"])


class TestMissionPlanParseRover(unittest.TestCase):
    def test_parse_valid_rover_spec(self):
        for spec in tdata.valid_rover_specs:
            for i, r in enumerate(missionplan.parse_rover_spec(spec["test_data"])):
                self.assertEqual(r.x, spec["rovers"][i]["x"], "Invalid X value!")
                self.assertEqual(r.y, spec["rovers"][i]["y"], "Invalid Y value!")
                self.assertEqual(
                    r.instructions,
                    spec["rovers"][i]["instructions"],
                    "Invalid instructions!",
                )

    def test_parse_invalid_rover_spec(self):
        for spec in tdata.invalid_rover_specs:
            with self.assertRaises(
                missionplan.InvalidRoverSpec, msg=spec["test_label"]
            ):
                missionplan.parse_rover_spec(spec["test_data"])
