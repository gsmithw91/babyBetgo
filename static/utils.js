// static/utils.js

export function showToast(message, duration = 3000) {
  const toast = document.getElementById("toast");
  const msg = document.getElementById("toast-message");
  msg.textContent = message;
  toast.classList.remove("hidden");
  toast.classList.add("opacity-100");

  setTimeout(() => {
    toast.classList.add("opacity-0");
    setTimeout(() => {
      toast.classList.remove("opacity-100");
      toast.classList.add("hidden");
    }, 300); // match transition duration
  }, duration);
}
