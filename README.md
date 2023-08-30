# Goal
have a go program that will send data to an sqs queue that will invoke a lambda function to convert the message to be encrypted and stored in s3, which will send an sns notification (text) that my logs have been encrypted.

# Description

# Local Lambda development

# setup github action
This project has a github action that will push an image to ecr. For this action to do such a thing, it needs IAM permissions, so this section will show how to setup the IAM role needed for the github action to work.

The first thing you will need to do is go into AWS, and setup a role with these permissions:
```

```