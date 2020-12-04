use std::env;

mod utils;
mod ex_1_1;
mod ex_1_2;

fn main() {
  let args: Vec<String> = env::args().skip(1).collect();
  if let Some((exercise_nb, exercise_args)) = args.split_first() {
    match exercise_nb.as_str() {
      "1-1" => ex_1_1::run(exercise_args),
      "1-2" => ex_1_2::run(exercise_args),
      _ => println!("Unknown exercise {}", exercise_nb)
    }
  }
}
