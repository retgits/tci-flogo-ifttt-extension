# IFTTT extension for TIBCO Cloud Integration - Web Integator
An extension for TIBCO Cloud Integration - Web Integrator to connect Web Integrator Apps to IFTT and do amazing things.

## Using the extension
In order to work with this extension you'll need to enable the `WebHooks` service on [IFTTT](https://ifttt.com/maker_webhooks). The connector takes two inputs:
* Key: The key you get after enabling the WebHooks service in IFTTT
* Event: _This is the name of the event you specify when creating a trigger in IFTTT_

After you enabled the WebHook service in IFTTT, you can go to the [settings](https://ifttt.com/services/maker_webhooks/settings) to see your `Key`. The `Key` is the last part of the URL that is shown. For example, if the URL is `https://maker.ifttt.com/use/blablabla` the `Key` would be `blablabla`

The activity `Send IFTTT WebHook` will let you trigger IFTTT WebHooks with up to three parameters (this is a limit set by IFTTT)

## License
The MIT License (MIT)

Copyright (c) 2017 retgits

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.