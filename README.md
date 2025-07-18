# goimail

library with icloud mail client _(iCloud+ required)_ that allows to use icloud-based mail sending.

installation:

```bash
go get github.com/tomek7667/goimail
```

## Usage

In order to have this working, you need the following variables: `icloudEmail`, `senderEmail`, `appSpecificPassword`.

```go
package main

import "github.com/tomek7667/goimail"

func main() {
	// in order to know how to obtain the following values, see instructions section below.
	icloudEmail := "...@icloud.com"
	senderEmail := "...@custom-domain.com"
	appSpecificPassword := "xxxx-xxxx-xxxx-xxxx"

	// initializing a client
	client, err := icloud.New(icloudEmail, senderEmail, appSpecificPassword)
	if err != nil {
		panic(err)
	}

	// sending an actual e-mail
	err = c.SendMail("This is test of the client", "<h1>hiii</h1>", nil, "receiver@example.com")
	if err != nil {
		panic(err)
	}
}
```

**Security notice**

I strongly encourage to pass those variables either from a secrets service _(e.g.: AWS Secrets Manager)_, or at least pass them via environment variables. Do not ever hard-code these values - they are considered extremely sensitive.

## Instructions on how to obtain each of initialization values

### `icloudEmail`

This is the standard email you use to log in to icloud. **It is not the email of the mail sender**

It is under your first and last name in [apple account website](https://account.apple.com/account/manage/).

### `senderEmail`

This is the actual email that will be as mail sender.

To set it up:

1. Go to [icloud plus website](https://www.icloud.com/icloudplus/).
2. Click on `Custom Email Domain`. If you don't already have a domain, add one _(See the buttons at the bottom of the dialog)_.
3. Click on a domain you want your sender email to be from.
4. Click `+` to add an email address, type your desired email address and click `Add email address`.

The added email address is the `senderEmail` value.

### `appSpecificPassword`

This is used to authenticate against smtp server.

To set it up:

1. Go to [apple account website, _Sign-In and Security_](https://account.apple.com/account/manage/section/security).
2. Click on `App-Specific Passwords`, and hit `+`.
3. Give it a name _(not relevant)_ and click `Create`. Confirm with your account password.
4. Do **NOT** hit `Done` before you copy your app-specific password into some secure place. _(I recommend using bitwarden as it's a free password manager)_
