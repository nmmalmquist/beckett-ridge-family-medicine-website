const requestForm = document.getElementById("request-form");
const modalTarget = document.querySelector("[data-modal-target]");
document.addEventListener("DOMContentLoaded", () => {
  requestForm.addEventListener(
    "blur",
    (event) => {
      if (event.target.tagName === "INPUT") {
        validateField(event.target, VALIDATION_MAP[event.target.id]);
      }
    },
    true
  );
  requestForm.addEventListener("input", (event) => {
    if (event.target.tagName === "INPUT") {
      if (errMessageExistsOnField(event.target)) {
        validateField(event.target, VALIDATION_MAP[event.target.id]);
      }
    }
  });

  requestForm.addEventListener("submit", async (event) => {
    event.preventDefault();
    const formData = new FormData(event.target);
    const dataArray = [...formData.entries()];
    const formIsValid = validateRequestAppointmentForm(dataArray);
    if (!formIsValid) {
      return;
    }
    const response = await fetch("/api/request-appointment", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: formData,
    });
    const html = await response.text();
    const modalTarget = document.querySelector("[data-modal-target]");
    modalTarget.innerHTML = html;
  });
  modalTarget.addEventListener("click", () => {
    closeModal();
  });
});

const validateName = (value) => {
  if (!value) {
    return "Please provide a full name";
  }
  return "";
};

const validateEmail = (value) => {
  const re =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  if (!re.test(value)) {
    return "Please provide a valid email";
  }
  return "";
};

const validatePhone = (value) => {
  const re = /^\(?([0-9]{3})\)?[-.●]?([0-9]{3})[-.●]?([0-9]{4})$/i;
  if (!re.test(value)) {
    return "Please provide a valid phone number";
  }
  return "";
};

const validateField = (field, validator) => {
  const errMessage = validator(field.value);
  if (errMessage) {
    addErrorToField(field, errMessage);
    return false;
  }
  removeErrorFromField(field);
  return true;
};

const VALIDATION_MAP = {
  "full-name": validateName,
  email: validateEmail,
  phone: validatePhone,
};

const validateRequestAppointmentForm = (data) => {
  let isValid = true;
  data.forEach(([id, value]) => {
    isValid = !VALIDATION_MAP[id]?.(value) && isValid;
    if (VALIDATION_MAP[id]) {
      validateField(document.getElementById(id), VALIDATION_MAP[id]);
    }
  });
  return isValid;
};

const closeModal = (clearForm) => {
  const modal = document.querySelector("[data-modal]");
  const shouldReset = !!document.querySelector("[data-reset-form]");
  shouldReset && requestForm.reset();
  modal.remove();
};
