


# zkdump
Dump Zookeeper data to JSON.

Add `--help` argument for more information on usage:

```
$ zkdump --help
usage: zkdump --server=SERVER [<flags>] [<path>]

A command-line utility to dump Zookeeper data.

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
      --version            Show application version.
  -s, --server=SERVER ...  Host name and port to connect to (host:port)
  -v, --verbose            Print verbose.
  -u, --user=USER          Username to use for digest authentication.
  -p, --password=PASSWORD  Password to use for digest authentication (will read from TTY if not given).
  -r, --recursive          Get nodes recursively.

Args:
  [<path>]  Root path (default: "/").

```

Combine with utilities such as [jq](https://stedolan.github.io/jq/) to convert transform and filter output.

## Siebel example

This utility was primarily created to read output from Siebel configuration data, which is stored in Zookeeper starting with Siebel IP2017.

When no path is supplied, all data is retrieved by default. For example, to retrieve all data from the Siebel Gateway with hostname `cgw`, registry port 2320, username "SADMIN" and password "oracle", use the following command:

```
$ zkdump --server cgw:2320 --user SADMIN --password oracle 


```

To retrieve a certain node, such as the Gateway security profile, use a command such as:

```
$ zkdump --server cgw:2320 --user SADMIN --password oracle /Config/Profiles/Security/Gateway

{
  "Name": "Gateway",
  "Path": "/Config/Profiles/Security/Gateway",
  "Data": "ewogICAgIlByb2ZpbGUiOiB7CiAgICAgICAgIlByb2ZpbGVOYW1lIjogIkdhdGV3YXkiLAogICAgICAgICJBY2Nlc3NQZXJtaXNzaW9uIjogIlJlYWRPbmx5IgogICAgfSwKICAgICJTZWN1cml0eUNvbmZpZ1BhcmFtcyI6IHsKICAgICAgICAiRGF0YVNvdXJjZXMiOiBbewogICAgICAgICAgICAiTmFtZSI6ICJQREIxIiwKICAgICAgICAgICAgIlR5cGUiOiAiREIiLAogICAgICAgICAgICAiSG9zdCI6ICIxOTIuMTY4LjU2LjE2MCIsCiAgICAgICAgICAgICJQb3J0IjogMTUyMSwKICAgICAgICAgICAgIlNxbFN0eWxlIjogIk9yYWNsZSIsCiAgICAgICAgICAgICJFbmRwb2ludCI6ICJQREIxIiwKICAgICAgICAgICAgIlRhYmxlT3duZXIiOiAiU0lFQkVMIiwKICAgICAgICAgICAgIkhhc2hVc2VyUHdkIjogZmFsc2UsCiAgICAgICAgICAgICJIYXNoQWxnb3JpdGhtIjogIlNIQTEiLAogICAgICAgICAgICAiQ1JDIjogIiIsCiAgICAgICAgICAgICJfcHJldlR5cGUiOiAiIgogICAgICAgIH1dLAogICAgICAgICJTZWNBZHB0TmFtZSI6ICJEQlNlY0FkcHQiLAogICAgICAgICJTZWNBZHB0TW9kZSI6ICJEQiIsCiAgICAgICAgIk5TQWRtaW5Sb2xlIjogWyJTaWViZWwgQWRtaW5pc3RyYXRvciJdLAogICAgICAgICJUZXN0VXNlck5hbWUiOiAiIiwKICAgICAgICAiVGVzdFVzZXJQd2QiOiAiIiwKICAgICAgICAiREJTZWN1cml0eUFkYXB0ZXJEYXRhU291cmNlIjogIlBEQjEiLAogICAgICAgICJEQlNlY3VyaXR5QWRhcHRlclByb3BhZ2F0ZUNoYW5nZSI6IGZhbHNlLAogICAgICAgICJDZXJ0aWZpY2F0ZVJvb3QiOiAiXC9zaWViZWxcL2NvbmZpZ1wvdHJ1c3RzdG9yZS5qa3MiCiAgICB9Cn0="
}
```

Pipe the output into `jq`, to retrieve only the Data element:

```
$ zkdump --server localhost:2320 --user SADMIN --password oracle /Config/Profiles/Security/Gateway | jq -r '.Data'

ewogICAgIlByb2ZpbGUiOiB7CiAgICAgICAgIlByb2ZpbGVOYW1lIjogIkdhdGV3YXkiLAogICAgICAgICJBY2Nlc3NQZXJtaXNzaW9uIjogIlJlYWRPbmx5IgogICAgfSwKICAgICJTZWN1cml0eUNvbmZpZ1BhcmFtcyI6IHsKICAgICAgICAiRGF0YVNvdXJjZXMiOiBbewogICAgICAgICAgICAiTmFtZSI6ICJQREIxIiwKICAgICAgICAgICAgIlR5cGUiOiAiREIiLAogICAgICAgICAgICAiSG9zdCI6ICIxOTIuMTY4LjU2LjE2MCIsCiAgICAgICAgICAgICJQb3J0IjogMTUyMSwKICAgICAgICAgICAgIlNxbFN0eWxlIjogIk9yYWNsZSIsCiAgICAgICAgICAgICJFbmRwb2ludCI6ICJQREIxIiwKICAgICAgICAgICAgIlRhYmxlT3duZXIiOiAiU0lFQkVMIiwKICAgICAgICAgICAgIkhhc2hVc2VyUHdkIjogZmFsc2UsCiAgICAgICAgICAgICJIYXNoQWxnb3JpdGhtIjogIlNIQTEiLAogICAgICAgICAgICAiQ1JDIjogIiIsCiAgICAgICAgICAgICJfcHJldlR5cGUiOiAiIgogICAgICAgIH1dLAogICAgICAgICJTZWNBZHB0TmFtZSI6ICJEQlNlY0FkcHQiLAogICAgICAgICJTZWNBZHB0TW9kZSI6ICJEQiIsCiAgICAgICAgIk5TQWRtaW5Sb2xlIjogWyJTaWViZWwgQWRtaW5pc3RyYXRvciJdLAogICAgICAgICJUZXN0VXNlck5hbWUiOiAiIiwKICAgICAgICAiVGVzdFVzZXJQd2QiOiAiIiwKICAgICAgICAiREJTZWN1cml0eUFkYXB0ZXJEYXRhU291cmNlIjogIlBEQjEiLAogICAgICAgICJEQlNlY3VyaXR5QWRhcHRlclByb3BhZ2F0ZUNoYW5nZSI6IGZhbHNlLAogICAgICAgICJDZXJ0aWZpY2F0ZVJvb3QiOiAiXC9zaWViZWxcL2NvbmZpZ1wvdHJ1c3RzdG9yZS5qa3MiCiAgICB9Cn0=
```

Note that in this case, the node's data is Base64 encoded. It can easily be decoded piping the output into the `base64` utility:

```
$ zkdump --server localhost:2320 --user SADMIN --password oracle /Config/Profiles/Security/Gateway | jq -r '.Data' | base64 -d

{
    "Profile": {
        "ProfileName": "Gateway",
        "AccessPermission": "ReadOnly"
    },
    "SecurityConfigParams": {
        "DataSources": [{
            "Name": "PDB1",
            "Type": "DB",
            "Host": "192.168.56.160",
            "Port": 1521,
            "SqlStyle": "Oracle",
            "Endpoint": "PDB1",
            "TableOwner": "SIEBEL",
            "HashUserPwd": false,
            "HashAlgorithm": "SHA1",
            "CRC": "",
            "_prevType": ""
        }],
        "SecAdptName": "DBSecAdpt",
        "SecAdptMode": "DB",
        "NSAdminRole": ["Siebel Administrator"],
        "TestUserName": "",
        "TestUserPwd": "",
        "DBSecurityAdapterDataSource": "PDB1",
        "DBSecurityAdapterPropagateChange": false,
        "CertificateRoot": "\/siebel\/config\/truststore.jks"
    }
}
```
