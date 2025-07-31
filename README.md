# SETUP:
- run
```sh
    make dev
```
- go to http://localhost:8000/
- read the description in the page
- click an item of the selectbox
- expects the number to appear on the display
- nothing appears on the display :(

# ISSUE:
## Describe the issue

So I have found this really weird(?) use case of the SelectBox that is kinda of breaking stuff.
The thing is, I threw a bunch of htmx attributes in each of the selectbox.Items inside the selectbox.Content, that ends up not working at all because of how the initSelect removes the eventListeners.
That is, the initSelect clones the selectbox.Content nodes and replaces it completely on the DOM, which does indeed remove the eventListeners that are about to be setup and avoid duplicates, but it also ends up removing eventListeners that have nothing to do with templUI, like the ones added by htmx (e.g. hx-trigger). 

---

## Reproduction steps

- Setup a templUI app that contains a SelectBox;
- Add some custom eventListeners in the selectbox.Item inside the selectbox.Content;
- Run the app and interact with the items trying to trigger the events;
- The events are not triggered.

---

## Environment

templUI version: 0.85.0
Go version: 1.24.3 linux/amd64
Templ version: 0.3.924
Tailwind version: 4.1.8

---

## 🔍 Please include a minimal sample repo if possible

https://github.com/CarraraVitor/templui_swap_issue
(the name is kinda wrong because it is the same repo that I used for another issue)
