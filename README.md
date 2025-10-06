# Open-Verification

The goal of this project is to create a way to verify Government issued IDs without exposing the tenants to sensitive or regulated information

The way it works is this server will handle file upload, proccess the id then answer prompts provided in the request keeping all request data in memory and never saving to disk

so for example if a service wants to do age verification they can upload the image including the flag is_over_18 and the server would only return a boolean response
