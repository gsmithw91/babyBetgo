import { showToast } from "./utils.js";

(function () {
  // ===============================
  // 🔧 DOM Element Cache
  // ===============================
  let modal, modalBtn, closeModal, tabLogin, tabRegister, loginWrap, registerWrap;

  const cacheDom = () => {
    modalBtn = document.getElementById("auth-modal-btn");
    modal = document.getElementById("auth-modal");
    closeModal = document.getElementById("close-modal");
    tabLogin = document.getElementById("tab-login");
    tabRegister = document.getElementById("tab-register");
    loginWrap = document.getElementById("login-form-wrap");
    registerWrap = document.getElementById("register-form-wrap");
  };

  // ===============================
  // 🚪 Modal Show/Hide
  // ===============================
  const setupModalToggle = () => {
    if (!modal || !closeModal) return;

    closeModal.onclick = () => modal.classList.add("hidden");

    window.onclick = (e) => {
      if (e.target === modal) {
        modal.classList.add("hidden");
      }
    };
  };

  // ===============================
  // 🔁 Login / Register Tabs
  // ===============================
  const setupTabs = () => {
    if (!tabLogin || !tabRegister || !loginWrap || !registerWrap) return;

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
  };

  // ===============================
  // 🔐 Login Form Submission
  // ===============================
  const setupLoginForm = () => {
    const form = document.getElementById("login-form");
    const message = document.getElementById("login-message");

    if (!form || !message) return;

    form.addEventListener("submit", async (e) => {
      e.preventDefault();

      const data = {
        username: form.username.value,
        password: form.password.value,
      };

      try {
        const res = await fetch("/login", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(data),
        });

        const result = await res.json();

        if (res.ok) {
          localStorage.setItem("jwt", result.token);
          location.reload(); // refresh to re-init user session
        } else {
          message.textContent = result.error || "Login failed";
        }
      } catch (err) {
        console.error("Login error:", err);
        message.textContent = "An error occurred during login.";
      }
    });
  };

  // ===============================
  // 📝 Register Form Submission
  // ===============================
  const setupRegisterForm = () => {
    const form = document.getElementById("register-form");
    const msg = document.getElementById("register-message");

    if (!form || !msg) return;

    form.onsubmit = async (e) => {
      e.preventDefault();

      const data = {
        username: form.username.value,
        email: form.email.value,
        password: form.password.value,
        phonenumber: form.phonenumber.value,
        display_name: form.display_name.value,
        bio: form.bio.value,
      };

      try {
        const res = await fetch("/register", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(data),
        });

        if (res.ok) {
          showToast("Registration successful!");
          setTimeout(() => {
            document.getElementById("tab-login").click();
            msg.textContent = "";
          }, 1000);
        } else {
          const errorText = await res.text();
          msg.textContent = errorText;
          msg.className = "mt-2 text-red-600 text-sm text-center";
          showToast(errorText, 5000);
        }
      } catch (err) {
        console.error("Registration error:", err);
        msg.textContent = "An error occurred during registration.";
      }
    };
  };

  // ===============================
  // 🧾 Add JWT Token to HTMX Requests
  // ===============================
  const setupJWTInjection = () => {
    document.addEventListener("htmx:configRequest", (event) => {
      const token = localStorage.getItem("jwt");
      if (token) {
        event.detail.headers["Authorization"] = `Bearer ${token}`;
      }
    });
  };

  // ===============================
  // 🚪 Reattach Logout Logic After HTMX Swap
  // ===============================
  const setupLogoutAfterSwap = () => {
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
  };

  // ===============================
  // 🚀 Init (on page load)
  // ===============================
  const init = () => {
    cacheDom();
    setupModalToggle();
    setupTabs();
    setupLoginForm();
    setupRegisterForm();
    setupJWTInjection();
    setupLogoutAfterSwap();
  };

  // Init on initial page load
  document.addEventListener("DOMContentLoaded", init);

  // Re-init modal logic after HTMX injects the modal
  document.body.addEventListener("htmx:afterSwap", (e) => {
    if (e.detail.target.id === "modal-slot") {
      const modal = document.getElementById("auth-modal");
      if (modal) {
        modal.classList.remove("hidden"); // show it
        cacheDom();
        setupModalToggle();
        setupTabs();
        setupLoginForm();
        setupRegisterForm();
      }
    }
  });
})();

