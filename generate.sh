#!/bin/bash

rm -fr client models

swagger generate client -f https://github.com/MyPureCloud/platform-client-sdk-go/raw/master/swagger.json -A genesys --skip-validation -T swagger-template

# Could be fixed by adjusting generation, but a bit of copypasta will get it done for now
cp monkeyPatch/models_wrap_up_code_mapping_verify.extra models/wrap_up_code_mapping_verify.go