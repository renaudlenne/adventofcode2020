use crate::utils::read_lines;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    if let Ok(lines) = read_lines(file_name) {
      let (_, nb_trees) = lines.skip(1).fold((3, 0), |(idx, nb_trees), res| {
        match res {
          Ok(line) => {
            let line_str: Vec<char> = line.as_str().chars().collect();
            match line_str.get(idx % line.len()) {
              Some('#') => (idx + 3, nb_trees + 1),
              _ => (idx + 3, nb_trees)
            }
          },
          _ => (idx, nb_trees)
        }
      });
      println!("{}", nb_trees);
    }
  }
}
