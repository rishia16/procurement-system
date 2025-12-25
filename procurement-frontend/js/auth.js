$("#btnLogin").on("click", function () {
  const username = $("#username").val().trim();
  const password = $("#password").val().trim();

  if (!username || !password) {
    Swal.fire("Error", "Username dan Password wajib diisi", "warning");
    return;
  }

  apiRequest({
    method: "POST",
    url: "/login",
    data: { username, password },
  })
    .done((res) => {
      localStorage.setItem("token", res.token);
      window.location.href = "dashboard.html";
    })
    .fail((err) => {
      Swal.fire(
        "Login Gagal",
        err.responseJSON?.error || "Server error",
        "error"
      );
    });
});
