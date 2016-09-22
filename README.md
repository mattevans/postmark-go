# postmark-go

[![GoDoc](https://godoc.org/github.com/mattevans/postmark-go?status.svg)](https://godoc.org/github.com/mattevans/postmark-go)

[Go](http://golang.org) bindings for the Postmark API (http://developer.postmarkapp.com/).

This is an unofficial library that is not affiliated with [Postmark](http://postmarkapp.com). Official libraries are available
[here](http://developer.postmarkapp.com/developer-official-libs.html).

Example usage
-------------

```go
auth := &http.Client{
  Transport: &postmark.AuthTransport{Token: "API_TOKEN"},
}
client := postmark.NewClient(auth)

emailReq := &postmark.Email{
  From:       "mail@company.com",
  To:         "jack@sparrow.com",
  TemplateID: 123456,
  TemplateModel: map[string]interface{}{
    "name": "Jack",
    "action_url": "http://click.company.com/welcome",
  },
  Tag:        "onboarding",
  TrackOpens: true,
}
email, response, err := client.Email.Send(emailReq)
if err != nil {
  fmt.Printf("Oh no! \n%v\n%v\n", response, err)
  return err
}
```

Setup
-----------------

You'll need to pass an `API_TOKEN` when initializing the client. This token can be
found under the 'Credentials' tab of your Postmark server. More info [here](http://developer.postmarkapp.com/developer-api-overview.html#authentication).

What's Implemented?
----------------

At the moment only a very small number of API endpoints are implemented. Open an
issue (or PR) if you required additional endpoints!

- [Send Email](http://developer.postmarkapp.com/developer-api-email.html#send-email)

Thanks &amp; Acknowledgements :ok_hand:
----------------

The packages's architecture is adapted from
[go-github](https://github.com/google/go-github), created by [Will
Norris](https://github.com/willnorris). :beers:
