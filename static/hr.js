setTimeout(() => {
  document.body.addEventListener("htmx:sendError", function (evt) {
    console.log(`SE ~ evt:`, evt);
  });
  document.body.addEventListener("htmx:afterRequest", function (evt) {
    const target = evt.detail.target;
    if (target.id === "hr") {
      const lastId = target.value;
      const res = evt.detail.xhr.response;
      if (!lastId) evt.detail.target.value = res;
      else if (res !== lastId) {
        target.value = evt.detail.xhr.response;
        console.log("detected change, reloading");
        setTimeout(() => {
          window.location.reload(true);
        }, 500);
      }
    }
  });
}, 150);
