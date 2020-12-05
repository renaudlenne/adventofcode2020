use crate::utils::read_lines;
use regex::Regex;


fn fit_password_policy(line: &String) -> bool {
  let password_regex: Regex = Regex::new(r"(\d{1,})-(\d{1,}) (\w{1}): (\w{1,})").unwrap();
  match password_regex.captures(line.as_str()) {
    Some(cap) => {
      let first_pos = cap[1].parse::<usize>().unwrap();
      let second_post = cap[2].parse::<usize>().unwrap();
      let char_to_check = &cap[3].parse::<char>().unwrap();
      let password:Vec<char> = cap[4].chars().collect();
      match password.get(first_pos-1) {
        None => false,
        Some(first_char) =>
          match password.get(second_post-1) {
            None => first_char == char_to_check,
            Some(second_char) if second_char == char_to_check => first_char != char_to_check,
            Some(_) => first_char == char_to_check
          }
      }
    },
    _ => false
  }
}

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      println!("{}", lines.fold::<i32, _>(0, |acc, res|
        match res {
          Ok(line) if fit_password_policy(&line) => acc+1,
          _ => acc
        }
      ))
    }
  }
}
