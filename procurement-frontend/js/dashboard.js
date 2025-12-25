requireAuth();

apiRequest({ method: "GET", url: "/api/items" })
  .done((items) => {
    items.forEach((item) => {
      $("#itemTable").append(`
        <tr>
          <td>${item.Name}</td>
          <td>${item.Stock}</td>
          <td>${item.Price}</td>
        </tr>
      `);
    });
  })
  .fail((err) => {
    Swal.fire("Error", err.responseJSON?.error || "Server error", "error");
  });
