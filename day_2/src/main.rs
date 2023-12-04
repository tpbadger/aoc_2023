use std::{
    fs::File,
    io::{prelude::*, BufReader},
    path::Path,
};

#[derive(Debug, PartialEq, Eq)]
struct Game {
    id: i32,
    subsets: Vec<Subset>,
    max_vals: [i32; 3] // red, green, blue
}

impl Game {
    fn is_valid(&self, max_red: i32, max_green: i32, max_blue: i32) -> bool {
        for subset in &self.subsets {
            if subset.n_red.is_some() && subset.n_red.unwrap() > max_red {
                return false
            }
            if subset.n_blue.is_some() && subset.n_blue.unwrap() > max_blue {
                return false
            }
            if subset.n_green.is_some() && subset.n_green.unwrap() > max_green {
                return false
            }
        }
        return true;
    }

    fn find_max_vals(&mut self) {
        self.max_vals = [1, 1, 1];
        for subset in &self.subsets {
            if subset.n_red.is_some() && subset.n_red.unwrap() > self.max_vals[0] {
                self.max_vals[0] = subset.n_red.unwrap();
            }   
            if subset.n_green.is_some() && subset.n_green.unwrap() > self.max_vals[1] {
                self.max_vals[1] = subset.n_green.unwrap();
            }  
            if subset.n_blue.is_some() && subset.n_blue.unwrap() > self.max_vals[2] {
                self.max_vals[2] = subset.n_blue.unwrap();
            }  
        }
    }


}

#[derive(Debug, PartialEq, Eq)]
struct Subset {
    n_red: Option<i32>,
    n_blue: Option<i32>,
    n_green: Option<i32>
}

fn subsets_from_string(str: &str) -> Vec<Subset> {
    let mut subsets = Vec::new();
    for subset in str.split(";") {
        let mut ss = Subset { n_red: None, n_blue: None, n_green: None };
        for colours in subset.split(",") {
            let num_colours:Vec<&str> = colours.trim().split(" ").collect();
            let colour = num_colours[1];
            match colour {
                "red" => {
                    ss.n_red =  Some(num_colours[0].parse::<i32>().unwrap());
                },
                "blue" => {
                    ss.n_blue = Some(num_colours[0].parse::<i32>().unwrap());
                },
                "green" => {
                    ss.n_green= Some(num_colours[0].parse::<i32>().unwrap());
                },
                _ => {
                    println!("hello");
                }
            }            
        }
        subsets.push(ss);
    }

    return subsets
}

fn read_input(filename: impl AsRef<Path>) -> Vec<String> {
    // read the input into a vector of strings
    let file = File::open(filename).expect("No file with that name");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect()
}


#[test]
fn test_subsets_from_string() {
    let input_string = "1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red";
    let res = subsets_from_string(input_string);
    assert_eq!(res[0], Subset{n_red:Some(3), n_blue:Some(6), n_green:Some(1)})
}

#[test]
fn test_input() {
    let inputs = read_input("./test_input.txt");
    let soln = solve(inputs, 12, 13, 14);
    assert!(soln==8)
}

fn solve(inputs: Vec<String>, max_red: i32, max_green: i32 , max_blue: i32) -> i32 {
    let mut acc = 0;
    let mut g_vec = Vec::new();
    for input in inputs {
        let splt:Vec<&str> = input.split(":").collect();
        let str_split: Vec<&str> = splt[0].split(" ").collect();
        let game_id = str_split[1].parse::<i32>().unwrap();
        let subsets = subsets_from_string(splt[1]);
        let mut g = Game{
            id:game_id,
            subsets:subsets,
            max_vals: [0, 0, 0]
        };
        if g.is_valid(max_red, max_green, max_blue) {
            acc += g.id;
        }
        g.find_max_vals();
        g_vec.push(g);
    }

    let mut acc_2 = 0;
    for game in g_vec {
        let mut prod = 1;
        prod *= game.max_vals[0];
        prod *= game.max_vals[1];
        prod *= game.max_vals[2];
        acc_2 += prod
    }

    println!("Solution to part 2 {acc_2}");

    return acc
}


fn main() {
    let inputs = read_input("./input.txt");
    let soln = solve(inputs, 12, 13, 14);
    println!("Day 2 solution is {soln}");
}