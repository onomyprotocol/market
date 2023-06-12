// for locally cleaning up log files and other temproraries

use onomy_test_lib::super_orchestrator::{remove_files_in_dir, stacked_errors::Result, std_init};

#[tokio::main]
async fn main() -> Result<()> {
    std_init()?;

    remove_files_in_dir("./tests/dockerfiles/dockerfile_resources", &[
        "marketd",
        "market_standaloned",
    ])
    .await?;
    remove_files_in_dir("./tests/logs", &["log", "json"]).await?;
    remove_files_in_dir("./tests/resources/keyring-test/", &["address", "info"]).await?;

    Ok(())
}