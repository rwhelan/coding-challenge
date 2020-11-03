valid_plateau_specs = [
    {"test_label": "Provided", "test_data": "5 5\n1 2 N\n", "x": 5, "y": 5,},
    {
        "test_label": "Provided w/ Windows EOL",
        "test_data": "5 50\r\n1 2 N\r\n",
        "x": 5,
        "y": 50,
    },
]

invalid_plateau_specs = [
    {"test_label": "non-numeric", "test_data": "5 d\n1 2 N\n"},
    {"test_label": "to many values", "test_data": "5 5 5\r\n1 2 N\r\n"},
]

valid_rover_specs = [
    {
        "test_label": "Provided",
        "test_data": "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\nMMRMMRMRRM",
        "rovers": [
            {"x": 1, "y": 2, "direction": "N", "instructions": "LMLMLMLMM",},
            {"x": 3, "y": 3, "direction": "E", "instructions": "MMRMMRMRRM",},
        ],
    },
    {
        "test_label": "Provided Extra Returns",
        "test_data": "5 5\n\n1 2 N\n\nLMLMLMLMM\n\n3 3 E\n\nMMRMMRMRRM\n\n",
        "rovers": [
            {"x": 1, "y": 2, "direction": "N", "instructions": "LMLMLMLMM",},
            {"x": 3, "y": 3, "direction": "E", "instructions": "MMRMMRMRRM",},
        ],
    },
    {
        "test_label": "Provided Extra Returns and Windows EOLs",
        "test_data": "5 5\r\n\r\n1 2 S\r\n\r\nLMLMLMLMR\r\n\r\n3 3 E\r\n\r\nMMRMMRMRRM\r\n\r\n",
        "rovers": [
            {"x": 1, "y": 2, "direction": "S", "instructions": "LMLMLMLMR",},
            {"x": 3, "y": 3, "direction": "E", "instructions": "MMRMMRMRRM",},
        ],
    },
]

invalid_rover_specs = [
    {
        "test_label": "Missing Directions",
        "test_data": "5 5\n1 2 N\nLMLMLMLMM\n3 3 E\n",
    },
    {
        "test_label": "Bad Init Coordinates",
        "test_data": "5 5\n1 D N\nLMLMLMLMM\n3 3 E\nMMRMMRMRDM\n",
    },
]

valid_rovers = [
    {
        "tdata": {
            "x": 3,
            "y": 3,
            "heading": "E",
            "instructions": "MMRMMRMRRM",
            "plateau_x": 5,
            "plateau_y": 5,
        },
        "expected_path": [
            (3, 3, "E"),
            (4, 3, "E"),
            (5, 3, "E"),
            (5, 3, "S"),
            (5, 2, "S"),
            (5, 1, "S"),
            (5, 1, "W"),
            (4, 1, "W"),
            (4, 1, "N"),
            (4, 1, "E"),
            (5, 1, "E"),
        ],
    },
    {
        "tdata": {
            "x": 1,
            "y": 2,
            "heading": "N",
            "instructions": "LMLMLMLMM",
            "plateau_x": 5,
            "plateau_y": 5,
        },
        "expected_path": [
            (1, 2, "N"),
            (1, 2, "W"),
            (0, 2, "W"),
            (0, 2, "S"),
            (0, 1, "S"),
            (0, 1, "E"),
            (1, 1, "E"),
            (1, 1, "N"),
            (1, 2, "N"),
            (1, 3, "N"),
        ],
    },
    {
        "tdata": {
            "x": 0,
            "y": 0,
            "heading": "N",
            "instructions": "MMMMMMMMMMRMRMMMMMMMMMMLMLMMMMMMMMMM",
            "plateau_x": 2,
            "plateau_y": 10,
        },
        "expected_path": [
            (0, 0, "N"),
            (0, 1, "N"),
            (0, 2, "N"),
            (0, 3, "N"),
            (0, 4, "N"),
            (0, 5, "N"),
            (0, 6, "N"),
            (0, 7, "N"),
            (0, 8, "N"),
            (0, 9, "N"),
            (0, 10, "N"),
            (0, 10, "E"),
            (1, 10, "E"),
            (1, 10, "S"),
            (1, 9, "S"),
            (1, 8, "S"),
            (1, 7, "S"),
            (1, 6, "S"),
            (1, 5, "S"),
            (1, 4, "S"),
            (1, 3, "S"),
            (1, 2, "S"),
            (1, 1, "S"),
            (1, 0, "S"),
            (1, 0, "E"),
            (2, 0, "E"),
            (2, 0, "N"),
            (2, 1, "N"),
            (2, 2, "N"),
            (2, 3, "N"),
            (2, 4, "N"),
            (2, 5, "N"),
            (2, 6, "N"),
            (2, 7, "N"),
            (2, 8, "N"),
            (2, 9, "N"),
            (2, 10, "N"),
        ],
    },
]
