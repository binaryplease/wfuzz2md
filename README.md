# wfuzz2md

Generate a Markdown table from [wfuzz](https://github.com/xmendez/wfuzz) JSON output.

## Usage

Use `-o json` to output as JSON and pipe the output of wfuzz directly into `wfuzz2md`

```bash
wfuzz -w <wordlist> -o json <url> | wfuzz2md
```

## Example

### Input

```json
[
    {
        "chars": 277,
        "code": 403,
        "payload": ".htaccess",
        "lines": 9,
        "location": "",
        "method": "GET",
        "post_data": [],
        "server": "Apache/2.4.29 (Ubuntu)",
        "url": "http://10.10.10.171/.htaccess/",
        "words": 28
    },
    {
        "chars": 277,
        "code": 403,
        "payload": ".htpasswd",
        "lines": 9,
        "location": "",
        "method": "GET",
        "post_data": [],
        "server": "Apache/2.4.29 (Ubuntu)",
        "url": "http://10.10.10.171/.htpasswd/",
        "words": 28
    },
    {
        "chars": 14461,
        "code": 200,
        "payload": "artwork",
        "lines": 371,
        "location": "",
        "method": "GET",
        "post_data": [],
        "server": "Apache/2.4.29 (Ubuntu)",
        "url": "http://10.10.10.171/artwork/",
        "words": 918
    },
    {
        "chars": 277,
        "code": 403,
        "payload": "server-status",
        "lines": 9,
        "location": "",
        "method": "GET",
        "post_data": [],
        "server": "Apache/2.4.29 (Ubuntu)",
        "url": "http://10.10.10.171/server-status/",
        "words": 28
    }
]
```

### Output

| Code | Chars | URL |
|---|---|---|
| 403 | 277 	|  http://10.10.10.171/.htaccess/  |
| 403 | 277 	|  http://10.10.10.171/.htpasswd/  |
| 200 | 14461 	|  http://10.10.10.171/artwork/  |
| 403 | 277 	|  http://10.10.10.171/server-status/  |
