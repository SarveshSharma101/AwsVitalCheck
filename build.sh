#!/bin/bash
go build -o awsvitalcheck
rm -rf vitalStatsCheck
mkdir vitalStatsCheck
mkdir vitalStatsCheck/configs
cp -r configs/awsConfig.yaml ./vitalStatsCheck/configs/config.yaml
cp awsvitalcheck ./vitalStatsCheck
tsh scp -r --proxy=fcftport.northeurope.cloudapp.azure.com -i configs/nvidiaIndetity.key vitalStatsCheck nvidia@azure-yolov5:/home/nvidia/          