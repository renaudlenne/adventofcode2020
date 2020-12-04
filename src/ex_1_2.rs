use crate::utils::read_lines;
use itertools::Itertools;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      for combination in lines.map(|line| {
        match line {
          Ok(value) => value.parse::<i32>().unwrap(),
          _ => 0
        }
      }).combinations(3) {
        match combination.as_slice() {
          [val1, val2, val3] if val1 + val2 + val3 == 2020 => println!("{}", val1*val2*val3),
          _ => ()// do nothing
        }
      };
    }
  }
}
