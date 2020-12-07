use crate::utils::read_lines;
use std::collections::HashSet;

fn is_passport_valid(fields: &Vec<String>) -> bool {
  let needed_fields: HashSet<_> = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"].iter().map(|&s| String::from(s)).collect();
  needed_fields.iter().all(|item| fields.contains(item))
}

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      let (last_passport_fields, current_valids) = lines.fold::<(Vec<String>, i32), _>((Vec::new(), 0), |(current_passport_fields, nb_valids), res|
        match res {
          Ok(line) if (line.is_empty()) && (is_passport_valid(&current_passport_fields)) => (Vec::new(), nb_valids+1),
          Ok(line) if line.is_empty() => (Vec::new(), nb_valids),
          Ok(line) => {
            let line_fields: Vec<String> = line.split(' ')
              .map(|field| field.split(':').next().map(|s| s.to_string()))
              .flatten()
              .collect();
            (current_passport_fields.iter().cloned().chain(line_fields).collect(), nb_valids)
          },
          _ => (Vec::new(), nb_valids)
        }
      );
      let nb_valids = match is_passport_valid(&last_passport_fields) {
        true => current_valids +1,
        false => current_valids,
      };

      println!("{}", nb_valids);
    }
  }
}
