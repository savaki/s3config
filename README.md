# s3config
Manage application config files from S3

### Decode JSON

```
config := Config{}
err := s3config.NewDecoder("bucket-name", "key-name").Decode(&config)
```

### Decode XML

```
config := Config{}
err := s3config.NewDecoder("bucket-name", "key-name", s3config.XML).Decode(&config)
```

### Decode HCL

```
config := Config{}
err := s3config.NewDecoder("bucket-name", "key-name", s3config.HCL).Decode(&config)
```

### Encode JSON

```
config := Config{}
err := s3config.NewEncoder("bucket-name", "key-name", s3config.HCL).Encode(&config)
```

