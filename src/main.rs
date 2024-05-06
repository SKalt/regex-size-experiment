use regex::Regex;
use std::mem::size_of;

fn main() {
    // let x = Regex::new(r"yeet").unwrap();
    println!("Size of Regex: {}", size_of::<Regex>());
}
