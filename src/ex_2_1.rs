use crate::utils::read_lines;
use regex::Regex;


fn fit_password_policy(line: &String) -> bool {
  let password_regex: Regex = Regex::new(r"(\d{1,})-(\d{1,}) (\w{1}): (\w{1,})").unwrap();
  match password_regex.captures(line.as_str()) {
    Some(cap) => {
      let min_repeat = cap[1].parse::<usize>().unwrap();
      let max_repeat = cap[2].parse::<usize>().unwrap();
      let char_to_repeat = cap[3].parse::<char>().unwrap();
      let password = &cap[4];
      let repeat_count = password.matches(char_to_repeat).count();
      repeat_count >= min_repeat && repeat_count <= max_repeat
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
