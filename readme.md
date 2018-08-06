# IFTTT extension for Flogo apps in TIBCO Cloud Integration
An extension for TIBCO Cloud Integration - Web Integrator to connect Web Integrator Apps to IFTT and do amazing things.

## Using the extension
In order to work with this extension you'll need to enable the `WebHooks` service on [IFTTT](https://ifttt.com/maker_webhooks). The connector takes two inputs:
* Key: The key you get after enabling the WebHooks service in IFTTT
* Event: _This is the name of the event you specify when creating a trigger in IFTTT_

After you enabled the WebHook service in IFTTT, you can go to the [settings](https://ifttt.com/services/maker_webhooks/settings) to see your `Key`. The `Key` is the last part of the URL that is shown. For example, if the URL is `https://maker.ifttt.com/use/blablabla` the `Key` would be `blablabla`

The activity `Send IFTTT WebHook` will let you trigger IFTTT WebHooks with up to three parameters (this is a limit set by IFTTT)
