use crate::ex_3::slide;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    let result = [
      (1, 1),
      (3, 1),
      (5, 1),
      (7, 1),
      (1, 2)
    ].iter().map(|(right, down)| {
      slide(file_name, *right, *down)
    }).fold(1, |acc, cur| {
      acc * cur
    });

    println!("{}", result);
  }
}
