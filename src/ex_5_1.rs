use crate::utils::read_lines;
use crate::ex_5::seat_number;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      let max_number = lines.map(|line| {
        match line {
          Ok(line) => seat_number(&line),
          _ => 0
        }
      }).max().unwrap_or(0);
      println!("{}", max_number);
    }
  }
}
