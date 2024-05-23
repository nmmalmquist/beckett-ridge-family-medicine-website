const useState = (init) => {
  let local = init;
  return [() => local, (value) => (local = value)];
};

const addErrorToField = (field, errMessage) => {
  if (document.querySelector(`.${field.id}-error`)) {
    return;
  }
  const el = document.createElement("span");
  el.classList.add("text-danger", "text-sm", `${field.id}-error`);
  el.textContent = errMessage;
  field.closest(".input-wrapper").after(el);
};
const removeErrorFromField = (field) => {
  document.querySelector(`.${field.id}-error`)?.remove();
};

const errMessageExistsOnField = (field) => {
  return !!document.querySelector(`.${field.id}-error`);
};
