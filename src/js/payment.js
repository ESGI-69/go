const createPayment = async (event) => {
  event.preventDefault();
  const paymentData = new FormData(document.getElementById('payment-form'))
  console.log(paymentData)
  console.log('PricePaid', (paymentData.get('PricePaid')))
  console.log('ProductId', (paymentData.get('ProductId')))
  try {
    const result = await fetch('http://localhost:3000/api/payments', {
      method: 'POST',
      // headers: {
      //   'Content-Type': 'application/json',
      // },
      body: paymentData,
    });
    console.log(result)
    console.log(result.json)
    console.log(result.body)
    await result.json();
    // redirect to /
    // window.location.href = '/';
  } catch (error) {
    console.error(error)
  }

}

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
