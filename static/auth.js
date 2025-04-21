
import { showToast } from "./utils.js";

// Modal logic for auth (login/register)
(function() {
  let modalBtn = document.getElementById("auth-modal-btn");
  let modal = document.getElementById("auth-modal");
  let closeModal = document.getElementById("close-modal");
  let tabLogin = document.getElementById("tab-login");
  let tabRegister = document.getElementById("tab-register");
  let loginWrap = document.getElementById("login-form-wrap");
  let registerWrap = document.getElementById("register-form-wrap");

  if (
    modalBtn &&
    modal &&
    closeModal &&
    tabLogin &&
    tabRegister &&
    loginWrap &&
    registerWrap
  ) {
    modalBtn.onclick = () => modal.classList.remove("hidden");
    closeModal.onclick = () => modal.classList.add("hidden");

    tabLogin.onclick = () => {
      tabLogin.classList.add("text-indigo-600", "border-indigo-600");
      tabLogin.classList.remove("text-gray-500", "border-transparent");
      tabRegister.classList.remove("text-indigo-600", "border-indigo-600");
      tabRegister.classList.add("text-gray-500", "border-transparent");
      loginWrap.classList.remove("hidden");
      registerWrap.classList.add("hidden");
    };

    tabRegister.onclick = () => {
      tabRegister.classList.add("text-indigo-600", "border-indigo-600");
      tabRegister.classList.remove("text-gray-500", "border-transparent");
      tabLogin.classList.remove("text-indigo-600", "border-indigo-600");
      tabLogin.classList.add("text-gray-500", "border-transparent");
      registerWrap.classList.remove("hidden");
      loginWrap.classList.add("hidden");
    };

    window.onclick = function(event) {
      if (event.target === modal) {
        modal.classList.add("hidden");
      }
    };

    // Handle login
    document
      .getElementById("login-form")
      .addEventListener("submit", async (e) => {
        e.preventDefault();
        const form = e.target;
        const data = {
          username: form.username.value,
          password: form.password.value,
        };

        const res = await fetch("/login", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(data),
        });

        const result = await res.json();

        if (res.ok) {
          localStorage.setItem("jwt", result.token);
          location.reload(); // refresh to re-init auth state
        } else {
          document.getElementById("login-message").textContent =
            result.error || "Login failed";
        }
      });

    // Handle register
    document.getElementById("register-form").onsubmit = async function(e) {
      e.preventDefault();
      const form = e.target;

      try {
        const res = await fetch("/register", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            username: form.username.value,
            email: form.email.value,
            password: form.password.value,
            phonenumber: form.phonenumber.value,
            display_name: form.display_name.value,
            bio: form.bio.value,
          }),
        });

        const msg = document.getElementById("register-message");

        if (res.ok) {
          showToast("Registration successful!");
          setTimeout(() => {
            tabLogin.click();
            msg.textContent = "";
          }, 1000);
        } else {
          const errorText = await res.text();
          msg.textContent = errorText;
          msg.className = "mt-2 text-red-600 text-sm text-center";
          showToast(errorText, 5000);
        }
      } catch (err) {
        console.error("Network error during registration:", err);
      }
    };
  }

  // Inject JWT token into all HTMX requests
  document.addEventListener("htmx:configRequest", (event) => {
    const token = localStorage.getItem("jwt");
    if (token) {
      event.detail.headers["Authorization"] = `Bearer ${token}`;
    }
  });

  // Automatically reattach logout behavior after HTMX updates
  document.body.addEventListener("htmx:afterSwap", (event) => {
    if (event.target.id === "user-info") {
      const logoutBtn = document.getElementById("logout-btn");
      if (logoutBtn) {
        logoutBtn.onclick = () => {
          localStorage.removeItem("jwt");
          showToast("Logged out!");
          setTimeout(() => location.reload(), 1000);
        };
      }
    }
  });
})();
