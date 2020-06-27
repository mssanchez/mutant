#!/bin/bash

gcloud app create --project=positive-apex-280419

gcloud components install app-engine-go

gcloud app deploy

gcloud app logs tail -s default