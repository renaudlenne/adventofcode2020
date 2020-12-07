use crate::utils::read_lines;
use std::collections::HashSet;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      let sum_count = lines.fold::<(HashSet<char>, usize), _>((HashSet::new(), 0), |(mut group_letters, sum_count), line| {
        match line {
          Ok(line) if line.is_empty() => (HashSet::new(), sum_count + group_letters.len()),
          Ok(line) => {
            for c in line.chars() {
              group_letters.insert(c);
            }
            (group_letters, sum_count)
          },
          _ => (group_letters, sum_count)
        }
      }).1;
      println!("{}", sum_count);
    }
  }
}
