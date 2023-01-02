const deleteProduct = async (productId) => {
  try {
    const result = await fetch(`http://localhost:3000/api/v1/products/${productId}`, {
      method: 'DELETE'
    });
    await result.json();
    document.getElementById(`product-${productId}`).remove();
  } catch (error) {
    console.error(error)
  }
}

const deletePayment = async (paymentId) => {
  try {
    const result = await fetch(`http://localhost:3000/api/v1/payments/${paymentId}`, {
      method: 'DELETE'
    });
    await result.json();
    document.getElementById(`payment-${paymentId}`).remove();
  } catch (error) {
    console.error(error)
  }
}

const source = new EventSource("http://localhost:3000/api/payments/sse");
// listen of state change 
source.onopen = function() {
  document.getElementById('feed-loading').remove();
};
source.onmessage = function(event) {
  const payment = JSON.parse(event.data);
  payment.CreatedAt = new Date(payment.CreatedAt).toLocaleString();
  payment.UpdatedAt = new Date(payment.UpdatedAt).toLocaleString();

  // Add to payment table
  const paymentTable = document.getElementById("payments-table");
  const paymentRow = document.createElement("tr");
  paymentRow.id = "payment-" + payment.id;
  paymentRow.innerHTML = `
    <td>${payment.ID}</td>
    <td>${payment.ProductID}</td>
    <td>${payment.Product.Name}</td>
    <td>${payment.PricePaid}</td>
    <td>${payment.CreatedAt}</td>
    <td>${payment.UpdatedAt}</td>
    <td><a href="/payments/${payment.ID}">🔎</a></td>
    <td><a href="/payments/${payment.ID}/edit">✏️</a></td>
    <td><a href="#" onclick="deletePayment(${payment.ID})">🗑️</a></td>
  `;
  paymentTable.appendChild(paymentRow);
  
  // Add to feed
  const paymentFeed = document.getElementById("payments");
  const paymentElement = document.createElement("li");
  paymentElement.innerText = `${payment.Product.Name}(${payment.ProductID}) bougth at ${payment.PricePaid}$. Bougth at ${payment.CreatedAt} with ID ${payment.ID}`;
  paymentFeed.appendChild(paymentElement);
};
