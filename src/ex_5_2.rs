use crate::utils::read_lines;
use crate::ex_5::seat_number;
use itertools::Itertools;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      let all_seats = lines.map(|line| {
        match line {
          Ok(line) => seat_number(&line),
          _ => 0
        }
      }).sorted();
      let mut prev_seat = 0;
      for seat_nb in all_seats {
        if (prev_seat != 0) && (prev_seat != seat_nb-1) {
          println!("{}", seat_nb-1);
          break;
        }
        prev_seat = seat_nb;
      }
    }
  }
}
