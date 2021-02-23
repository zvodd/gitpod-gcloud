This is a hellow world project ready to run locally or deploy to GCP GEA Project Instance.
Using LabStack's Echo framework V4

```
go build && my_example_app -local
```

or

```
gcloud app deploy
```

note, when deploying:  `rm my_example_app` or add the excutable name to `.gcloudignore` to prevent compiled binary being uploaded.
