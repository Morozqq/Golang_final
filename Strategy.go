class FlowerState:
    def water(self):
        pass

class DryState(FlowerState):
    def water(self):
        print("The flower is dry. Watering...")

class HealthyState(FlowerState):
    def water(self):
        print("The flower is healthy. No need to water.")

class Flower:
    def __init__(self):
        self.state = DryState()

    def water(self):
        self.state.water()

    def set_state(self, state):
        self.state = state
