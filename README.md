# go-toolkit

A collection of small but useful packages that I use in my go applications. I primarily created this collection as a way to learn go as I learn best by solving problems. As a bonus, I've grown to also learn a little bit about how things are implemented at a lower level as I am attempting to implement these packages using just the standard library.

## Packages

### Mail

The mail package allows you to create and send emails using a simple API.

#### Installation

```bash
    go get github.com/mwprogrammer/go-toolkit/mail
```

#### Usage

1. Get and configure the email sender with your email server's credentials.

```go
    sender := mail.GetSender()
    sender.Configure("host", "username", "password", "port")
```

2. Create an email message.

```go
    message := mail.CreateMessage()

    // Add receipients
    var receipients = [...]string{"example1@gmail.com", "example1@outlook.com"}
    var cc = "example2@gmail.com"
    var bcc = "example3@outlook.com"

    for _, receipient := range receipients {
        message.AddTo(receipient)
    }

    message.AddCC(cc)
    message.AddBCC(bcc)

    // Add a subject
    message.AddSubject("TEST SIMPLE MESSAGE WITH ATTACHMENTS")

    // Add a body. Note you can also add html as well.
    message.AddBody("Test sending message with attachments.")

    // Add an attachment
    err := message.AttachFile("../path/to/file")

    if err != nil {
        fmt.Println("could not attach file.")
    }

```

3. Send email.

```go
    err = sender.Send(message)
```

## Roadmap

-   [x] Mail - Had fun learning a bit more about the SMTP protocol.
-   [ ] Http - Implementing and dogfooding this largely informs many of the ideas I have learned managing the balance between abstraction and flexibility.
-   [ ] Logging - I honestly thought I'd provide a wrapper for slog and call it a day. Dogfooding this was a bit painful but educational.
-   [ ] Databases? - I'm trying to implement my own 'mini-orm' in another application and depending on the results of that experiment ... maybe I will add this package. I dunno.
-   [ ] Jobs? - I recently wrote a windows service successfully in Go. The design for what to externalize into a separate package felt finicky but the experience made me wonder if perhaps I could design a more general package for creating background jobs and the like. Would be my most ambitious package to date.

## Contributions

This is a project I work on primarily for learning purposes. Sometimes I rewrite entire lines of code when I see something doesn't work in a personal application that I'm dogfooding the packages on or change significant logic that will break changes. For instance; go-toolkit used to be go-utilities which was a large package with different modules for mail etc. Now it is akin to a mono repo of sorts for different much smaller packages. Use the packages at your own risk. On the other hand; I am accepting code reviews and suggestions for what to work on next. If interested, send an email to **malawianprogrammer@gmail.com**. 

## License

Distributed under the MIT license. See [LICENSE.md](LICENSE.md) for more details.
