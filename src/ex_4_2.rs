use crate::utils::read_lines;
use std::collections::{HashSet, HashMap};
use regex::Regex;

fn is_passport_valid(fields: HashMap<String, String>) -> bool {
  let needed_fields: HashSet<_> = ["byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"].iter().map(|&s| String::from(s)).collect();

  let valid_fields: Vec<String> = fields.iter().fold(Vec::new(), |mut valid_fields, (field_name, field_value)| {
    let is_match = match field_name.as_str() {
      "byr" => {
        match field_value.parse::<i32>() {
          Ok(byr) if (byr >= 1920) && (byr <= 2002) => true,
          _ => false
        }
      },
      "iyr" => {
        match field_value.parse::<i32>() {
          Ok(iyr) if (iyr >= 2010) && (iyr <= 2020) => true,
          _ => false
        }
      },
      "eyr" => {
        match field_value.parse::<i32>() {
          Ok(eyr) if (eyr >= 2020) && (eyr <= 2030) => true,
          _ => false
        }
      },
      "hgt" => {
        let metric_hgt_regex = Regex::new(r"^(\d{3})cm$").unwrap();
        match metric_hgt_regex.captures(field_value) {
          Some(cap) => {
            let size = cap[1].parse::<usize>().unwrap();
            size >= 150 && size <= 193
          }
          _ => {
            let imperial_hgt_regex = Regex::new(r"^(\d{2})in$").unwrap();
            match imperial_hgt_regex.captures(field_value) {
              Some(cap) => {
                let size = cap[1].parse::<usize>().unwrap();
                size >= 59 && size <= 76
              }
              _ => false
            }
          }
        }
      },
      "hcl" => {
        let hcl_regex = Regex::new(r"^#([0-9a-f]{6})$").unwrap();
        hcl_regex.is_match(field_value)
      },
      "ecl" => {
        let valid_ecls: HashSet<_> = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"].iter().map(|&s| String::from(s)).collect();
        valid_ecls.contains(field_value)
      },
      "pid" => {
        let pid_regex = Regex::new(r"^(\d{9})$").unwrap();
        pid_regex.is_match(field_value)
      },
      _ => true
    };
    if is_match {
      valid_fields.push(field_name.clone());
      valid_fields
    } else {
      valid_fields
    }
  });

  needed_fields.iter().all(|item| valid_fields.contains(&item))
}

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      let (last_passport_fields, current_valids) = lines.fold::<(HashMap<String, String>, i32), _>((HashMap::new(), 0), |(mut current_passport_fields, nb_valids), res| {
        let cloned_fields = current_passport_fields.clone();
        match res {
          Ok(line) if (line.is_empty()) && (is_passport_valid(cloned_fields)) => (HashMap::new(), nb_valids + 1),
          Ok(line) if line.is_empty() => (HashMap::new(), nb_valids),
          Ok(line) => {
            for field in line.split(' ') {
              let mut split_iter = field.split(':');
              current_passport_fields.insert(split_iter.next().unwrap().to_string(), split_iter.next().unwrap().to_string());
            }
            (current_passport_fields, nb_valids)
          },
          _ => (HashMap::new(), nb_valids)
        }
      });
      let nb_valids = match is_passport_valid(last_passport_fields) {
        true => current_valids +1,
        false => current_valids,
      };

      println!("{}", nb_valids);
    }
  }
}
