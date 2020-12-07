use std::env;

mod utils;
mod ex_1_1;
mod ex_1_2;
mod ex_2_1;
mod ex_2_2;
mod ex_3;
mod ex_3_1;
mod ex_3_2;
mod ex_4_1;
mod ex_4_2;
mod ex_5;
mod ex_5_1;
mod ex_5_2;
mod ex_6_1;
mod ex_6_2;

fn main() {
  let args: Vec<String> = env::args().skip(1).collect();
  if let Some((exercise_nb, exercise_args)) = args.split_first() {
    match exercise_nb.as_str() {
      "1-1" => ex_1_1::run(exercise_args),
      "1-2" => ex_1_2::run(exercise_args),
      "2-1" => ex_2_1::run(exercise_args),
      "2-2" => ex_2_2::run(exercise_args),
      "3-1" => ex_3_1::run(exercise_args),
      "3-2" => ex_3_2::run(exercise_args),
      "4-1" => ex_4_1::run(exercise_args),
      "4-2" => ex_4_2::run(exercise_args),
      "5-1" => ex_5_1::run(exercise_args),
      "5-2" => ex_5_2::run(exercise_args),
      "6-1" => ex_6_1::run(exercise_args),
      "6-2" => ex_6_2::run(exercise_args),
      _ => println!("Unknown exercise {}", exercise_nb)
    }
  }
}
