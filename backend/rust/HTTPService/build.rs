fn main() -> Result<(), Box<dyn std::error::Error>> {
    println!("HERE I PASS");
    tonic_build::compile_protos("../../proto/http_action_service.proto")?;
    tonic_build::compile_protos("../../proto/pmodel.proto")?;
    tonic_build::compile_protos("../../proto/core.proto")?;
    Ok(())
}
