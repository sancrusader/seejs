# JavaScript URL & API Key Analyzer

This Go-based tool analyzes JavaScript files from a list of URLs to identify potentially sensitive data such as API keys, secrets, and URLs in JavaScript files. It’s designed to help security researchers, developers, and penetration testers audit JavaScript files for exposed secrets and endpoints.

# Features

API Key Detection: Identifies Google Maps, AWS, Azure, Stripe, Twilio, GitHub, Firebase, Slack, Dropbox, Facebook, PayPal, Mailgun, SendGrid, Mailchimp, Telegram, and other popular API keys.

Advanced Regex Matching: Uses complex regex patterns to capture various API key formats and secret tokens.

URL Extraction: Detects endpoints and paths within JavaScript files, including absolute and relative URLs.

Automated Output: Saves findings (URLs, API keys, and secrets) to a specified output file.


# Usage

Prerequisites

Go 1.16 or newer.


# Installation

1. Clone the repository:

```
git clone https://github.com/sancrusader/seejs.git
cd seejs
```


2. Build the project (optional):

```
go build main.go
```



# Running the Analyzer

1. Prepare a list of URLs in a file (e.g., js.txt), with each URL on a new line. Example:

```
https://example.com/script1.js
https://example.com/script2.js
```


2. Run the tool with the following command:

```
go run main.go -l path/to/js.txt
```


3. The results will be saved in output.txt with only the detected API keys, secrets, and URLs.


Example Output

The output file will contain only the detected secrets, without additional text. Sample:

```
AIzaSyD... (Google Maps API Key)
AKIAIOSFODNN... (AWS Access Key)
sk_test_BQokik... (Stripe API Key)
https://example.com/api/endpoint
```

Advanced Configuration

The regex patterns for key detection can be easily modified within the script to capture additional or custom key formats.

# Use Cases

Bug Bounty Hunting: Quickly scan JavaScript files for exposed keys, secrets, and sensitive endpoints.

Security Audits: Ideal for security researchers auditing applications for misconfigured or exposed sensitive information.

Application Testing: Aids developers and testers in verifying that secrets are not accidentally pushed to production.


Contributing

Contributions are welcome! Feel free to submit a pull request or create an issue for suggestions and improvements.


---

Let me know if you'd like further customizations or additions!

