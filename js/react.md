# React

## Terminology

#### Element

An element is an [object](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object) that describes a dom element. It is also a recursive data structure because it's a tree node.

```js
// Element
{
  type: 'div',      // html element
  props: {          // attributes
    children: ...,  // Element
    ...,
  },
}
```

#### Component

A react component is a function that returns an element (functional component) or returns an instance
that returns an element (classbased component).

## React Implementation Details

React is split into a couple different distinct abstractions, core, renderers,
and reconcilers.

#### Core

APIs that define components, available under the global namespace `React`.

#### Renderers

Takes react element tree as input and outputs platform calls to render the given
tree.

- React DOM
- React Native
- React Test

#### Reconcilers

React element tree manipulation, work scheduler.

Renderer will pass a component into the reconciler, the reconciler will then
check the type of component it is and perform the appropriate logic to extract
the element from the component and its subtree.

#### Event System

?
