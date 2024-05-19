document.querySelector("#request-form").addEventListener("submit", (event) => {
  event.preventDefault();
  const data = [...new FormData(event.target).entries()];
  console.log(data);
});
