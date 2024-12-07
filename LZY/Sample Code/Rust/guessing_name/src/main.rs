use std::io;
use rand::Rng;
use std::cmp::Ordering;

fn main() {
    let secret_number = rand::thread_rng().gen_range(1..=100);
    loop {
        println!("Please enter a guess value:");
        let mut guess = String::new();

        io::stdin().read_line(&mut guess).expect("You have entered the wrong value");
        println!("The value you entered is {guess}");

        //let guess: u32 = guess.trim().parse().expect("You have entered the wrong type of value");
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue,
        };

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("The guess value is less"),
            Ordering::Greater => println!("The guess value is greater"),
            Ordering::Equal => {
                println!("You guess the right value, Congratulations!!!");
                break;
            }
        }
    }
}
