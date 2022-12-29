// const createPayment = async (event) => {
//   event.preventDefault();
//   const paymentData = new FormData(document.getElementById('payment-form'))
//   try {
//     const result = await fetch('http://localhost:3000/api/payments', {
//       method: 'POST',
//       headers: {
//         'Content-Type': 'application/json',
//       },
//       body: JSON.stringify({
//         PricePaid: parseFloat(paymentData.get('PricePaid')),
//         ProductID: parseInt(paymentData.get('ProductID')),
//       }),
//     });
//     await result.json();
//     window.location.href = '/';
//   } catch (error) {
//     console.error(error)
//   }
// }
function createPayment(event) {
  event.preventDefault();
  const paymentData = new FormData(document.getElementById('payment-form'))
  fetch('http://localhost:3000/api/payments/', {
    method: 'POST',
    body: JSON.stringify({
      PricePaid: parseFloat(paymentData.get('PricePaid')),
      ProductID: parseInt(paymentData.get('ProductID')),
    }),
  })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      window.location.href = '/';
    }
    )
    .catch((error) => {
      console.error('Error:', error);
    }
    );
}

function editPayment(event) {
  event.preventDefault();
  const paymentData = new FormData(document.getElementById('payment-form'))
  fetch(`http://localhost:3000/api/payments/${paymentData.get('PaymentId')}`, {
    method: 'PATCH',
    body: JSON.stringify({
      PricePaid: parseFloat(paymentData.get('PricePaid')),
      ProductID: parseInt(paymentData.get('ProductID')),
    }),
  })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      window.location.href = '/';
    }
    )
    .catch((error) => {
      console.error('Error:', error);
    }
    );
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
