# Cloud `cat`

## Get ready

You must setup a service account like so:

```bash
gcloud iam service-accounts create [NAME]
gcloud projects add-iam-policy-binding [PROJECT_ID] --member "serviceAccount:[NAME]@[PROJECT_ID].iam.gserviceaccount.com" --role "roles/owner"
gcloud iam service-accounts keys create [FILE_NAME].json --iam-account [NAME]@[PROJECT_ID].iam.gserviceaccount.com
export GOOGLE_APPLICATION_CREDENTIALS="[PATH]"
```

## Go!

`ccat PROJECT_ID`