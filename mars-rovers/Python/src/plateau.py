class Plateau:
    def __init__(self, x, y):
        self.x = x
        self.y = y
        self.rovers = []

    def add_rover(self, rover):
        self.rovers.append(rover)

    def all_rovers(self):
        return self.rovers
