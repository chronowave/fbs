use std::io::Write;
use std::path::Path;

use flatc_rust;

fn main() -> std::io::Result<()> {
    println!("cargo:rerun-if-changed=../lib.fbs");
    println!("cargo:rerun-if-changed=./src/lib.rs");
    flatc_rust::run(flatc_rust::Args {
        lang: "rust",
        inputs: &[Path::new("../lib.fbs")],
        out_dir: Path::new("src"),
        extra: &["--filename-suffix", ""],
        ..Default::default()
    })
    .expect("flatc");

    let mut file = std::fs::OpenOptions::new()
        .write(true)
        .append(true)
        .open("./src/lib.rs")
        .unwrap();

    write!(
        file,
        r#"
pub static TEXT_LUT: [Marker; 3] = [
    Marker::SENTINEL,
    Marker::SOH,
    Marker::FRAG,
];"#
    )
}
