const API_BASE = "http://localhost:3000";

function apiRequest(options) {
  return $.ajax({
    method: options.method,
    url: API_BASE + options.url,
    contentType: "application/json",
    data: options.data ? JSON.stringify(options.data) : null,
    headers:
      options.url === "/login"
        ? {}
        : {
            Authorization: "Bearer " + localStorage.getItem("token"),
          },
  });
}

function requireAuth() {
  if (!localStorage.getItem("token")) {
    window.location.href = "login.html";
  }
}
