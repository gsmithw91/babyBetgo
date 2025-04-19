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
    document.getElementById("login-form").onsubmit = async function(e) {
      e.preventDefault();
      const form = e.target;
      console.log("Login form submitted");
      try {
        console.log("Sending login request:", form.username.value);
        const res = await fetch("/login", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({
            username: form.username.value,
            password: form.password.value,
          }),
        });
        const msg = document.getElementById("login-message");
        console.log("Login response status:", res.status);
        if (res.ok) {
          msg.textContent = "Login successful!";
          msg.className = "mt-2 text-green-600 text-sm text-center";
          setTimeout(() => {
            modal.classList.add("hidden");
            // Refresh balance value using HTMX (if htmx is loaded)
            if (window.htmx) {
              const balanceEl = document.querySelector("#balance-value");
              if (balanceEl && window.htmx) {
                window.htmx.trigger(balanceEl, "refresh");
              }
            }
          }, 1000);
        } else {
          const errorText = await res.text();
          console.error("Login error:", errorText);
          msg.textContent = errorText;
          msg.className = "mt-2 text-red-600 text-sm text-center";
        }
      } catch (err) {
        console.error("Network error during login:", err);
      }
    };
    // Handle register
    document.getElementById("register-form").onsubmit = async function(e) {
      e.preventDefault();
      const form = e.target;
      console.log("Register form submitted");
      try {
        console.log(
          "Sending register request:",
          form.username.value,
          form.email.value,
          form.display_name.value,
        );
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
        console.log("Sending Request Payload", {
          username: form.username.value,
          email: form.email.value,
          password: form.password.value,
          phonenumber: form.phonenumber.value,
          display_name: form.display_name.value,
          bio: form.bio.value,
        });
        const msg = document.getElementById("register-message");
        console.log("Register response status:", res.status);
        if (res.ok) {
          msg.textContent = "Registration successful!";
          msg.className = "mt-2 text-green-600 text-sm text-center";
          setTimeout(() => {
            tabLogin.click();
            msg.textContent = "";
          }, 1000);
        } else {
          const errorText = await res.text();
          console.error("Register error:", errorText);
          msg.textContent = errorText;
          msg.className = "mt-2 text-red-600 text-sm text-center";
        }
      } catch (err) {
        console.error("Network error during registration:", err);
      }
    };
  }
})();
