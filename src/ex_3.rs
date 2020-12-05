use crate::utils::read_lines;

pub(crate) fn slide(file_name: &String, right: usize, down: usize) -> i64 {
  match read_lines(file_name) {
    Ok(lines) => {
      let (_, nb_trees, _) = lines.skip(down).fold((right, 0, 0), |(idx, nb_trees, skip), res| {
        if skip > 0 { return (idx, nb_trees, skip-1)}
        match res {
          Ok(line) => {
            let line_str: Vec<char> = line.as_str().chars().collect();
            let next_nb_trees = match line_str.get(idx % line.len()) {
              Some('#') => nb_trees+1,
              _ => nb_trees
            };
            (idx + right, next_nb_trees, down-1)
          },
          _ => (idx, nb_trees, skip)
        }
      });
      nb_trees
    },
    _ => 0
  }
}
