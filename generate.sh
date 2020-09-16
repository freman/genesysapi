#!/bin/bash

# Cleanup
rm -fr client models swagger.json

# Download and patch laatest swagger.json
GO111MODULE="on" go run swagger-patch/main.go https://github.com/MyPureCloud/platform-client-sdk-go/raw/master/swagger.json swagger-patch.json

swagger generate client --allow-template-override -f swagger.json -A genesys --skip-validation -T swagger-template

# Could be fixed by adjusting generation, but a bit of copypasta will get it done for now
cp monkeyPatch/models_wrap_up_code_mapping_verify.extra models/wrap_up_code_mapping_verify.go