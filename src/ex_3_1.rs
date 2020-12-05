use crate::ex_3::slide;

pub(crate) fn run(args: &[String]) {
  if let Some(file_name) = args.first() {
    let nb_trees = slide(file_name, 3, 1);
    println!("{}", nb_trees);
  }
}
