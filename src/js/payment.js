const deletePayment = async (paymentId) => {
  try {
    const result = await fetch(`http://localhost:3000/api/payments/${paymentId}`, {
      method: 'DELETE'
    });
    await result.json();
    document.getElementById(`payment-${paymentId}`).remove();
  } catch (error) {
    console.error(error)
  }
}
