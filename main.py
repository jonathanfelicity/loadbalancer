import random
from collections import defaultdict

class LoadBalancer:
    def __init__(self):
        self.servers = []
        self.server_weights = defaultdict(int)

    def add_server(self, server, weight=1):
        self.servers.append(server)
        self.server_weights[server] = weight

    def remove_server(self, server):
        if server in self.servers:
            self.servers.remove(server)
            del self.server_weights[server]

    def get_server(self):
        total_weight = sum(self.server_weights.values())
        random_num = random.uniform(0, total_weight)
        current_weight = 0

        for server, weight in self.server_weights.items():
            current_weight += weight
            if random_num <= current_weight:
                return server

        # In case no server is selected, return None
        return None

# Example usage
lb = LoadBalancer()
lb.add_server("Server A", weight=3)
lb.add_server("Server B", weight=2)
lb.add_server("Server C", weight=1)

for _ in range(10):
    server = lb.get_server()
    print("Request sent to:", server)
