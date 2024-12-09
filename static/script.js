document.addEventListener('htmx:beforeRequest', (event) => {
      if (document.startViewTransition) {
        document.startViewTransition(() => {});
      }
});
