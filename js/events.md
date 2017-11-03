# Events

The browser runs a perpectual event loop which listens and reacts to events.
This includes loading resources (html, js, css), dom manipulations to draw
resources onto screen, observing user interactions (clicks, mouse movements etc.).

The atomic unit of the event system is an `event` which is has 3 properties:
- a name (string)
- a set of properties (object)
- an emitter (object)

Browsers implement this system using `EventProtoType` as the base class for
event properties, and `addEventListener` as a way of interacting with the
emitter.

#### Example

```
const buttonDOMElement = document.querySelect('#button')
buttonDOMElement.addEventListener('click', (event) => {
  console.log(event)
})
```

In this example, `buttonDOMElement` is the emitter. `'click'` is the name, and `event`
contains the set of properties.
