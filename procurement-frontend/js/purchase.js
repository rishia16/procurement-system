requireAuth();

let cart = [];
let itemsMap = {};

apiRequest({ method: "GET", url: "/api/suppliers" }).done((res) => {
  res.forEach((s) =>
    $("#supplier").append(`<option value="${s.ID}">${s.Name}</option>`)
  );
});

apiRequest({ method: "GET", url: "/api/items" }).done((res) => {
  res.forEach((i) => {
    itemsMap[i.ID] = i.Name;
    $("#item").append(`<option value="${i.ID}">${i.Name}</option>`);
  });
});

$("#addItem").on("click", function () {
  const itemId = parseInt($("#item").val());
  const qty = parseInt($("#qty").val());

  if (!qty || qty <= 0) {
    Swal.fire("Error", "Qty harus > 0", "warning");
    return;
  }

  cart.push({ item_id: itemId, qty });
  renderCart();
});

$("#cart").on("click", ".remove", function () {
  cart.splice($(this).data("index"), 1);
  renderCart();
});

function renderCart() {
  $("#cart").empty();
  cart.forEach((c, i) => {
    $("#cart").append(`
      <tr>
        <td>${itemsMap[c.item_id]}</td>
        <td>${c.qty}</td>
        <td><button class="btn btn-danger remove" data-index="${i}">Hapus</button></td>
      </tr>
    `);
  });
}

$("#submitOrder").on("click", function () {
  if (cart.length === 0) {
    Swal.fire("Error", "Keranjang kosong", "warning");
    return;
  }

  apiRequest({
    method: "POST",
    url: "/api/purchasings",
    data: { supplier_id: parseInt($("#supplier").val()), items: cart },
  })
    .done((res) => {
      Swal.fire("Success", "Grand Total: " + res.grand_total, "success");
      cart = [];
      renderCart();
    })
    .fail((err) => {
      Swal.fire("Error", err.responseJSON?.error || "Server error", "error");
    });
});
