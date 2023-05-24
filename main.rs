use rand::Rng;
use std::collections::HashMap;

struct LoadBalancer {
    servers: Vec<String>,
    server_weights: HashMap<String, u32>,
}

impl LoadBalancer {
    fn new() -> LoadBalancer {
        LoadBalancer {
            servers: Vec::new(),
            server_weights: HashMap::new(),
        }
    }

    fn add_server(&mut self, server: String, weight: u32) {
        self.servers.push(server.clone());
        self.server_weights.insert(server, weight);
    }

    fn remove_server(&mut self, server: &str) {
        if let Some(index) = self.servers.iter().position(|s| s == server) {
            self.servers.remove(index);
            self.server_weights.remove(server);
        }
    }

    fn get_server(&self) -> Option<&str> {
        let total_weight: u32 = self.server_weights.values().sum();
        let random_num: u32 = rand::thread_rng().gen_range(0..total_weight + 1);
        let mut current_weight: u32 = 0;

        for (server, weight) in &self.server_weights {
            current_weight += weight;
            if random_num <= current_weight {
                return Some(server.as_str());
            }
        }

        None
    }
}

fn main() {
    let mut lb = LoadBalancer::new();
    lb.add_server(String::from("Server A"), 3);
    lb.add_server(String::from("Server B"), 2);
    lb.add_server(String::from("Server C"), 1);

    for _ in 0..10 {
        if let Some(server) = lb.get_server() {
            println!("Request sent to: {}", server);
        }
    }
}
