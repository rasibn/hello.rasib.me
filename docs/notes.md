# Using the CLI

Install the Tailwind CSS v4 alpha and the separate CLI package:

```bash
npm install tailwindcss@next @tailwindcss/cli@next
```

Next, import Tailwind in your main CSS file:

```css
app.css
@import "tailwindcss";
```

Finally, compile your CSS using the CLI tool:

```bash
$ npx @tailwindcss/cli@next -i app.css -o dist/app.css
```

---

The /lib/systemd/system/hello.rasib.me-go.service

```
[Unit]
Description=hello.rasib.me

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/ubuntu/hello.rasib.me/build/Release
WorkingDirectory=/home/ubuntu/hello.rasib.me

[Install]
WantedBy=multi-user.target
```

---

Check logs on the server using

```
sudo journalctl -u hello.rasib.me-go.service -f
```
