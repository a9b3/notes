## Web performance

Use the Performance API to measure page load metrics.

Access metrics `window.performance`

Access metrics for resources `window.performance.getEntriesByType('resource')`

#### Prefetch

This is the built-in prefetch api.

*link tag*: supports html filetypes

```
<link rel="prefetch" href="..." />
```

*js*: does not support html filetypes

```
const xhrRequest = new XMLHttpRequest()
xhrRequest.open('GET', url, true)
xhrRequest.send()
```

#### Links

- [MDN docs](https://developer.mozilla.org/en-US/docs/Web/API/Performance)
- [Critical path](https://developers.google.com/web/fundamentals/performance/critical-rendering-path/optimizing-critical-rendering-path)
- [Performance budgets](https://infrequently.org/2017/10/can-you-afford-it-real-world-web-performance-budgets/)
