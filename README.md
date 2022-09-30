# helpwave api
The official restful application programming interface of helpwave

## Github Actions
### CodeQL Analysis
[Code scanning](https://docs.github.com/en/code-security/code-scanning/automatically-scanning-your-code-for-vulnerabilities-and-errors/about-code-scanning) is a feature of GitHub that can be used to analyze the code. This helps in the search for security vulnerabilities and programming errors.

### Continuous Integration
The go application will be build by GitHub actions,  
afterwards a docker image is build and pushed to the 
[GitHub container registry](https://ghcr.io)

The following docker image tags exist:
* `branch-xxx`, where `xxx` is the branch name
* `edge` for the main branch
* `latest` for the latest tag / release
* `v1`, `v1.0`, `v1.0.0` for the matching tags

Make sure to prefix the tags with `v` if you wan't a image for them, I suggest matching the format: [`v\d+\.\d+\.\d+`](https://regexr.com/6v4qh) (Examples: v1.0.0, v14.15.161, ...)

### Create release
I suggest creating new releases using the GitHub actions workflow. The primary advantage is, that the release is made by `github-actions` and the helpwave artifact is included.

![image](https://user-images.githubusercontent.com/26925347/193222515-98220b50-b320-497d-a012-af4be7cdbe3b.png)

Afterwards you should adjust the title (which is the tag name by default) / description of the release.

![image](https://user-images.githubusercontent.com/26925347/193222838-c2f16900-371d-495f-ab55-9d75b6489cfc.png)
