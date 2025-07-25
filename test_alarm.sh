#!/bin/bash

curl -d '{"org_id":"00000000-0000-0000-0000-000000000000","device_id":"1","device_key":"1234567890","content":{"message":"test"}}' 'http://127.0.0.1:4151/pub?topic=core.device.alarm'