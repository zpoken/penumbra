## How to redo the parameter setup

Note: this is a setup process only for testnet purposes since it is done by
a single party

* `cargo run` in this folder.

The verifying and proving keys for each circuit will be created in a serialized
form in the `proof-params/src/gen` folder. The proving keys should be tracked using git lfs. The verifying keys are stored directly in git since they are small.

Each time you regenerate the parameters, increment the version of the
`penumbra-proof-params` crate.

To add a _new_ circuit to the parameter setup, you should modify `src/main.rs`
before running `cargo run`. Then edit `penumbra-proof-params` to reference
the new paramters.
