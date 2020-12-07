use crate::utils::read_lines;
use std::collections::HashSet;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      let sum_count = lines.fold::<(HashSet<char>, usize, bool), _>((HashSet::new(), 0, true), |(mut group_letters, sum_count, first_in_group), line| {
        match line {
          Ok(line) if line.is_empty() => (HashSet::new(), sum_count + group_letters.len(), true),
          Ok(line) => {
            if first_in_group {
              for c in line.chars() {
                group_letters.insert(c);
              }
            } else {
              for gc in group_letters.clone() {
                if !line.chars().any(|c| c == gc) {
                  group_letters.remove(&gc);
                }
              }
            }
            (group_letters, sum_count, false)
          },
          _ => (group_letters, sum_count, first_in_group)
        }
      }).1;
      println!("{}", sum_count);
    }
  }
}
