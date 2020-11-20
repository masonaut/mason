// build:mason v0.1

name: "sample web"
description: "description of web workflow"

workflow "web build": {
    description: "build web files"
    actions: [
        "install-files"
    ]
}

action "install-files": {
    commands: [
        "ls -la ."
        "echo 'install files'"
    ]
}

action "package-files": {
    commands: [
        "echo 'packaging files'"
    ]
}
