pub(crate) fn seat_number(boarding_pass: &String) -> i32 {
  let row_number = boarding_pass[..7].chars().fold::<(i32, i32), _>((0, 127), |(min, max), current_char| {
    match current_char {
      'F' => (min, min+((max-min)/2)),
      'B' => (min+((max-min)/2)+1, max),
      _ => (min, max),
    }
  }).0;
  let col_number = boarding_pass[7..].chars().fold::<(i32, i32), _>((0, 7), |(min, max), current_char| {
    match current_char {
      'L' => (min, min+((max-min)/2)),
      'R' => (min+((max-min)/2)+1, max),
      _ => (min, max),
    }
  }).0;
  row_number*8+col_number
}
